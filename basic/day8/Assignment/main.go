package main

import (
	"fmt"
	Customer "hotel/Customer"
	Restaurant "hotel/Restaurant"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	//item := []string{"pavBhaji", "tea", "idli", "dosa"}

	neha := Customer.New("neha", "neha@123", 5612345892, "@124", 1000)
	pooja := Customer.New("pooja", "pooja@345", 7634126739, "@567t", 200)
	shardha := Customer.New("shradha", "shradha@345", 7634187639, "@98t", 300)
	Customer.SignUp(neha)
	Customer.SignUp(pooja)
	Purohit := Restaurant.New("purohit")
	// fmt.Println("choose 0  for pav Bhaji;1 for tea;2 for idli;3 for dosa")
	// var choice int
	// fmt.Scanf("%d", &choice)
	Purohit.AddNewItem("pav bhaji", 80, true, 4)
	Purohit.AddNewItem("tea", 10, false, 4)
	Purohit.AddNewItem("idli", 50, true, 5)
	Purohit.AddNewItem("dosa", 60, true, 4)

	Purohit.GetRestDetails()
	wg.Add(4)

	go neha.PlaceOrder(Purohit, "pav bhaji", &wg)
	go shardha.PlaceOrder(Purohit, "tea", &wg)
	go neha.PlaceOrder(Purohit, "idli", &wg)
	go pooja.PlaceOrder(Purohit, "dosa", &wg)
	wg.Wait()
	fmt.Println(neha.GetBalance())
	fmt.Println(shardha.GetBalance())
	fmt.Println(pooja.GetBalance())

}
