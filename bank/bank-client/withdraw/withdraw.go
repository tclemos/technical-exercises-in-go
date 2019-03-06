package withdraw

type WithdrawID string

type ContextID string

type Request struct {
	ContextID ContextID
	Amount    float64
}

type Response struct {
	ID     WithdrawID
	Amount float64
}

func Withdraw(request Request) (Response, error) {
	return Response{}, nil
}
