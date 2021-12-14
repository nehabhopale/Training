package board

import ("testing"
"tictac/mark"
)

var newBoard = NewBoard(4)

func TestGetSize( t *testing.T){
	actual:=uint8(4)
	expected:=newBoard.GetSize()
	if expected!=actual{
		t.Errorf("expected was :%v and actual is %v ",expected,actual)
	}
}
func TestGetAt(t *testing.T){
	actual :=mark.Empty
	expected:=newBoard.GetAt(0)
	if expected!=actual{
		t.Errorf("expected was :%v and actual is %v ",expected,actual)
	}
}
func TestGet(t *testing.T){
	actual :=mark.Empty
	expected:=newBoard.Get(0,0)
	if expected!=actual{
		t.Errorf("expected was :%v and actual is %v ",expected,actual)
	}
}
func TestSet(t *testing.T){
	actual:=mark.Not

	newBoard.Set(0,0,mark.Not)
	expected:=newBoard.Get(0,0)
	if expected!=actual{
		t.Errorf("expected was :%v and actual is %v ",expected,actual)
	}

}

func TestIsFull(t *testing.T) {
	actual := false
	expected := newBoard.IsFull()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}

func TestIsEmpty(t *testing.T) {
	actual := false
	expected := newBoard.IsEmpty()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}