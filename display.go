package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

var runeArr = []rune{}

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

func addRune(r rune) {
	runeArr = append(runeArr, r)
}

func backspace() {
	runeArr = runeArr[:len(runeArr)-1]
}

func draw(screen tcell.Screen) {
	for i, r := range runeArr {
		screen.SetContent(i, 0, r, nil, tcell.StyleDefault)
	}
}
