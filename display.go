package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func createDisplay() tcell.Screen {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	screen.SetStyle(defStyle)
	screen.Clear()

	return screen
}

func closeDisplay(screen tcell.Screen) {
	screen.Fini()
	os.Exit(0)
}

func draw(screen tcell.Screen, text []rune) {
	row, col := 0, 0
	for _, r := range text {
		if r == rune(0) {
			r = '_'
		}

		if r == '\n' || r == '\r' {
			col++
			row = 0
		} else {
			screen.SetContent(row, col, r, nil, tcell.StyleDefault)
			row++
		}
	}
}
