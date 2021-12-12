package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//writing a file

func main() {

	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("old falcon\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	val := "writing using ioutil\n"
	data := []byte(val)

	err1 := ioutil.WriteFile("data.txt", data, 0)

	if err1 != nil {
		log.Fatal(err)
	}

	fmt.Println("done")

	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString("second line"); err != nil {
		log.Fatal(err)
	}

	//Print the contents of the file
	data1, err22 := ioutil.ReadFile("data.txt")
	if err22 != nil {
		log.Fatal(err22)
	}
	fmt.Println(string(data1))

	//writing to file at diiff location
	file3, err3 := os.OpenFile("C:/Users/neha/Documents/FileHandling/test.txt", os.O_APPEND|os.O_WRONLY, 0777)
	if err3 != nil {
		log.Println(err)
	}
	defer file3.Close()
	if _, err := file3.WriteString("second line"); err != nil {
		log.Fatal(err)
	}

	//Print the contents of the file
	data11, err4 := ioutil.ReadFile("C:/Users/neha/Documents/FileHandling/test.txt")
	if err4 != nil {
		log.Fatal(err22)
	}
	fmt.Println(string(data11))

	byteSlice := make([]byte, 5)
	bytesRead, err9 := file3.Read(byteSlice)
	if err9 != nil {
		log.Fatal(err9)
	}
	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", byteSlice)

}
