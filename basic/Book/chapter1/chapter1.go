package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("cannonball")
	fmt.Println("hello,\n Go")
	fmt.Println("hello,\t Go")
	fmt.Println("hello,\"\" Go")
	fmt.Println("hello,\\ Go")
	fmt.Println(reflect.TypeOf(25))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(5.1))
	fmt.Println(reflect.TypeOf("hello"))
	prizeWithTax()
}
func prizeWithTax() {
	var price int = 100
	fmt.Println("Price is", price, "dolloars.")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("tax is", tax, "dollors.")

	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "dollors.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("within budget?", total <= float64(availableFunds))

}
