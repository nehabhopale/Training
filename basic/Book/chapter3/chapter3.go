package main

import (
	"errors"
	"fmt"
)

func fA(a int, b int) {
	fmt.Println(a + b)
}
func fB(a int, b int) {
	fmt.Println(a * b)
}
func fC(a bool) {
	fmt.Println(!a)
}
func fD(a string, b int) {
	for i := 0; i < b; i++ {
		fmt.Print(a)
	}
	fmt.Println()
}
func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return dividend / divisor, nil

}
func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}
func pointer() {
	var myInt int
	var myIntPointer *int
	myInt = 42
	myIntPointer = &myInt
	fmt.Println(*myIntPointer)
}
func main() {
	fA(2, 3) //5
	fB(2, 3) //6
	fC(true) //false
	fD("$", 4)
	fA(5, 6)  //11
	fB(5, 6)  //30
	fC(false) //false
	fD("ha", 4)
	quotient, err := divide(5.6, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
	pointer()
}
