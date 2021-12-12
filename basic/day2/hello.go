package main

import (
	"fmt"
	"reflect"
	"test/math"
	"time"
)

func main() {
	fmt.Println("hi")
	fmt.Println("addition of two numbers", math.AddTwoNumber(4, 5))
	num1 := 2
	num2 := 2.4
	name := "neha"
	fmt.Println(reflect.TypeOf(num1))
	fmt.Printf("number2 type is %T\n", num2)
	fmt.Println(reflect.TypeOf(name))
	greeting()
	day()
	evenNumbers()
}

func greeting() {
	Time := time.Now()
	Currenthour := Time.Hour()
	if 8 <= Currenthour && Currenthour < 12 {
		fmt.Println("good morning")
	} else if 12 <= Currenthour && Currenthour < 16 {
		fmt.Println("afternoon")
	} else if 16 <= Currenthour && Currenthour <= 20 {
		fmt.Println("evening")
	} else {
		fmt.Println("night")
	}

}

func day() {
	day := time.Now().Weekday()
	switch day {
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday:
		fmt.Println("Tuesday")
	case time.Wednesday:
		fmt.Println("Wednesday")
	case time.Thursday:
		fmt.Println("Thursday")
	case time.Friday:
		fmt.Println("Friday")
	case time.Saturday:
		fmt.Println("Saturday")
	case time.Sunday:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}

}
func evenNumbers() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		i++
	}
}
