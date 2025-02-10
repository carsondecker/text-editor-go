package main

import (
	"fmt"
	"log"
	"os"
)

func getPath() string {
	// implement later
	return ""
}

func readFile(path string) []rune {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Panic("Could not read file.")
	}
	fmt.Print(string(data))
	return []rune(string(data))
}

func (gb *GapBuffer) saveFile(path string) {
	textOnly := make([]rune, len(gb.text[:gb.gapStart])+len(gb.text[gb.gapEnd:]))
	copy(textOnly, gb.text[:gb.gapStart])
	copy(textOnly[len(gb.text[:gb.gapStart]):], gb.text[gb.gapEnd:])
	err := os.WriteFile(path, []byte(string(textOnly)), 0644)
	if err != nil {
		log.Panic("Could not save file.")
	}
}
