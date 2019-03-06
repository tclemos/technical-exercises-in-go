package deposit

import (
	"github.com/tclemos/technical-exercises-in-go/bank/bank-client/account"
)

type ContextID string
type DepositID string

type Request struct {
	ContextID ContextID
	AccountID account.ID
	Amount    float64
}

type Response struct {
	ID DepositID
}

func Deposit(request Request) (Response, error) {
	return Response{}, nil
}
