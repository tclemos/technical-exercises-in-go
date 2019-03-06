package session

import (
	twofactorauth "github.com/tclemos/technical-exercises-in-go/bank/bank-client/two-factor-auth"
	"github.com/tclemos/technical-exercises-in-go/bank/bank-client/withdraw"
)

type SessionID string
type Token twofactorauth.Token

type WithdrawContextRequest struct {
	SessionID SessionID
	Token     twofactorauth.Token
}

func CreateWithdrawContext(request WithdrawContextRequest) (withdraw.ContextID, error) {
	return "", nil
}
