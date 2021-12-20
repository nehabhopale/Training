package main

import(
"tictac/mark"
"tictac/player"
"tictac/board"
"tictac/resultanalyzer"
"tictac/game"
"fmt")

func main(){
	var name1 string
	var name2 string
	label1: fmt.Println("enter your name for player 1")
	fmt.Scanf("%s\n", &name1)
	if name1==""{
		fmt.Println("name can't be empty string")
		goto label1
	}
	label2: fmt.Println("enter your name for player 2")
	fmt.Scanln(&name2)
	if name2==""{
		fmt.Println("name can't be empty string")
		goto label2
	}

	playerA := player.NewPlayer(name1, mark.Not)
	playerB := player.NewPlayer(name2, mark.Cross)
	var players []*player.Player
	players = append(players, playerA)
	players = append(players, playerB)
	var size int
	start: fmt.Println("enter the size for board(size*size) ")
	_, err := fmt.Scanln(&size)
	if size<=2 {
		fmt.Println("board size should be grater than 2")
		goto start
	}
	if err!=nil {
			fmt.Println(err)
			
	}
	board1:=board.NewBoard(uint8(size))
	resultanalyzer1:=resultanalyzer.NewAnalyzer(board1)

	currentPlayer:=playerA
	game1:=game.NewGame(players,currentPlayer,board1,resultanalyzer1)
	game1.Play()
}