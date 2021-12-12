package main

import "fmt"

var i int

func main() {
	defer f1()
	fmt.Println("hey i m at start")
	defer f2() // defer f2( will be executed first then der1)
	var i = 10
	defer f3(i) // giving i =10 ?????????
	i = 20
	// if i == 0 {
	// 	panic("i m unable to guess ")
	// }
	if i == 20 {
		panic("i m unable to guess ")
	}
	fmt.Println("value of i inside main", i)
	fmt.Println("   ")
	defer f4(&i) //20
	//panic search for defer.
	defer func() {
		details := recover()
		fmt.Println("inside defer using recover", details) //will print panic statement if panic nott executed then with nil

	}()
	//panic("i m unable to guess ")
}
func f1() {
	fmt.Println("inside defer f1")
	defer f2()                                        //defer inside defer
	fmt.Println(" inside defer f1 bt after defer f2") //it will print everyting from f1 then jump to f2
}
func f2() {
	fmt.Println("i m from f2 inside defer f1")
}
func f3(i int) {
	fmt.Println("value of i is", i)

}
func f4(i *int) {
	fmt.Println("value of i by considering pointer to i", *i)
}
