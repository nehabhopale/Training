package main

//file create and then read
// importing the packages
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile() {

	fmt.Println("Writing to a file in Go lang")
	//file creating
	file, err := os.Create("test.txt") //file creation

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer file.Close()

	len, err := file.WriteString("Welcome to swabhav-Training.")

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func ReadFile() {

	fmt.Println("Reading a file in Go lang")
	fileName := "test.txt"

	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Println("File Name:", fileName)
	fmt.Println("Size(in bytes):", len(data))
	fmt.Printf("Data: %s", data)

}

// main function
func main() {

	CreateFile()
	ReadFile()
}
