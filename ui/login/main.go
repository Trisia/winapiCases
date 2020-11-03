package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

func main() {
	mw, err := walk.NewMainWindow()
	if err != nil {
		panic(err)
	}

	mw.SetTitle("配置连接地址")
	mw.SetWidth(400)
	mw.SetHeight(100)
	mw.SetLayout(walk.NewHBoxLayout())

	mw.SetX((int(win.GetSystemMetrics(0)) - mw.Width()) / 2)
	mw.SetY((int(win.GetSystemMetrics(1)) - mw.Height()) / 2)
	mw.Show()

	ctn1, err := walk.NewComposite(mw)
	ctn1.SetLayout(walk.NewHBoxLayout())

	edit, err := walk.NewTextEdit(ctn1)
	if err != nil {
		panic(err)
	}
	edit.SetCompactHeight(true)

	edit.KeyDown().Attach(func(key walk.Key) {
		if key == walk.KeyExecute {
			fmt.Println(edit.Text())
		}
	})
	btn, err := walk.NewPushButton(ctn1)
	if err != nil {
		panic(err)
	}
	btn.SetText("确定")

	mw.Run()
}
