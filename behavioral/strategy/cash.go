package strategy

import "fmt"

type Cash struct {
}

func (p *Cash) MakePayment(amount float64) {
	fmt.Printf("Payment by Cash method, amount: %f \n", amount)
}
