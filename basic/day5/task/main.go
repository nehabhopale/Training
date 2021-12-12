package main

import (
	"fmt"
	ShapeInterface "task/ShapeInterface"
	circle "task/circle"
	rectangle "task/rectangle"
	square "task/square"
	triangle "task/triangle"
)

func main() {
	cir := circle.New(3.2)
	getDetails(cir)
	tri := triangle.New(3, 4, 5, 6)
	getDetails(tri)
	rect := rectangle.New(2.3, 1.2)
	getDetails(rect)
	sq := square.New(2.5)
	getDetails(sq)
	fmt.Println(cir.GetParameters())
	fmt.Println(tri.GetParameters())
	fmt.Println(rect.GetParameters())
	fmt.Println(sq.GetParameters())
	cir.SetParameters(6.7)
	tri.SetParameters(2, 3, 4, 5)
	rect.SetParameters(7, 8)
	sq.SetParameters(3.4)
	fmt.Println(cir.GetParameters())

}
func getDetails(shape ShapeInterface.ShapeInterface) {
	fmt.Println("area of a shape-", shape.Area())
	fmt.Println("perimeter of a shape-", shape.Perimeter())
}
