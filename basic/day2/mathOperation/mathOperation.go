package main

import "fmt"

func mathoperation(f1, f2 float32, f func(float32, float32) float32) float32 {
	return f(f1, f2)
}
func add(num1, num2 float32) float32 {
	return num1 + num2
}
func sub(num1, num2 float32) float32 {
	return num1 - num2
}
func mul(num1, num2 float32) float32 {
	return num1 * num2
}
func div(num1, num2 float32) float32 {
	return num1 / num2
}
func main() {
	fmt.Println(mathoperation(12.0, 13.0, add))
	fmt.Println(mathoperation(12.0, 13.0, sub))
	fmt.Println(mathoperation(12.0, 13.0, mul))
	fmt.Println(mathoperation(12.0, 13.0, div))
}
