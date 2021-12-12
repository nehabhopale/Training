package main

import (
	"fmt"
	"log"
)

func find(item string, slice []string) bool {
	for _, sliceItem := range slice {
		if item == sliceItem {
			return true
		}

	}
	return false
}

type Refrigerator []string

func (r Refrigerator) Open() {
	fmt.Println("opening refrigerator")
}
func (r Refrigerator) Close() {
	fmt.Println("closing refrigerator")
}
func main() {
	fridge := Refrigerator{"Milk", "pizza"}
	for _, food := range []string{"Milk", "banana"} {
		err := fridge.FindFood(food)
		if err != nil {
			log.Fatal(err)
		}
	}
	snack()
}
func snack() {
	defer fmt.Println("closing refrigerator")
	fmt.Println("opening refrigerator")
	panic("refrigerator is empty")
}
func (r Refrigerator) FindFood(food string) error {
	r.Open()
	defer r.Close()
	if find(food, r) {
		fmt.Println("found", food)
	} else {
		return fmt.Errorf("%s not found", food)
	}
	return nil
}
