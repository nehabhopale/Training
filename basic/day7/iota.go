package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
const (
	_           = iota // ignore first value by assigning to blank identifier
	KB int = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {
	fmt.Print("Enter your filepath: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//fmt.Println(scanner.Text())
	filepath := scanner.Text()
	fmt.Print("Enter the size format you want: ")
	scaning := bufio.NewScanner(os.Stdin)
	scaning.Scan()
	sizeFormat := scaning.Text()
	showSize(filepath,sizeFormat)
	
	
}
func showSize(filepath string,sizeFormat string){
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := float64(fileInfo.Size())
	switch sizeFormat {
	case "KB":
		fileSize = fileSize / float64(KB)
	case "MB":
		fileSize = fileSize / float64(MB)
	case "GB":
		fileSize = fileSize / float64(GB)
	case "TB":
		fileSize = fileSize / float64(TB)
	}
	fmt.Println("file path", filepath, " is ", fileSize, " ", sizeFormat)
}
