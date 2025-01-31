package main

import (
	"github.com/gdamore/tcell/v2"
)

func getInput(screen tcell.Screen) {
	ev := screen.PollEvent()

	switch ev := ev.(type) {
	case *tcell.EventResize:
		(screen).Sync()
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyBackspace:
			fallthrough
		case tcell.KeyBackspace2:
			backspace()
		case tcell.KeyEscape:
			fallthrough
		case tcell.KeyCtrlC:
			closeDisplay(screen)
		default:
			addRune(ev.Rune())
		}
	}
}
