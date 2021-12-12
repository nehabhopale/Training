package main

import (
	"fmt"
	"time"
)

func repeat(s string) {
	for i := 0; i < 25; i++ {
		fmt.Print(s)

	}
}
func odd(channel chan int) {
	channel <- 1
	channel <- 3
}
func even(channel chan int) {
	channel <- 2
	channel <- 4
}

func main() {
	go repeat("x")
	go repeat("y")
	time.Sleep(time.Second)
	channelA := make(chan int)
	channelB := make(chan int)
	go odd(channelA)
	go even(channelB)
	fmt.Println(<-channelA)
	fmt.Println(<-channelA)
	fmt.Println(<-channelB)
	fmt.Println(<-channelB)
}
