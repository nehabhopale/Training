package player

import (
	"testing"
	"tictac/mark"
)

var newPlayer = NewPlayer("abc", mark.Not)

func TestNewPlayer(t *testing.T) {
	actualName := "abc"
	expectedName := newPlayer.name
	actualMark := mark.Not
	expectedMark := newPlayer.mark

	if actualName != expectedName {
		t.Errorf("expected:%v and actual:%v", expectedName, actualName)
	} else if actualMark != expectedMark {
		t.Errorf("expected:%v and actual:%v", expectedMark, actualMark)
	}
}

func TestGetMark(t *testing.T) {
	actual := mark.Not
	expected := newPlayer.GetMark()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}

func TestGetName(t *testing.T) {
	actual := "abc"
	expected := newPlayer.GetName()
	if expected != actual {
		t.Errorf("expected:%v and actual:%v", expected, actual)
	}
}