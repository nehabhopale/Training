package resultanalyzer

import("tictac/result"
"tictac/mark"
"tictac/board"
//"fmt"

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
	
	if  r.checkRowsForResult() || r.checkColsForResult()||r.checkDiagonal1()||r.checkDiagonal2() {
		return result.Winner
	}else if r.board.IsFull() {
		return result.Draw
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
	
	var rowFirstEntry mark.Mark
	var i uint8
	size:=r.board.GetSize()
	for i=0;i<size*size;i+=size{
		
		var count uint8 =0
		rowFirstEntry=r.board.GetAt(i)
		if rowFirstEntry !=mark.Empty{
			//fmt.//Println("i",i)
			for j:=i+1;j<size*size;j++{
				//fmt.Println("j",j)
				if r.board.GetAt(j)==rowFirstEntry{
					count++
				}else{
					break
				}
				
			}
			//fmt.Println("outside loop ",count)
			if count==(size-1){
				//fmt.Println("inside row",count)
				return true 
			}
		}else{
			continue
		}
	}
	return false
}
func(r Resultanalyzer) checkColsForResult()bool{
	
	var colFirstEntry mark.Mark
	var i uint8
	size:=r.board.GetSize()
	for i=0;i<size;i++{
		colFirstEntry=r.board.GetAt(i)
		if colFirstEntry !=mark.Empty{
			var count uint8 =0
			for j:=i+size;j<size*size;j+=size{
				if r.board.GetAt(j)==colFirstEntry{
					count++
				}else{
					break
				}
			}
			if count==size-1{
				//fmt.Println("inside row",count)
				return true 
			}
		}
	}
	return false

	
}
func(r Resultanalyzer) checkDiagonal1() bool{
	diag1Entry:=r.board.GetAt(0)
	size :=r.board.GetSize()
	var c uint8 =size+1
	var i uint8
	var count uint8 =0
	for i = 1; i <r.board.GetSize(); i++ {
		if diag1Entry== r.board.GetAt(i*c) && (diag1Entry!=mark.Empty){
			count++
		}
	}
	// for i:=size+1;i<size*size;i+=(size+1){
	// 		if (r.board.GetAt(i)==diag1Entry )&& (diag1Entry!=mark.Empty) {
	// 			count++
	// 		}
	// }
	if count==size-1{
		//fmt.Println("inside diag1",count)
			return true 
	}
	return false 

}
func(r Resultanalyzer) checkDiagonal2() bool{
	size :=r.board.GetSize()
	var i uint8
	diag2Entry:=r.board.GetAt(size-1)
	var count uint8 =0
	// for i:=size+1;i<size*size;i+=(size-1){
	// 		fmt.Println("i in diag2",i)
	// 		if (r.board.GetAt(i)==diag2Entry )&& (diag2Entry!=mark.Empty){
	// 			count++
	// 		}
	// }
	var c uint8 =size-1
	for i = 2; i <=r.board.GetSize(); i++ {
		if diag2Entry== r.board.GetAt(i*c) && (diag2Entry!=mark.Empty){
			count++
		}
	}

	if count==size-1{
		//fmt.Println("inside diag2",count)
			return true 
	}
	return false 
}




