package main

import (
	"github.com/gdamore/tcell/v2"
)

func getInput(screen tcell.Screen, path string, gb *GapBuffer) {
	ev := screen.PollEvent()

	switch ev := ev.(type) {
	case *tcell.EventResize:
		(screen).Sync()
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyBackspace:
			fallthrough
		case tcell.KeyBackspace2:
			gb.backspace()
		case tcell.KeyLeft:
			gb.left()
		case tcell.KeyRight:
			gb.right()
		case tcell.KeyUp:
			gb.up()
		case tcell.KeyDown:
			gb.down()
		case tcell.KeyEscape:
			fallthrough
		case tcell.KeyCtrlC:
			closeDisplay(screen)
		case tcell.KeyCtrlS:
			gb.saveFile(path)
		default:
			gb.insert(ev.Rune())
		}
	}
}
