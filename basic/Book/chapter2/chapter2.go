package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	validations()
	fmt.Println()
	printValues()
	fmt.Println()
	fileHandling()
	fmt.Println()

}

func validations() {
	if true {
		fmt.Println("true")
	}
	if false {
		fmt.Println("false")
	}
	if !false {
		fmt.Println("not true")
	}
	if 12 == 12 {
		fmt.Println("12==12")
	}
	if 12 != 12 {
		fmt.Println("this is not 12")
	}
}

func fileHandling() {
	fileInfo, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo)
}

func printValues() {
	for x := 1; x <= 3; x++ {
		fmt.Print(x)
	}
	fmt.Println()
	for x := 2; x <= 3; x++ {
		fmt.Print(x)
	}
	fmt.Println()
	for x := 1; x <= 3; x += 2 {
		fmt.Print(x)
	}
	fmt.Println()
	for x := 3; x >= 1; x-- {
		fmt.Print(x)
	}
	fmt.Println()
	for x := 1; x < 3; x++ {
		fmt.Print(x)
	}
	fmt.Println()
	for x := 1; x >= 3; x++ {
		fmt.Print(x)
	}

}
