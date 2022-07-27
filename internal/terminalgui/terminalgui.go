package terminalgui

import (
	"fmt"

	ct "github.com/daviddengcn/go-colortext"
	"github.com/vladkaar/chess/internal/chess"
)

const (
	tr = iota
	one
)

func TermGui() {
	fmt.Println(tr, one)
	fmt.Printf("%d %d %d %d %d %d\n", chess.BQ, chess.BK, chess.BR, chess.BB, chess.BN, chess.BP)
	fmt.Printf("%x %x %x %x %x %x\n", chess.BQ, chess.BK, chess.BR, chess.BB, chess.BN, chess.BP)
	fmt.Printf("%c %c %c %c %c %c\n", chess.BQ, chess.BK, chess.BR, chess.BB, chess.BN, chess.BP)
	fmt.Printf("%d %d %d %d %d %d\n", chess.WQ, chess.WK, chess.WR, chess.WB, chess.WN, chess.WP)
	fmt.Printf("%x %x %x %x %x %x\n", chess.WQ, chess.WK, chess.WR, chess.WB, chess.WN, chess.WP)
	fmt.Printf("%c %c %c %c %c %c\n", chess.WQ, chess.WK, chess.WR, chess.WB, chess.WN, chess.WP)
	ct.Foreground(ct.Green, false)
	fmt.Println("Green text starts here...")
	ct.ChangeColor(ct.Black, true, ct.White, false)
	fmt.Printf("%c %c %c %c %c %c .", chess.WQ, chess.WK, chess.WR, chess.WB, chess.WN, chess.WP)
	ct.ResetColor()
	// time.Sleep(100)
	fmt.Println()
	fmt.Printf("%s", "addada")

	ct.ChangeColor(ct.Black, false, ct.White, false)
	fmt.Printf("%c ", chess.WQ)
	ct.ChangeColor(ct.White, false, ct.Black, false)
	fmt.Printf("%c ", chess.WQ)
	ct.ResetColor()
	fmt.Printf("%d\n", ' ')
	showField()

}

func showField() {

	field := chess.Field

	for i := range field {
		ct.Foreground(ct.White, false)
		fmt.Printf("%d ┃ ", 8-i)
		for j := range field[i] {
			if (i+j)%2 == 0 {
				ct.Background(ct.Black, true)
			} else {
				ct.Background(ct.Black, false)
			}
			// if field[i][j] < 9818 {
			// 	ct.Foreground(ct.Black, false)
			// } else {
			// 	ct.Foreground(ct.White, false)
			// }
			// fmt.Printf("%c ", WQ)
			fmt.Printf("%c ", field[i][j])
		}
		ct.ResetColor()
		fmt.Println()
	}
	fmt.Println("━━╋━━━━━━━━━━━━━━━━━")
	fmt.Println("  ┃ a b c d e f g h")
	/*
		for i := 0; i < N; i++ {
			ct.ChangeColor(ct.Black, false, ct.White, false)
			fmt.Printf("%c ", WQ)
			fmt.Println()
		}
	*/

}
