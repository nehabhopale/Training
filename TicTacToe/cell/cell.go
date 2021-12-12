package cell

import ("tictac/mark"
"fmt")

type Cell struct{
	mark mark.Mark
}

func(c *Cell)GetMark() mark.Mark{
	return c.mark
}
func (c *Cell)SetMark(value mark.Mark){
	if value==" "{
	 	fmt.Println("mark can't be empty")
	}
	c.mark=value

}
func NewCell(mark mark.Mark) *Cell{
	return &Cell{
		mark:mark,
	}
}
