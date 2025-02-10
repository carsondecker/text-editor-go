package main

func main() {
	screen := createDisplay()
	// Close screen on error
	defer func() {
		maybePanic := recover()
		screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}()

	path := getPath()
	gb := newGapBuffer(readFile(path))
	for {
		screen.Show()
		getInput(screen, path, gb)
		screen.Clear()
		draw(screen, gb.text)
	}
}
