package main

import "fmt"

func arr() {
	var numbers [3]int
	numbers[0] = 42
	numbers[2] = 108
	var letters = [3]string{"a", "b", "c"}
	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(letters[2])
	fmt.Println(letters[0])
	fmt.Println(letters[1])

}
func print() {
	numbers := [6]int{3, 6, -2, 10, 23, 12}
	for i, number := range numbers {
		if number >= 10 && number <= 20 {
			fmt.Println(i, number)
		}
	}
}
func main() {
	arr()
	print()
}
