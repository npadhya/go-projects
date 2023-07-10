package concretes

import (
	"log"
)

type ICICI struct {
	balance int
}

func NewICICI() *ICICI {
	return &ICICI{
		balance: 0,
	}
}

func (icici *ICICI) GetBalance() int {
	log.Println("Getting Balance")
	return icici.balance
}

func (icici *ICICI) DepositBalance(amount int) (int, error) {
	log.Println("Depositing {amount}")
	icici.balance += amount
	return icici.balance, nil
}

func (icici *ICICI) WithdrawBalance(amount int) (int, error) {
	log.Printf("Withdrwaing {amount}")
	icici.balance -= amount
	return icici.balance, nil
}

func (c *ICICI) DoRandom() {
	log.Println("ICICI doing random stuff")
}
