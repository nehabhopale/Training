package main

import "fmt"

func main() {
	map1()
	fmt.Println(" ")
	slice()
}
func map1() {
	jewelry := make(map[string]float64)
	jewelry["necklace"] = 89.99
	jewelry["earings"] = 79.99
	clothing := map[string]float64{"pants": 59.99, "shirt": 39.99}
	fmt.Println("Earrings:", jewelry["earings"])
	fmt.Println("Earrings:", jewelry["necklace"])
	fmt.Println("Earrings:", clothing["shirt"])
	fmt.Println("Earrings:", clothing["pants"])
}
func slice() {
	data := []string{"a", "c", "e", "a", "e"}
	counts := make(map[string]int)
	for _, item := range data {
		counts[item]++
	}
	letters := []string{"a", "b", "c", "d", "e"}
	for _, letter := range letters {
		count, ok := counts[letter]
		if !ok {
			fmt.Printf("%s:not found\n", letter)
		} else {
			fmt.Printf("%s:%d\n", letter, count)
		}
	}
}
func travelmap() {
	ranks := map[string]int{"bronze": 3, "silver": 2, "gold": 1}
	for medal, rank := range ranks {
		fmt.Printf("The %s medal's rank id %d\n", medal, rank)
	}

}
