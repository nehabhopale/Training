package main

import ("fmt"
"pizza/pizzamenu")

func main() {

	extravagnzaPizza := &pizzamenu.ExtraVaganza {}

	//Add cheese topping
	extravagnzaPizzaWithCheese := &pizzamenu.CheeseTopping{
		Pizza: extravagnzaPizza,
	}

	fmt.Println("pizza with cheese",extravagnzaPizzaWithCheese.GetPrize())
	//Add onion topping
	extravagnzaPizzaWithOnion := &pizzamenu.OnionTopping{
		Pizza: extravagnzaPizza,
	 }
	 fmt.Println("pizza with onion",extravagnzaPizzaWithOnion.GetPrize())

	 extravagnzaPizzaWithCheeseAnOnion:=&pizzamenu.OnionTopping{
		Pizza: extravagnzaPizzaWithCheese,
	}
	fmt.Println("pizza with cheese and Onion",extravagnzaPizzaWithCheeseAnOnion.GetPrize())

	
	
}