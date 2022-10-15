package game

type Game struct {
	Board
}

func NewGame(board [][]bool) *Game {
	b := Board{board}
	g := Game{b}
	return &g
}

func (g *Game) NextGen() {

}

func (g *Game) GetBoard() [][]bool {
	return g.Board.board
}
