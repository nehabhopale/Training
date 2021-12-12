
package main

import (
	"fmt"
	"reflect"
)

type Shape interface {
	perimeter() int
}

type Square struct {
	side int
}

type Rectangle struct {
	length  int
	breadth int
}

func (s *Square) perimeter() int {
	return 4*s.side
}
func (r *Rectangle) perimeter() int {
	return (2*(r.length + r.breadth))
}

/*func Measure(s interface{}) { // Violates OCP because for new Shape eg Triangle  we need to add another case
	switch s.(type) {
	case Square:
		side := s.(Square).side
		fmt.Println("perimeter of square is ", 4*side)
	case Rectangle:
		l := s.(Rectangle).length
		b := s.(Rectangle).breadth
		fmt.Println("perimeter of rectangle is", 2*(l+b))
	}
}*/
func Measure(s Shape) {
	fmt.Printf("perimeter of %v is %v \n", reflect.TypeOf(s), s.perimeter())
}

func main() {
	s := &Square{side: 4}
	r := &Rectangle{length: 2, breadth: 3}
	Measure(s)
	Measure(r)
}