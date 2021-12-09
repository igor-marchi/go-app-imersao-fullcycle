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

	_, invalidCC := entity.NewCreditCard(
		input.CreditCardNumber,
		input.CreditCardName,
		input.CreditCardExpirationMonth,
		input.CreditCardExpirationYear,
		input.CreditCardCVV)

	if invalidCC != nil {
		err := p.Repository.Insert(transaction.Id, transaction.AccountId, transaction.Amount, entity.REJECTED, invalidCC.Error())
		if err != nil {
			return TransactionDtoOutput{}, err
		}

		output := TransactionDtoOutput{
			Id:           transaction.Id,
			Status:       entity.REJECTED,
			ErrorMessage: invalidCC.Error(),
		}

		return output, nil
	}

	return TransactionDtoOutput{}, nil
}
