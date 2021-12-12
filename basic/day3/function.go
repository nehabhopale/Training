package main

import "fmt"

func a() func(float32, float32, func(float32, float32) float64) float64 {
	fmt.Println("inside a function")
	/*b := func(f1 float32, f2 float32, f func(float32, float32) float64) float64 {
		fmt.Println("function inside function")
		return f(f1, f2)
	}
	return b*/
	return func(f1 float32, f2 float32, f func(float32, float32) float64) float64 {
		return float64(f(f1, f2))
	}
}
func mathoperation(f1, f2 float32, f func(float32, float32) float64) float64 {
	return float64(f(f1, f2))
}
func add(num1, num2 float32) float64 {
	return float64(num1 + num2)
}
func sub(num1, num2 float32) float64 {
	return float64(num1 - num2)
}
func mul(num1, num2 float32) float64 {
	return float64(num1 * num2)
}
func div(num1, num2 float32) float64 {
	return float64(num1 / num2)
}
func main() {
	f := a()
	fmt.Println(f(12.3, 13.0, add))
	fmt.Println(f(12.3, 13.0, sub))
	fmt.Println(f(12.3, 13.0, mul))
	fmt.Println(f(12.3, 13.0, div))
}
