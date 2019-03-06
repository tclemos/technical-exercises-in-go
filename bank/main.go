package main

import (
	"log"

	"github.com/tclemos/technical-exercises-in-go/bank/bank-client/account"

	"github.com/tclemos/technical-exercises-in-go/bank/bank"
	bankclient "github.com/tclemos/technical-exercises-in-go/bank/bank-client"
	"github.com/tclemos/technical-exercises-in-go/bank/bank-client/deposit"
	"github.com/tclemos/technical-exercises-in-go/bank/bank-client/session"
	twofactorauth "github.com/tclemos/technical-exercises-in-go/bank/bank-client/two-factor-auth"
	"github.com/tclemos/technical-exercises-in-go/bank/bank-client/withdraw"
)

func main() {

	go bank.Start()

	var customerAccountID account.ID = "tclemos"
	var customerPassword account.Password = "anypassword"
	var customerToken twofactorauth.Token = "anytoken"

	doASimpleDeposit(customerAccountID)
	doASimpleWithdraw(customerAccountID, customerPassword, customerToken)
	// performConcurrentWithdraws()
}

func doASimpleDeposit(customerAccountID account.ID) {

	log.Println("Doing a simple deposit")

	depositContextID, err := bankclient.CreateDepositContext()

	if err != nil {
		log.Println(err)
		return
	} else {
		log.Printf("Deposit context ID: $s\n", depositContextID)
	}

	depositRequest := deposit.Request{
		ContextID: depositContextID,
		AccountID: customerAccountID,
		Amount:    10.5,
	}

	if response, err := deposit.Deposit(depositRequest); err != nil {
		log.Println(err)
	} else {
		log.Printf("Deposit worked successfully: %v\n", response)
	}

	log.Println("Simple deposit done!")
}

func doASimpleWithdraw(customerAccountID account.ID, customerPassword account.Password, token twofactorauth.Token) {
	log.Println("Doing a simple withdraw")

	createSessionRequest := bankclient.CreateSessionRequest{
		AccountID: customerAccountID,
		Password:  customerPassword,
	}

	sessionID, err := bankclient.CreateSession(createSessionRequest)
	if err != nil {
		log.Println(err)
		return
	}

	createWithdrawContextRequest := session.WithdrawContextRequest{
		SessionID: sessionID,
		Token:     token,
	}

	withdrawContextID, err := session.CreateWithdrawContext(createWithdrawContextRequest)
	if err != nil {
		log.Println(err)
		return
	}

	withdrawRequest := withdraw.Request{
		ContextID: withdrawContextID,
		Amount:    7,
	}

	withdrawResponse, err := withdraw.Withdraw(withdrawRequest)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Simple withdraw done!")
}

// func performConcurrentWithdraws() {
// }
