package main
 
import (
    "fmt"
)
//inheritance by structure embedding
type A struct{
	name string			//here name is encapsulated in A struct 
}
type B struct{
	id string  // here i is small kept to achieve the abstraction
	A
}
func (b B)details(){
	fmt.Println(b.id)
	fmt.Println(b.A.name) //we exported Println from fmt package

}
func main(){
	A:=A{"neha"}
	B:=B{"1"}
	B.details()
	rectangle := Rectangle{
     
        length: 5.0,
        width: 2.5,
    }
     

    square := Square{
     
        side: 1.0,
    }
     
    
    fmt.Println( rectangle.Area())
    fmt.Println(square.Area())
}

//polymorphism by interfaces
//polymorphism is having many forms

 //using figure interface we can execute area method on diff forms like rec sq etc
type Figure interface{
 
    Area() float64
}
 
type Rectangle struct{
     
    length float64
    width float64
}
type Square struct{
     
   
    side float64
}

func (rect Rectangle) Area() float64{
 
    area := rect.length * rect.width
    return area
}
func (sq Square) Area() float64{
	 
    area := sq.side * sq.side
    return area
}
