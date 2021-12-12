package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 5)
	fmt.Println("address of myslice", &mySlice)
	fmt.Printf("address of myslice %p \n", &mySlice)
	fmt.Println(&mySlice[0])
	fmt.Printf("address of slice using printf %p \n", &mySlice[1])
	mySlice[0] = 12
	mySlice[1] = 1
	mySlice[2] = 14
	fmt.Println(mySlice) //12 1 14 0 0
	copyMySlice := mySlice
	copyMySlice[2] = 3
	fmt.Printf("address of copymyslice %p \n", &copyMySlice)
	fmt.Println("copy of myslice", copyMySlice)
	fmt.Println(" myslice after update in its copy", mySlice)
	fmt.Println("len of mySlice", len(mySlice))
	fmt.Println("len of mySlice", cap(mySlice))

	mySlice = append(mySlice, 12) //create a new mySlice of double cap becaz we have appended when its cap is full
	fmt.Println(mySlice)          //12 1 3 0 0 12
	fmt.Println(len(mySlice))
	fmt.Println(len(mySlice))
	fmt.Printf("address after appending in slice whose cap is full %p", &mySlice[0])

	myNewSlice := make([]int, 5, 10)
	fmt.Println(myNewSlice)
	myNewSlice[0] = 23
	myNewSlice[1] = 2
	myNewSlice[3] = 14
	fmt.Println(myNewSlice)
	fmt.Printf("adress of first ele before append %p \n", &myNewSlice[0])
	myNewSlice = append(myNewSlice, 45)
	fmt.Printf("adress of first ele after append %p \n", &myNewSlice[0])

	slice := make([]string, 5, 5)
	slice[0] = "neha"
	slice[2] = "n"
	//slice of string pointers
	pointerSlice := make([]*string, 5, 5)
	pointerSlice[0] = &slice[0]
	fmt.Println("value at that pointer", *pointerSlice[0])
}
