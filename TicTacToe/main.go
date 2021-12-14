package main

import(
"tictac/mark"
"tictac/player"
"tictac/board"
"tictac/resultanalyzer"
"tictac/game")

func main(){
	playerA := player.NewPlayer("k", mark.Not)
	playerB := player.NewPlayer("l", mark.Cross)
	var players []*player.Player
	players = append(players, playerA)
	players = append(players, playerB)
	board1:=board.NewBoard(4)
	resultanalyzer1:=resultanalyzer.NewAnalyzer(board1)

	currentPlayer:=playerA
	game1:=game.NewGame(players,currentPlayer,board1,resultanalyzer1)
	game1.Play()
}