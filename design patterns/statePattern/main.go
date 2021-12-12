package main

import (
	"fmt"
	"log"
	"machine/state"
)

func main() {
	vendingMachine := state.New(1, 20) //1 machine of prize 20
	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.InsertMoney(20)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
