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

func Test_cells_keep_alive_when_have_two_or_tree_neighbours(t *testing.T) {
	board := [][]bool{
		{false, false, false, false},
		{false, true, true, false},
		{false, true, true, false},
		{false, false, false, false},
	}
	g := NewGame(board)
	g.NextGen()
	g.NextGen()
	g.NextGen()
	result := g.GetBoard()
	if !(reflect.DeepEqual(board, result)) {
		t.Errorf("Expected %v; but got %v", board, result)
	}
}

func Test_cells_die_by_underpopulation_one_or_less_neighbours(t *testing.T) {
	board := [][]bool{
		{true, false, false, false},
		{false, true, false, false},
		{false, false, false, false},
		{false, false, true, false},
	}
	expected := [][]bool{
		{false, false, false, false},
		{false, false, false, false},
		{false, false, false, false},
		{false, false, false, false},
	}
	g := NewGame(board)
	g.NextGen()
	result := g.GetBoard()
	if !(reflect.DeepEqual(expected, result)) {
		t.Errorf("Expected %v; but got %v", board, result)
	}
}
