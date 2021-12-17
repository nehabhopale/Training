package game

import("tictac/board"
"tictac/resultanalyzer"
"tictac/player"
"tictac/result"
"fmt"

)

type Game struct{
	players []*player.Player
	currentPlayer *player.Player
	board   *board.Board
	resultanalyzer *resultanalyzer.Resultanalyzer
}
func NewGame(players []*player.Player,currentPlayer *player.Player,board *board.Board,resultanalyzer *resultanalyzer.Resultanalyzer)*Game{
	return &Game{
		players:players,
		currentPlayer:currentPlayer,
		board:board,
		resultanalyzer:resultanalyzer,
	}
}

func (g *Game) Play(){
	var row int
	var col int
	var status result.Result
	g.currentPlayer=g.players[0]
	
	for ok := true; ok; ok = (status==result.InProgress){
	Start:
		fmt.Printf("enter row  and col  in a less than %d and greater than equal to 0 for %s ",g.board.GetSize(),g.currentPlayer.GetName())
		_, err := fmt.Scanln(&row,&col)
		if err!=nil{
			fmt.Println(err)
			
		}
		
		if row<0 ||col<0{
			fmt.Println("row and column can't be negative")
			goto Start
		}

		fmt.Println(row)
		fmt.Println(col)
		ok:=g.board.Set(uint8(row),uint8(col),g.currentPlayer.GetMark())
		if !ok{
			fmt.Println("Your entry is invalid .Consider valid position")
			goto Start
		}
		g.board.PrintBoard()
		status=g.resultanalyzer.Analyze()
		
		if status==result.Winner{
			fmt.Println(g.currentPlayer.GetName(),"is winner")
				break
		} else if status==result.Draw{
			fmt.Println("game tie")
			break
		}
		
		if g.currentPlayer == g.players[0] {
					g.currentPlayer = g.players[1]
				} else {
					g.currentPlayer= g.players[0]
				}
		}
	

}

