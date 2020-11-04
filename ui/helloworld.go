package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

func main() {
	//var inTE, outTe *walk.TextEdit
	//mw, err := walk.NewMainWindowWithName("HelloWorld")
	mw, _ := walk.NewMainWindow()

	mw.SetTitle("Hello")
	mw.SetSize(walk.Size{400, 200})
	//mw.SetWidth(600)
	//mw.SetHeight(400)
	mw.SetLayout(walk.NewVBoxLayout())

	mw.SetX((int(win.GetSystemMetrics(0)) - mw.Width()) / 2)
	mw.SetY((int(win.GetSystemMetrics(1)) - mw.Height()) / 2)
	mw.Show()

	txt, _ := walk.NewTextLabel(mw.Form())
	txt.SetText("Info")

	button, _ := walk.NewPushButton(mw.Form())
	button.SetText("CLOSE")
	button.Clicked().Attach(func() {
		button.SetEnabled(false)
		walk.App().Exit(0)
	})

	mw.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		fmt.Println("I am accepted Exist instruct（wait 5 second）", *canceled, reason)
	})

	dialog, _ := walk.NewDialogWithFixedSize(mw)
	dialog.SetLayout(walk.NewFlowLayout())
	dialog.SetClientSize(walk.Size{200, 200})

	btn2, _ := walk.NewPushButton(mw)
	btn2.SetText("CCCC")
	btn2.Clicked().Attach(func() {
		dialog.Run()
	})

	code := mw.Run()
	fmt.Println("Exist Code:", code)
	//go func() {
	//	code := mw.Run()
	//	fmt.Println("Exist Code:", code)
	//}()
}
