package main

func main() {
	path := getPath()
	gb := newGapBuffer(readFile(path))

	screen := createDisplay()
	// Close screen on error
	defer func() {
		maybePanic := recover()
		screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}()

	for {
		screen.Show()
		getInput(screen, path, gb)
		screen.Clear()
		draw(screen, gb.text)
	}
}
