package concretes

import (
	"log"
)

type SBI struct {
	balance int
	fee     int
}

func NewSBI() *SBI {
	return &SBI{
		balance: 0,
		fee:     10,
	}
}

func (icici *SBI) GetBalance() int {
	log.Println("Getting Balance")
	return icici.balance
}

func (icici *SBI) DepositBalance(amount int) (int, error) {
	log.Println("Depositing {amount}")
	icici.balance += amount - icici.fee
	return icici.balance, nil
}

func (icici *SBI) WithdrawBalance(amount int) (int, error) {
	log.Printf("Withdrwaing {amount}")
	icici.balance -= amount + icici.fee
	return icici.balance, nil
}

func (c *SBI) DoRandom() {
	log.Println("SBI doing random stuff")
}
