package triangle

type triangle struct {
	side1, base, side2, height float64
}

func (t triangle) Area() float64 {
	return (((t.base) * (t.height)) / 2)
}
func (t triangle) Perimeter() float64 {
	return t.side1 + t.side2 + t.base
}
func (t *triangle) GetParameters() (float64, float64, float64, float64) {
	return t.side1, t.side2, t.base, t.height
}

func (t *triangle) SetParameters(newside1 float64, newbase float64, newside2 float64, newheight float64) {
	t.side1 = newside1
	t.side2 = newside2
	t.height = newheight
	t.base = newbase

}
func New(side1, side2, base, height float64) *triangle {
	return &triangle{
		side1:  side1,
		base:   base,
		side2:  side2,
		height: height,
	}
}
