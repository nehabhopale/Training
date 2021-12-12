package resultanalyzer

import("tictac/result"
"tictac/mark"
"tictac/board"
"fmt"
)

type Resultanalyzer struct{
	board *board.Board

}
func NewAnalyzer(board *board.Board) *Resultanalyzer{
	return &Resultanalyzer{
		board:board,

	}
}
func( r *Resultanalyzer ) Analyze() result.Result{
	
	if r.board.IsFull(){
		return result.Draw
	}else if r.checkRowsForResult() || r.checkColsForResult()||r.checkDiagonals(){
		return result.Winner
	}
	return result.InProgress
	
}

// func( r Resultanalyzer ) HasResult() bool{
// 	if r.checkRowsForResult(){
// 		return true
// 	}
// 	if r.checkColsForResult(){
// 		return true
// 	}

// 	if r.checkDiagonals(){
// 		return true
// 	}
// 	return false

// }
func(r *Resultanalyzer) checkRowsForResult()bool{
	
	var rowEntry mark.Mark
	for i:=0;i<3;i++{
		rowEntry=r.board.Get(i,0)
		if rowEntry!=mark.Empty{
			if(r.board.Get(i,0)==r.board.Get(i,1))&&(r.board.Get(i,1)==r.board.Get(i,2))&&(r.board.Get(i,0)==rowEntry){
				fmt.Println("rows",r.board.Get(i,0))
				return true
			} 
		}
	}
	return false
}
func(r Resultanalyzer) checkColsForResult()bool{
	
	var colEntry mark.Mark
	for i:=0;i<3;i++{
		colEntry=r.board.Get(0,i)
		if colEntry!=mark.Empty{
			if(r.board.Get(0,i)==r.board.Get(1,i))&&(r.board.Get(1,i)==r.board.Get(2,i))&&(r.board.Get(0,i)==colEntry){
				fmt.Println("cols",r.board.Get(0,i))
				return true
			} 
		}
	}
	return false

	
}
func(r Resultanalyzer) checkDiagonals() bool{

	
	diag1:=r.board.Get(0,0)
	if diag1!=mark.Empty{
		if (r.board.Get(0, 0) == r.board.Get(1, 1)) && (r.board.Get(1, 1) == r.board.Get(2, 2) )&& (r.board.Get(0, 0) ==diag1){
			fmt.Println("diag1",r.board.Get(0,0))	
			return true
		}
	}
	diag2:=r.board.Get(0,2)
	if diag2!=mark.Empty{
		if (r.board.Get(2, 0) == r.board.Get(1, 1)) &&(r.board.Get(1, 1) == r.board.Get(0, 2) )&& (r.board.Get(2, 0) ==diag2){
			fmt.Println("diag2",r.board.Get(2,0))
			return true
		}
	}

	return false
}




