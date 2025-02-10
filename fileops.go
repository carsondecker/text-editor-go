package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func getPath() string {
	var path string
	fmt.Println("*** Welcome to this text editor! ***")
	fmt.Println("Disclaimer: Please do not edit important files with this, there are chances of crashes or loss of data.")
	fmt.Print("Please paste the file path of the file you would like to edit: ")
	fmt.Scanln(&path)
	return path
}

func readFile(path string) []rune {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return []rune{}
	}
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
