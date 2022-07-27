package main

import (
	// _ "github.com/vladkaar/chess/internal/gui"

	"fmt"
	"os"

	"github.com/vladkaar/chess/internal/chess"
	"github.com/vladkaar/chess/internal/terminalgui"
)

func main() {
	testGame()
	// checkChar()
}

func testGame() {
	fmt.Println("Start chess:")
	game := chess.NewGame()
	fmt.Println(game.Name)
	game.Name = "ds"
	fmt.Println(game.Name)
	// gui.StartGui()
	terminalgui.TermGui()
}

func checkChar() {
	var data string
	start := 0
	finish := (1 << 14) - (1 << 11)
	for ch := start; ch < finish; ch++ {
		if ch%32 == 0 {
			data += "\n"
			data += fmt.Sprintf("%x\t%d\t", ch, ch)
		}
		data += fmt.Sprintf("%c \t", ch)
	}
	write := false
	// write := true
	if write {
		file, err := os.Create("hello.txt")
		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
		defer file.Close()
		file.Write([]byte(data))

	} else {
		fmt.Println(data)
	}

	//♔ = false

	fmt.Println("Done.")
}
