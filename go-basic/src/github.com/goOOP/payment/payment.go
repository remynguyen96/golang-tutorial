package payment

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

type PaymentOption interface {
	ProcessPayment(float32) bool
}

type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

func CreateCreditAccount(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
}

func (c *CreditCard) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a credit card payment ...")
	return true
}

func (c CreditCard) OwnerName() string {
	return c.ownerName
}

	func (c *CreditCard) SetOwnerName(value string) error {
		if len(value) == 0 {
		return errors.New("Invalid owner name provided!")
	}

	c.ownerName = value
	return nil
}

func (c CreditCard) CardNumber() string {
	return c.cardNumber
}

var cardNumberPattern = regexp.MustCompile("\\d{4}-\\d{4}-\\d{4}-\\d{4}")
func (c *CreditCard) SetCardNumber(value string) error {
	if !cardNumberPattern.Match([]byte(value)) {
		return errors.New("Invalid credit card number format")
	}

	c.cardNumber = value
	return nil
}

func (c CreditCard) ExpirationDate() (int, int) {
	return c.expirationMonth, c.expirationYear
}

func (c *CreditCard) SetExpirationDate(month, year int) error {
	now := time.Now()
	if year < now.Year() || (year == now.Year() && time.Month(month) < now.Month()) {
		return errors.New("Expiration date must lie in the future")
	}

	c.expirationMonth, c.expirationYear = month, year
	return nil
}

func (c CreditCard) SecurityCode() int {
	return c.securityCode
}

func (c *CreditCard) SetSecurityCode(value int) {
	if value < 100 || value > 999 {
		//return errors.New("Invalid ")
	}
}

func (c CreditCard) AvailableCredit() float32 {
	return 5000
}

type CheckingAccount struct {}

func (c CheckingAccount) ProcessPayment(amount float32) bool {
	fmt.Println("Paying with checking account")
	return true
}
