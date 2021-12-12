package main
import "fmt"
// Solve this without running the code
func main() {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4)
	a := append(s, 5)
	fmt.Println(a) // O/P :- [1 2 3 4 5]____________________
	b := append(s, 6)
	fmt.Println(b) // O/P :- [1 2 3 4 6]
	fmt.Println(a) // O/P :-[1 2 3 4 6]
	fmt.Println(s) // O/P :- [1 2 3 4]
	// What is the expected output and why?
	//for a and b we are making the changes at the same address of s

}
