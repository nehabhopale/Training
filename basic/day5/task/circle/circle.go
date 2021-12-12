package circle

import (
	"math"
)

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c *circle) GetParameters() float64 {
	return c.radius
}

func (c *circle) SetParameters(newRadius float64) {
	c.radius = newRadius
}
func New(radius float64) *circle { //if u decide to return calue then return should be cir only not &cir
	return &circle{
		radius: radius,
	}
}
