package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// testify package contains assertions for various types of unit tests
func TestCreateAccount(t *testing.T) {

	arg := CreateAccountParams{
		Owner:    "Gerardo",
		Balance:  100000,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	// verify if there is no error
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
