package strategy

import "fmt"

type Visa struct {
}

func (p *Visa) MakePayment(amount float64) {
	fmt.Printf("Payment by Visa method, amount: %f \n", amount)
}
