package main

import (
	"log"

	"InterfaceExample/concretes"
	"InterfaceExample/interfaces"
)


func main() {
	icici := concretes.NewICICI()

	bal, _ := icici.DepositBalance(20)
	log.Println(bal)

	icici.DepositBalance(30)

	log.Println(icici.GetBalance())

	myAccounts := []interfaces.IBankAccount{
		concretes.NewICICI(),
		concretes.NewSBI(),
	}

	for _, accountToBeAdded := range myAccounts {
		accountToBeAdded.DepositBalance(100)
	}

	for _, account := range myAccounts {
		balance := account.GetBalance()
		log.Println(balance)
	}

	vehicles := []interfaces.IVehicle{
		concretes.NewElectricVehicle(),
		concretes.NewICVehicle(),
	}

	for _,v := range vehicles {
		v.Start()
		v.IsRunning()
		v.Stop()
		v.IsRunning()
	}

	randomStuff := []interfaces.IAbstract{
		concretes.NewElectricVehicle(),
		concretes.NewICICI(),
		concretes.NewICVehicle(),
		concretes.NewSBI(),
	}

	for _,r := range randomStuff{
		r.DoRandom()
	}
}
