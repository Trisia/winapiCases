package main

import "github.com/lxn/walk"

func main() {
	wmain, err := walk.NewMainWindow()
	if err != nil {
		panic(err)
	}
	wmain.SetLayout(walk.NewHBoxLayout())
	wmain.Show()
	wmain.Run()
}
