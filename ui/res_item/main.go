package main

import (
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

type ResourceItem struct {
	*walk.Composite
	resName  *walk.Label
	ip       *walk.LinkLabel
	protocal *walk.TextLabel
	ports    *walk.TextLabel
}

func NewResourceItem(form walk.Form) (*ResourceItem, error) {
	com, err := walk.NewCompositeWithStyle(form, win.WS_BORDER)
	//com.SetSize(walk.Size{100,100})
	//com, err := walk.NewComposite(form)

	if err != nil {
		return nil, err
	}
	r := &ResourceItem{Composite: com}
	layout := walk.NewVBoxLayout()
	r.SetLayout(layout)

	r.resName, err = walk.NewLabel(r)
	if err != nil {
		return nil, err
	}
	r.resName.SetText("资源名称")

	c, _ := walk.NewComposite(r)
	boxLayout := walk.NewHBoxLayout()
	boxLayout.SetMargins(walk.Margins{})
	c.SetLayout(boxLayout)
	r.protocal, _ = walk.NewTextLabel(c)
	r.ports, _ = walk.NewTextLabel(c)
	r.protocal.SetText("TCP")
	r.ports.SetText("Port: 8080")

	r.ip, err = walk.NewLinkLabel(r)
	if err != nil {
		return nil, err
	}
	r.ip.SetText("10.0.0.1")

	return r, nil
}

func main() {
	mw, err := walk.NewMainWindow()
	if err != nil {
		panic(err)
	}
	layout := walk.NewFlowLayout()
	//layout := walk.NewGridLayout()
	//layout.SetRowStretchFactor(200,10)
	mw.SetLayout(layout)
	//listBox, _ := walk.NewListBox(mw)
	//fmt.Println(listBox.Model())
	NewResourceItem(mw)
	NewResourceItem(mw)
	NewResourceItem(mw)
	NewResourceItem(mw)
	NewResourceItem(mw)
	NewResourceItem(mw)
	NewResourceItem(mw)
	NewResourceItem(mw)
	NewResourceItem(mw)
	NewResourceItem(mw)
	NewResourceItem(mw)
	NewResourceItem(mw)

	mw.Show()
	mw.Run()
}
