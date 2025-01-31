package main

func main() {
	screen := createDisplay()
	for {
		screen.Show()
		getInput(screen)
		screen.Clear()
		draw(screen)
	}
}
