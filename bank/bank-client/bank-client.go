package bankclient

import (
	"github.com/tclemos/technical-exercises-in-go/bank/bank-client/account"
	"github.com/tclemos/technical-exercises-in-go/bank/bank-client/deposit"
	"github.com/tclemos/technical-exercises-in-go/bank/bank-client/session"
)

type CreateSessionRequest struct {
	AccountID account.ID
	Password  account.Password
}

func CreateDepositContext() (deposit.ContextID, error) {
	return "", nil
}

func CreateSession(request CreateSessionRequest) (session.SessionID, error) {
	return "", nil
}
