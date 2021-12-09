package process_transaction

import (
	"testing"
	"time"

	"github.com/fullcycle/gateway/domain/entity"
	mock_repository "github.com/fullcycle/gateway/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	input := TransactionDtoInput{
		Id:                        "1",
		AccountId:                 "1",
		CreditCardNumber:          "400000000000000000000000",
		CreditCardName:            "Wesley Silva",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    200,
	}

	expectedOutput := TransactionDtoOutput{
		Id:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.Id, input.AccountId, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)

	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransaction_ExecuteRejectedTransaction(t *testing.T) {
	input := TransactionDtoInput{
		Id:                        "1",
		AccountId:                 "1",
		CreditCardNumber:          "4193523830170205",
		CreditCardName:            "Wesley Silva",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    1200,
	}

	expectedOutput := TransactionDtoOutput{
		Id:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you dont have limit for this transaction",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.Id, input.AccountId, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)

	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransaction_ExecuteApprovedTransaction(t *testing.T) {
	input := TransactionDtoInput{
		Id:                        "1",
		AccountId:                 "1",
		CreditCardNumber:          "4193523830170205",
		CreditCardName:            "Wesley Silva",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    900,
	}

	expectedOutput := TransactionDtoOutput{
		Id:           "1",
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.Id, input.AccountId, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)

	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
