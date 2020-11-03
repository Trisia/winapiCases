package main

import (
	"fmt"
	"github.com/lxn/walk"
)

func main() {
	mw, err := walk.NewMainWindow()
	if err != nil {
		panic(err)
	}

	icon, err := walk.Resources.Icon("../img/logo.ico")
	if err != nil {
		panic(err)
	}

	ni, err := walk.NewNotifyIcon(mw)
	if err != nil {
		panic(err)
	}
	// 退出时清除资源
	defer ni.Dispose()

	if err := ni.SetIcon(icon); err != nil {
		panic(err)
	}
	// 鼠标悬浮形式文字
	if err := ni.SetToolTip("测试托盘点击展开"); err != nil {
		panic(err)
	}

	ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}
		if err := ni.ShowCustom("Walk 任务栏通知标题", "Walk任务栏通知内容", icon); err != nil {
			panic(err)
		}
	})
	existAction := walk.NewAction()
	if err := existAction.SetText("右键菜单"); err != nil {
		panic(err)
	}
	existAction.Triggered().Attach(func() {
		walk.App().Exit(0)
	})
	if err := ni.ContextMenu().Actions().Add(existAction); err != nil {
		panic(err)
	}

	if err := ni.SetVisible(true); err != nil {
		panic(err)
	}

	if err := ni.ShowInfo("Walk NotifyIcno Example", "点击重试"); err != nil {
		panic(err)
	}
	mw.Run()

	fmt.Println("End")
}
