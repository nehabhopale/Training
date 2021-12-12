package main
import "fmt"

type value struct{
	num int
}
//single resposibility

//func should do only one thing 

// violation
// //here square method is performing multiplicaation also and printing result also
// func (v value )square() {
// 	num1:=v.num*v.num
// 	fmt.Println(num1)
// }
func(v value) square()int{
	return v.num*v.num
}
func (v value) display(){
	fmt.Println(v.num*v.num)
}
func main(){
	v:=value{
		num:4,
	}
	fmt.Println(v.square())
	v.display()
}
