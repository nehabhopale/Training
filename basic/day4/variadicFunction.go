package main

import "fmt"

func main() {
	test("n", 10, 20)
	test("b", 12, 76, 5, 6)

}
func test(a string, b ...int) { //it will take 1 string and n number of int ,  there can be only one arg like b

	fmt.Println("a", a)
	fmt.Println("b", b)
}
