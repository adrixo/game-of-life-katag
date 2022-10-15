package game

type game struct {
	board [][]bool
}

func NewGame(board [][]bool) *game {
	g := game{board}
	return &g
}

func (g *game) NextGen() {

}

func (g *game) GetBoard() [][]bool {
	return g.board
}
