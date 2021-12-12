package main

import (
	"fmt"
	"sync"
)

type log struct {
}

var singleInstance *log
var lock = &sync.Mutex{}
var wg = sync.WaitGroup{}

func getInstance() *log {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {

			fmt.Println("Creting Single Instance Now")
			singleInstance = &log{}
		} else {
			fmt.Println("Single Instance 1")
		}
	}
	wg.Done()
	return singleInstance

}
func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {

		go getInstance()
	}
	wg.Wait()

}
