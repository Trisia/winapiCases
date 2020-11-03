package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

func main() {

	width := 400
	height := 50

	mw, err := walk.NewMainWindow()
	if err != nil {
		panic(err)
	}
	// 隐藏 最小化和最大化按钮
	win.SetWindowLong(mw.Handle(), win.GWL_STYLE,
		win.GetWindowLong(mw.Handle(), win.GWL_STYLE) & ^win.WS_MINIMIZEBOX & ^win.WS_MAXIMIZEBOX)

	mw.SetTitle("测试")
	mw.SetWidth(width)
	mw.SetHeight(height)

	layout := walk.NewVBoxLayout()
	//fmt.Println(layout.LayoutBase.Spacing())
	//layout.SetSpacing(0)
	//layout.SetMargins(walk.Margins{20,0,0,20})
	mw.SetLayout(layout)

	mw.SetX((int(win.GetSystemMetrics(0)) - mw.Width()) / 2)
	mw.SetY((int(win.GetSystemMetrics(1)) - mw.Height()) / 2)
	mw.Show()

	//title, _ := walk.NewLabel(mw)
	//title.SetText("配置连接地址")
	//title.SetAlignment(walk.AlignHCenterVCenter)
	//f, err := walk.NewFont(
	//	"Simsum",
	//	50,
	//	0,
	//)
	//if err != nil {
	//	panic(err)
	//}
	//title.SetFont(f)

	ctn1, err := walk.NewComposite(mw)
	ctn1.SetLayout(walk.NewHBoxLayout())

	edit, err := walk.NewLineEdit(ctn1)
	if err != nil {
		panic(err)
	}
	edit.SetToolTipText("格式为 https://ip:port")
	// 设置placeholder
	edit.SetCueBanner("请输入服务器地址")

	btn, err := walk.NewPushButton(ctn1)
	if err != nil {
		panic(err)
	}
	btn.SetText("确定")

	progressBar, _ := walk.NewProgressBar(mw)
	// 设置进度条样式为无限滚动
	progressBar.SetMarqueeMode(true)
	progressBar.SetVisible(false)
	btn.Clicked().Attach(func() {
		// 显示进度条，隐藏其他模块
		progressBar.SetVisible(true)
		ctn1.SetVisible(false)
	})

	edit.EditingFinished().Attach(func() {
		// pare Enter
		fmt.Println("End")
	})
	//mw.Activating().Attach(func() {
	//	// 窗口聚焦
	//	fmt.Println("激活")
	//})

	mw.Run()
}
