package main
import ("fmt"
"time"
"sync"
)
var wg=sync.WaitGroup{}
func main(){
	fmt.Println("start of main")
	wg.Add(2)
	go delayedIteration1()
	go delayedIteration2()
	wg.Wait()
	fmt.Println("end of main")
}
func delayedIteration1(){
	for i:=0;i<3;i++{
		time.Sleep(time.Second)
		fmt.Println(i)
	}
	wg.Done()

}
func delayedIteration2(){
	for i:=0;i<3;i++{
		time.Sleep(time.Second)
		fmt.Println(i)
	}
	hi()
	wg.Done()
}
func hi(){
	fmt.Println("hiiii neha")
}