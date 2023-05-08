package strategy

import "fmt"

type InKind struct {
}

func (p *InKind) MakePayment(amount float64) {
	fmt.Printf("It is awful, but you paid all the bills, amount: %f \n", amount)
}
