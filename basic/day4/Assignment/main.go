package main

import (
	customer "bankinfo/bankcustomer"
	"fmt"
)

func main() {
	neha := customer.New(2, "neha", 123, 1000, "savings", 1223)
	pooja := customer.New(4, "pooja", 456, 3000, "savings", 3456)

	neha.AddNewAccount(333, 1002, "savings", 1223)
	pooja.AddNewAccount(444, 6000, "savings", 3456)

	neha.AddNewAccount(222, 3000, "current", 1223)
	pooja.AddNewAccount(111, 2000, "current", 3456)

	fmt.Println("Before Transaction...")
	neha.GetCustomerDetails()
	fmt.Println("...................")
	pooja.GetCustomerDetails()

	err := neha.MoneyTransfer(pooja, 333, 444, 1223, 1000, "savings")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("After Transaction...")
		neha.GetCustomerDetails()
		fmt.Println("-----------------")
		pooja.GetCustomerDetails()
		fmt.Println("")
	}

}
