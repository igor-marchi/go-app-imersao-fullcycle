package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_IsValid(t *testing.T) {
	transaction := NewTransaction()
	transaction.Id = "1"
	transaction.AccountId = "1"
	transaction.Amount = 900

	assert.Nil(t, transaction.IsValid())
}

func TestTransaction_IsNotValidWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.Id = "1"
	transaction.AccountId = "1"
	transaction.Amount = 1001
	err := transaction.IsValid()

	assert.Error(t, err)
	assert.Equal(t, "you dont have limit for this transaction", err.Error())
}

func TestTransaction_IsNotValidWithAmountLessThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.Id = "1"
	transaction.AccountId = "1"
	transaction.Amount = 0
	err := transaction.IsValid()

	assert.Error(t, err)
	assert.Equal(t, "the amount must be grater than one", err.Error())
}
