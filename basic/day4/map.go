package main

import "fmt"

func main() {
	//Reslicing
	mySlice := make([]int, 5, 10)
	mySlice[0] = 12
	mySlice[1] = 15
	mySlice[2] = 23
	fmt.Println("mySlice-", mySlice)
	reslice := mySlice[:2] //slice of slice (start till 1)
	reslice[0] = 334       //change in reslice will make change in slice also
	fmt.Println("reslice-", reslice)
	fmt.Println("mySlice-", mySlice)

	newslice := mySlice
	newslice[0] = 223
	fmt.Println("newslice-", newslice)
	fmt.Println("myslice-", mySlice) //changes in new slice effect original slice to avoid this we use copy func

	copyslice := make([]int, 5, 10) //its lenth and cap should be same as slice whos copy we r making
	num := copy(copyslice, mySlice)
	copyslice[0] = 989
	fmt.Println("number of elemets copied-", num)
	fmt.Println("copyslice-", copyslice)
	fmt.Println("mySlice-", mySlice)

	//maps
	myMap := make(map[string]int)

	myMap["mike"] = 23
	myMap["john"] = 45
	myMap["taylor"] = 89
	fmt.Println("map-", myMap)
	myMap["taylor"] = 45 //changing map values
	fmt.Println("after changeMap fn call myMap-", myMap)

	//delete in map
	delete(myMap, "mike")
	delete(myMap, "wyn")
	fmt.Println("after deleting from map", myMap)

	map1 := make(map[string]int)
	var map3 map[string]int
	fmt.Println("map1", map1)
	map1["n"] = 2
	//fmt.Println("map2", map2)
	//map2["t"] = 5
	map2 := map[string]int{}
	map2["n"] = 2
	fmt.Println("map2", map2)
	map3 =map[string]int{
		"p":4
	}
	fmt.Pritnln(map3)
}
