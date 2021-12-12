package main

import "fmt"

func main() {
	myArray := [4]int{1, 2, 3, 4}
	fmt.Println(arrUpdate(myArray)) //it will consider copy of myarray
	fmt.Println(myArray)            //original myarray will come
	myArray = arrUpdate(myArray)
	fmt.Println(myArray)

}
func arrUpdate(arr [4]int) [4]int {
	arr[0] = 3
	return (arr)
	//arr[1]=arr[0]+arr[4]

}
