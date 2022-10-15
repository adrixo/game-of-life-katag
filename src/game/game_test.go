package game

import (
	"reflect"
	"testing"
)

func Test_Keep_empty_board_when_input_empty_board(t *testing.T) {
	board := [][]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}
	g := NewGame(board)
	g.NextGen()
	result := g.GetBoard()
	if !(reflect.DeepEqual(board, result)) {
		t.Errorf("Expected %v; but got %v", board, result)
	}
}
