package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//1) taking input using scanln
	fmt.Println("Enter Your First Name: ")
	var first string
	// Taking input from user
	fmt.Scanln(&first)
	//2) using scanf
	var name string
	fmt.Print("Enter your name: ")
	_, err := fmt.Scanf("%s\n", &name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hello", name)
	//3) using NewReader
	fmt.Print("Enter your hobby: ")
	reader := bufio.NewReader(os.Stdin)
	hobby, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s\n", hobby)
	//fmt.Println(hobby)
	//4)using NewScanner
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())
}
