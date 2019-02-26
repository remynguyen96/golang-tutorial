package main

import (
	"fmt"
	"github.com/goOOP/paybroker"
	"github.com/goOOP/payment"
)

type CreditAccount struct{}

func (c *CreditAccount) processPayment(amount float32) {
	fmt.Println("Processing credit card payment", amount)
}

func CreateCreditAccount(chargeCh chan float32) *CreditAccount {
	creditAccount := &CreditAccount{}
	go func(charge chan float32) {
		for amount := range chargeCh {
			creditAccount.processPayment(amount)
		}
	}(chargeCh)

	return creditAccount
}

// --------------------------------------
type Account struct {}

func (a *Account) AvailableFunds() float32 {
	fmt.Println("Getting account funds")
	return 0
}
func (a *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Process payment with", amount)
	return true
}

type ExtentAccount struct {
	Account
}

// --------------------------------------

type CreditAccountNext struct {}

func (a *CreditAccountNext) AvailableFunds() float32 {
	fmt.Println("Getting credit funds")
	return 250
}

type CheckingAccount struct {}

func (c *CheckingAccount) AvailableFunds() float32 {
	fmt.Println("Getting checking funds")
	return 125
}

type HyBridAccount struct {
	creditAccount CreditAccountNext
	CheckingAccount
}

func (h *HyBridAccount) AvailableFunds() float32 {
	return h.creditAccount.AvailableFunds() + h.CheckingAccount.AvailableFunds()
}

func main() {
	var option payment.PaymentOption
	option = &payment.CreditCard{}
	option.ProcessPayment(500)

	option = &payment.CheckingAccount{}
	option.ProcessPayment(500)

	option = &paybroker.PaymentBrokerAccount{}
	option.ProcessPayment(500)


/*	ha := &HyBridAccount{}
	fmt.Println(ha.AvailableFunds())*/

	/*	ca1 := &ExtentAccount{}
	ca1.AvailableFunds()
	ca1.ProcessPayment(300)*/

	/*chargeCh := make(chan float32)
	CreateCreditAccount(chargeCh)
	chargeCh <- 500
	var a string
	fmt.Scanln(&a)*/

	/*	var option payment.PaymentOption
		option = payment.CreateCreditAccount(
			"Debra Ingram",
			"1111-2222-3333-4444",
			5,
			2021,
			123,
		)
		option.ProcessPayment(500)
		option = payment.CreateCashAccount()
		option.ProcessPayment(500)*/

	/*	credit := payment.CreateCreditAccount(
			"Debra Ingram",
			"1111-2222-3333-4444",
			5,
			2021,
			123,
		)
		fmt.Printf("Owner name: %v\n", credit.OwnerName())
		fmt.Printf("Card number: %v\n", credit.CardNumber())
		fmt.Printf("Security code: %v\n", credit.SecurityCode())
		fmt.Println("Trying to change card number")
		err := credit.SetCardNumber("invalid")
		if err != nil {
			fmt.Printf("That didn't work: %v\n", err)
		}
		fmt.Printf("Available credit: %v\n", credit.AvailableCredit())*/

}
