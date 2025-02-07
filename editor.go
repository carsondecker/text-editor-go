package main

import (
	"slices"
)

const init_gap_size = 10

type GapBuffer struct {
	text                    []rune
	gapStart                int
	gapEnd                  int
	cursorPosOnLine         int
	farthestCursorPosOnLine int
}

func newGapBuffer(text []rune) *GapBuffer {
	return &GapBuffer{
		text:                    text,
		gapStart:                0,
		gapEnd:                  0,
		cursorPosOnLine:         0,
		farthestCursorPosOnLine: 0,
	}
}

func (gb *GapBuffer) insert(r rune) {
	if gb.gapEnd-gb.gapStart <= 1 {
		gb.grow()
	}
	gb.text[gb.gapStart] = r
	gb.gapStart++
	gb.cursorPosOnLine = gb.gapStart - gb.getStartOfLine(gb.gapStart)
	gb.farthestCursorPosOnLine = gb.cursorPosOnLine
}

func (gb *GapBuffer) grow() {
	gb.text = slices.Concat(gb.text[:gb.gapStart], make([]rune, init_gap_size), gb.text[gb.gapStart:])
	gb.gapEnd += init_gap_size
}

func (gb *GapBuffer) backspace() {
	if gb.gapStart != 0 {
		gb.gapStart--
		gb.text[gb.gapStart] = rune(0)
		gb.cursorPosOnLine = gb.gapStart - gb.getStartOfLine(gb.gapStart)
		gb.farthestCursorPosOnLine = gb.cursorPosOnLine
	}
}

func (gb *GapBuffer) left() {
	if gb.gapStart != 0 {
		gb.gapStart--
		gb.gapEnd--
		gb.text[gb.gapEnd] = gb.text[gb.gapStart]
		gb.text[gb.gapStart] = rune(0)
		gb.cursorPosOnLine = gb.gapStart - gb.getStartOfLine(gb.gapStart)
		gb.farthestCursorPosOnLine = gb.cursorPosOnLine
	}
}

func (gb *GapBuffer) right() {
	if gb.gapEnd != len(gb.text) {
		gb.gapEnd++
		gb.text[gb.gapStart] = gb.text[gb.gapEnd-1]
		gb.text[gb.gapEnd-1] = rune(0)
		gb.gapStart++
		gb.cursorPosOnLine = gb.gapStart - gb.getStartOfLine(gb.gapStart)
		gb.farthestCursorPosOnLine = gb.cursorPosOnLine
	}
}

func (gb *GapBuffer) up() {
	if gb.getStartOfLine(gb.gapStart) == 0 {
		gb.moveGapToPos(0)
		return
	}
	startOfPrevLine := gb.getStartOfLine(gb.getStartOfLine(gb.gapStart) - 1)
	temp := gb.farthestCursorPosOnLine
	if gb.getEndOfLine(startOfPrevLine)-startOfPrevLine > gb.farthestCursorPosOnLine {
		gb.moveGapToPos(startOfPrevLine + gb.farthestCursorPosOnLine)
	} else {
		gb.moveGapToPos(gb.getEndOfLine(startOfPrevLine))
	}
	gb.farthestCursorPosOnLine = temp
}

func (gb *GapBuffer) down() {
	if gb.getEndOfLine(gb.gapStart)+(gb.gapEnd-gb.gapStart) == len(gb.text) {
		gb.moveGapToPos(gb.getEndOfLine(gb.gapStart))
		return
	}
	farthestPosOnLine := gb.farthestCursorPosOnLine
	gb.moveGapToPos(gb.getEndOfLine(gb.gapStart) + 1)
	if gb.gapStart+farthestPosOnLine >= gb.getEndOfLine(gb.gapStart) {
		gb.moveGapToPos(gb.getEndOfLine(gb.gapStart))
	} else {
		gb.moveGapToPos(gb.gapStart + farthestPosOnLine)
	}
	gb.farthestCursorPosOnLine = farthestPosOnLine
}

func (gb *GapBuffer) moveGapToPos(pos int) {
	for gb.gapStart != pos {
		if gb.gapStart < pos {
			gb.right()
		} else {
			gb.left()
		}
	}
}

func (gb *GapBuffer) getStartOfLine(pos int) int {
	curr := pos
	for curr != 0 && gb.text[curr-1] != '\r' {
		curr--
	}
	return curr
}

func (gb *GapBuffer) getEndOfLine(pos int) int {
	curr := pos
	if len(gb.text) <= 1 {
		return 0
	}
	if curr >= len(gb.text)-1 {
		return len(gb.text) - 1
	}

	if gb.text[curr] == '\r' && gb.text[curr+1] != '\r' {
		curr++
	}

	for curr != len(gb.text) && gb.text[curr] != '\r' {
		curr++
	}
	if curr >= gb.gapEnd {
		return curr - (gb.gapEnd - gb.gapStart)
	}
	return curr
}
