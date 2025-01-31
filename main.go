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

	gb := newGapBuffer()
	for {
		screen.Show()
		getInput(screen, gb)
		screen.Clear()
		draw(screen, gb.text)
	}
}
