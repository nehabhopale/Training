package main

import "os"

func main() {
	file, err := os.Create("file.txt")

	if err != nil {
		return
	}
	defer file.Close()

	file1, err1 := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		return
	}
	defer file1.Close()
	if _, err := file.WriteString("second line"); err != nil {
		return
	}

}
