package cell

import (
	"testing"
	"tictac/mark"
)

var newCell = NewCell(mark.Not)

func TestNewCell(t *testing.T) {
	actual := mark.Not
	expected := newCell.mark
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}

func TestGetMark(t *testing.T) {
	actual := mark.Not
	expected := newCell.GetMark()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}

func TestSetMark(t *testing.T) {
	actual := mark.Cross
	newCell.SetMark(mark.Cross)
	expected := newCell.GetMark()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}