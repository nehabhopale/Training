package rectangle

type rectangle struct {
	length, breadth float64
}

func (r rectangle) Area() float64 {
	return r.length * r.breadth
}
func (r rectangle) Perimeter() float64 {
	return 2*r.length + 2*r.breadth
}
func (r *rectangle) GetParameters() (float64, float64) {
	return r.length, r.breadth
}

func (r *rectangle) SetParameters(newlength float64, newBreadth float64) {
	r.length = newlength
	r.breadth = newBreadth

}
func New(length, breadth float64) *rectangle {
	return &rectangle{
		length:  length,
		breadth: breadth,
	}
}
