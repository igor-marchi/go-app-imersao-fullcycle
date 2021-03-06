package process_transaction

import (
	"github.com/fullcycle/gateway/domain/entity"
	"github.com/fullcycle/gateway/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
}

func NewProcessTransaction(repository repository.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := entity.NewTransaction()
	transaction.Id = input.Id
	transaction.AccountId = input.AccountId
	transaction.Amount = input.Amount

	cc, invalidCC := entity.NewCreditCard(
		input.CreditCardNumber,
		input.CreditCardName,
		input.CreditCardExpirationMonth,
		input.CreditCardExpirationYear,
		input.CreditCardCVV)

	if invalidCC != nil {
		return p.rejectTransaction(transaction, invalidCC)
	}

	transaction.SetCreditCard(*cc)
	invalidTransaction := transaction.IsValid()
	if invalidTransaction != nil {
		return p.rejectTransaction(transaction, invalidTransaction)
	}

	return p.approveTransaction(transaction)
}

func (p *ProcessTransaction) approveTransaction(transaction *entity.Transaction) (TransactionDtoOutput, error) {
	err := p.Repository.Insert(transaction.Id, transaction.AccountId, transaction.Amount, entity.APPROVED, "")
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		Id:           transaction.Id,
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}
	return output, nil
}

func (p *ProcessTransaction) rejectTransaction(transaction *entity.Transaction, invalidTransaction error) (TransactionDtoOutput, error) {
	err := p.Repository.Insert(transaction.Id, transaction.AccountId, transaction.Amount, entity.REJECTED, invalidTransaction.Error())
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		Id:           transaction.Id,
		Status:       entity.REJECTED,
		ErrorMessage: invalidTransaction.Error(),
	}

	return output, nil
}
