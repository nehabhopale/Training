package main

import (
	"fmt"
	"io/ioutil"

	"path/filepath"
)

func printDirectory(path string, depth int) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("error reading %s: %s\n", path, err.Error())
		return
	}

	for _, entry := range entries {

		for i := 0; i < depth; i++ {
			fmt.Print("\t")
		}

		if entry.IsDir() {
			fmt.Println("{" + entry.Name() + "}")
			printDirectory(filepath.Join(path, entry.Name()), depth+1)
		} else {

			fmt.Println(entry.Name())
		}
	}
}

func main() {
	path := "C:/Users/neha/Desktop/training/Book"

	printDirectory(path, 0)
}
