package board

import ("tictac/cell"
"tictac/mark"
"fmt"
)

type Board struct{
	cells []*cell.Cell
	size uint8

}
func NewBoard(size uint8) *Board {
	cells := make([]*cell.Cell, size*size)
	for i := 0; i < int(size)*int(size); i++ {
		cells[i] = cell.NewCell(mark.Empty)
	}
	return &Board{
		cells: cells,
		size:  size,
	}
}
func (b *Board)GetSize()uint8{
	return b.size
}
func (b *Board)GetAt(pos uint8)mark.Mark{
	return b.cells[pos].GetMark()
}
func (b *Board)PrintBoard(){
	s := int(b.size)
	count := 1
	for i := 0; i < s*s; i++ {
		fmt.Printf("%v", *b.cells[i])
		count++
		if count > s {
			fmt.Println()
			count = 1
		}
	}
}
func(b *Board)IsEmpty() bool{
	cellsEmpty := true
	var i uint8
	for  i = 0; i < b.size*b.size; i++ {
			if b.cells[i].GetMark() != mark.Empty{
				cellsEmpty = false
			}
	}

	return cellsEmpty
}


func(b *Board)IsFull() bool {
	cellsFull:=true
	var i uint8
	for i = 0; i < b.size*b.size; i++ {	
			if (b.cells[i].GetMark()== mark.Empty){
				cellsFull=false
			}
	}

	return cellsFull
}

func(b *Board)Set(row uint8, col uint8 , mark1 mark.Mark) bool {
	cellLocation:=getLocation(row,col,b.size)
	fmt.Println("location is ",cellLocation)
	if row > (b.size-1) || col > (b.size-1) || b.cells[cellLocation].GetMark() != mark.Empty {
		return false
	}
	b.cells[cellLocation].SetMark(mark1)
	return true
	// if b.cells[cellLocation].GetMark() !=mark.Empty{
	// 	fmt.Println("its an occupied position give a try in next move.Its a turn of  ",b.cells[cellLocation].GetMark())
	// }
	// if  b.cells[cellLocation].GetMark() ==mark.Empty && getLocation(row, col,b.size)!=90{
		
	// 	b.cells[cellLocation].SetMark(mark1)
	// }
}

func (b *Board)Get(row uint8, col uint8) mark.Mark {
	cellLocation:=getLocation(row,col,b.size)
	return b.cells[cellLocation].GetMark()
}

func getLocation(row uint8, col uint8,n uint8 ) uint8 {
	// if (row > (n-1)){
	// 	fmt.Println("please enter valid rows ")
	// 	return 90

	// }

	// if (col > (n-1) ){
	// 	fmt.Println("please enter valid columns ")
	// 	return 90
		
	// }
		
	return row*n+col

}
