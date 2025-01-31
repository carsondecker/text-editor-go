package main

import "slices"

const init_gap_size = 10

type GapBuffer struct {
	text     []rune
	gapStart int
	gapEnd   int
}

func newGapBuffer() *GapBuffer {
	return &GapBuffer{
		text:     []rune{},
		gapStart: 0,
		gapEnd:   0,
	}
}

func (gb *GapBuffer) insert(r rune) {
	if gb.gapEnd-gb.gapStart <= 1 {
		gb.grow()
	}
	gb.text[gb.gapStart] = r
	gb.gapStart++
}

func (gb *GapBuffer) grow() {
	gb.text = slices.Concat(gb.text[:gb.gapStart], make([]rune, init_gap_size), gb.text[gb.gapStart:])
	gb.gapEnd += init_gap_size
}

func (gb *GapBuffer) backspace() {
	if gb.gapStart != 0 {
		gb.gapStart--
		gb.text[gb.gapStart] = rune(0)
	}
}

func (gb *GapBuffer) left() {
	if gb.gapStart != 0 {
		gb.gapStart--
		gb.text[gb.gapEnd-1] = gb.text[gb.gapStart]
		gb.text[gb.gapStart] = rune(0)
		gb.gapEnd--
	}
}

func (gb *GapBuffer) right() {
	if gb.gapEnd != len(gb.text) {
		gb.gapEnd++
		gb.text[gb.gapStart] = gb.text[gb.gapEnd-1]
		gb.text[gb.gapEnd-1] = rune(0)
		gb.gapStart++
	}
}
