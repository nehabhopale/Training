package cell

import ("tictac/mark"
"fmt")

type Cell struct{
	mark mark.Mark
}

func(c *Cell)GetMark() mark.Mark{
	return c.mark
}
func (c *Cell)SetMark(value mark.Mark)error{
	if value==mark.Empty{
	 	return fmt.Errorf("mark can't be empty")
	}
	c.mark=value
	return nil

}
func NewCell(mark mark.Mark) *Cell{
	return &Cell{
		mark:mark,
	}
}
