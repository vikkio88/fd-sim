package main

import "fdsim/ui"

func main() {

	a := ui.NewApp()
	a.Run()

	a.Cleanup()
}
