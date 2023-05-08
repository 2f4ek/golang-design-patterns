package strategy

type IPayment interface {
	MakePayment(amount float64)
}
