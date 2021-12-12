package board

import ("tictac/cell"
"tictac/mark"
"fmt"
)

type Board struct{
	cells [][]*cell.Cell

}
func NewBoard() *Board {
	cells := make([][] *cell.Cell,3)
	for i := 0; i < 3; i++ {
		cells[i]=make([]*cell.Cell,3)
		for j:=0;j<3;j++{
			cells[i][j] = cell.NewCell(mark.Empty)	
		}
		
	}
	return &Board{
		cells: cells,	
	}
}
func (b *Board)PrintBoard(){
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			fmt.Print(*b.cells[i][j])
		}
		fmt.Println()
	}
}
func(b *Board)IsEmpty() bool{
	cellsEmpty := true
	for  i := 0; i < 3; i++ {
		for j:=0;j<3;j++{
			if b.cells[i][j].GetMark() != mark.Empty{
				cellsEmpty = false
			}
		}
	}

	return cellsEmpty
}


func(b *Board)IsFull() bool {
	cellsFull:=true
	for i := 0; i < 3; i++ {
		for j:=0;j<3;j++{
			if (b.cells[i][j].GetMark()== mark.Empty){
				cellsFull=false
			}
		}
	}

	return cellsFull
}

func(b *Board)Set(row int, col int , mark1 mark.Mark) {
	if b.cells[row][col].GetMark() !=mark.Empty{
		fmt.Println("its an occupied position give a try in next move")
	}
	if   b.cells[row][col].GetMark() ==mark.Empty && checkLocation(row, col){
		b.cells[row][col].SetMark(mark1)
	}
}

func (b *Board)Get(row int, col int) mark.Mark {
	
	return b.cells[row][col].GetMark()
}

func checkLocation(row int, col int ) bool {
	if (row > 2 || row < 0){
		fmt.Println("please enter valid rows ")
		return false

	}

	if (col > 2 || col < 0){
		fmt.Println("please enter valid columns ")
		return false
		
	}
		
	return true

}
