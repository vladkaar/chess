package chess

type Game struct {
	Name string
}

const (
	BK = '♔' // Black King
	BQ = '♕' // Black Queen
	BR = '♖' // Black Bishop
	BB = '♗' // Black Knight
	BN = '♘' // Black Rook
	BP = '♙' // Black Pawn
	WK = '♚' // White King
	WQ = '♛' // White Queen
	WR = '♜' // White Bishop
	WB = '♝' // White Knight
	WN = '♞' // White Rook
	WP = '♟' // White Pawn
)

var Field [8][8]int

//NewGame create work fields.
func NewGame() (instance Game) {
	instance.Name = "qw"
	createField()
	return
}

func createField() {
	// N := 8

	// field := make([][]int, N)
	for i := range Field {
		// field[i] = make([]int, N)
		for j := range Field[i] {
			Field[i][j] = 32
		}
	}
	// field[3][0] = WQ
	// field[3][1] = BQ
	// field[3][2] = BQ
	// field[3][3] = WQ

	Field[0][0] = BR
	Field[0][1] = BN
	Field[0][2] = BB
	Field[0][3] = BQ
	Field[0][4] = BK
	Field[0][5] = BB
	Field[0][6] = BB
	Field[0][7] = BR

	Field[1][0] = BP
	Field[1][1] = BP
	Field[1][2] = BP
	Field[1][3] = BP
	Field[1][4] = BP
	Field[1][5] = BP
	Field[1][6] = BP
	Field[1][7] = BP

	Field[6][0] = WP
	Field[6][1] = WP
	Field[6][2] = WP
	Field[6][3] = WP
	Field[6][4] = WP
	Field[6][5] = WP
	Field[6][6] = WP
	Field[6][7] = WP

	Field[7][0] = WR
	Field[7][1] = WN
	Field[7][2] = WB
	Field[7][3] = WQ
	Field[7][4] = WK
	Field[7][5] = WB
	Field[7][6] = WN
	Field[7][7] = WR

}
