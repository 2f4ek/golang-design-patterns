package main

import (
	"fmt"
	"github.com/2f4ek/golang-design-patterns/behavioral/iterator"
	"github.com/2f4ek/golang-design-patterns/behavioral/strategy"
	"github.com/2f4ek/golang-design-patterns/creational/factory"
	"github.com/2f4ek/golang-design-patterns/creational/singleton"
	"github.com/2f4ek/golang-design-patterns/structural/proxy"
)

func main() {
	initSingletonPattern()
	initMinionsWaveByFactoryPattern()
	makeSomePaymentsWithStrategyPattern()
	letsCountMentalProblemsWithIteratorPattern()
	letsSendSomeRequestsToNotExistingAppWithProxyPattern()
}

func initSingletonPattern() {
	singleton.GetInstance()
	singleton.GetInstance()
	singleton.GetInstance()
	singleton.GetInstance()
	singleton.GetInstance()
	singleton.GetInstance()
}

func makeSomePaymentsWithStrategyPattern() {
	var paymentStrategy strategy.IPayment
	method := "inKind"

	switch method {
	case "inKind":
		paymentStrategy = &strategy.InKind{}
	case "visa":
		paymentStrategy = &strategy.Visa{}
	default:
		paymentStrategy = &strategy.Cash{}
	}

	paymentStrategy.MakePayment(100.01)
}

func initMinionsWaveByFactoryPattern() {
	rangeMinion, _ := factory.InitMinion("range")
	siegeMinion, _ := factory.InitMinion("siege")
	meleeMinion, _ := factory.InitMinion("melee")

	rangeMinion.GetDetails()
	siegeMinion.GetDetails()
	meleeMinion.GetDetails()
}

func letsCountMentalProblemsWithIteratorPattern() {
	mentalProblem1 := &iterator.MentalProblem{
		Name:         "Depression",
		AgeOfReceipt: 25,
	}
	mentalProblem2 := &iterator.MentalProblem{
		Name:         "Bipolar Disorder",
		AgeOfReceipt: 22,
	}

	mpCollection := &iterator.MentalProblemCollection{
		MentalProblems: []*iterator.MentalProblem{mentalProblem1, mentalProblem2},
	}

	myIterator := mpCollection.CreateIterator()

	for myIterator.HasNext() {
		problem := myIterator.GetNext()
		fmt.Printf("It is necessary to pay attention for %s\n", problem.Name)
	}
}

func letsSendSomeRequestsToNotExistingAppWithProxyPattern() {
	nginxServer := proxy.NewNginxServer()
	registerUserUrl := "/user"
	loginUserUrl := "/user/login"

	//success call
	httpCode, method := nginxServer.HandleRequest(registerUserUrl, "POST")
	fmt.Printf("Url: %s, HttpCode: %d, Method: %s\n", registerUserUrl, httpCode, method)
	//wrong endpoint
	httpCode, method = nginxServer.HandleRequest(registerUserUrl, "GET")
	fmt.Printf("Url: %s, HttpCode: %d, Method: %s\n", registerUserUrl, httpCode, method)
	//success call
	httpCode, method = nginxServer.HandleRequest(loginUserUrl, "POST")
	fmt.Printf("Url: %s, HttpCode: %d, Method: %s\n", loginUserUrl, httpCode, method)
	//wrong endpoint
	httpCode, method = nginxServer.HandleRequest(loginUserUrl, "GET")
	fmt.Printf("Url: %s, HttpCode: %d, Method: %s\n", loginUserUrl, httpCode, method)

	// force access denied
	httpCode, method = nginxServer.HandleRequest(loginUserUrl, "POST")
	fmt.Printf("Url: %s, HttpCode: %d, Method: %s\n", loginUserUrl, httpCode, method)
	httpCode, method = nginxServer.HandleRequest(loginUserUrl, "POST")
	fmt.Printf("Url: %s, HttpCode: %d, Method: %s\n", loginUserUrl, httpCode, method)
}
