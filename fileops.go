package main

import (
	"fmt"
	"log"
	"os"
)

func getFileFromPath(path string) []rune {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Could not read file.")
	}
	fmt.Print(string(data))
	return []rune(string(data))
}

func saveFile() {

}
