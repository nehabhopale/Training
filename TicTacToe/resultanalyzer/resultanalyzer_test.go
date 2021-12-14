package resultanalyzer

import (
	"testing"
	"tictac/board"
	"tictac/mark"
	"tictac/result"
)

var board1 *board.Board = board.NewBoard(3)

var r = NewAnalyzer(board1)

func TestCheckBoard(t *testing.T) {
	board1.Set(0, 0, mark.Cross)
	board1.Set(0, 1, mark.Cross)
	board1.Set(0, 2, mark.Cross)
	actual := result.Winner
	expected := r.Analyze()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}