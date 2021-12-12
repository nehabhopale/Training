package main

import (
	calc "assignment/calc"
	mypack "assignment/mypackage"
	"fmt"
)

func main() {
	fmt.Println(calc.Add(1, 2))
	fmt.Println(calc.Sub(7, 3))
	mypack.MyFunction()
}
