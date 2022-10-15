package game

import "strconv"

type Board struct {
	board [][]bool
	m     map[string]bool
	width int
	high  int
}

func NewBoard(board [][]bool) *Board {
	m, h, w := createHM(board)
	b := Board{board: board, m: m, high: h, width: w}
	return &b
}

func (b *Board) GetBoard() [][]bool {
	var bb [b.high][b.width]bool
	for x := 0; x < b.width; x++ {
		for y := 0; y < b.high; y++ {
			sX := strconv.FormatInt(int64(x), 10)
			sY := strconv.FormatInt(int64(y), 10)
			bb = b.m[sX+sY]
		}
	}
	return bb[:][:]
}

func createHM(board [][]bool) (m map[string]bool, width int, high int) {
	width = len(board)
	high = len(board[])
	m = make(map[string]bool)
	for x := 0; x < width; x++ {
		for y := 0; y < high; y++ {
			sX := strconv.FormatInt(int64(x), 10)
			sY := strconv.FormatInt(int64(y), 10)
			m[sX+sY] = board[x][y]
		}
	}
	return m, width, high
}
