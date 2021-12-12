package main

import "fmt"

func slice() {
	numbers := make([]float64, 3)
	numbers[0] = 19.7
	numbers[2] = 25.2
	for i, number := range numbers {
		fmt.Println(i, number)
	}
	var letters = []string{"a", "b", "c"}
	for i, letter := range letters {
		fmt.Println(i, letter)
	}
}
func myslice() {
	array := [5]string{"a", "b", "c", "d", "e"}
	slice := array[1:3]
	slice = append(slice, "x")
	slice = append(slice, "y", "z")
	for _, letter := range slice {
		fmt.Println(letter)
	}
}
func sum(numbers ...int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func main() {
	slice()
	fmt.Println(sum(7, 9))
	fmt.Println(sum(1, 2, 4))

}
