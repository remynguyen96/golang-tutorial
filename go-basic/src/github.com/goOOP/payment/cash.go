package payment

import "fmt"

type Cash struct {}

func CreateCashAccount() *Cash {
	return &Cash{}
}

func (c Cash) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a cash transaction...")
	return true
}
