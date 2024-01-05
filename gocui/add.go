package main

import (
	"go.wit.com/gui/widget"
)

var fakeStartWidth int = me.FakeW
var fakeStartHeight int = me.TabH + me.FramePadH
// setup fake labels for non-visible things off screen
func (n *node) setFake() {
	w := n.tk
	w.isFake = true

	n.gocuiSetWH(fakeStartWidth, fakeStartHeight)

	fakeStartHeight += w.gocuiSize.Height()
	// TODO: use the actual max hight of the terminal window
	if (fakeStartHeight > 24) {
		fakeStartHeight = me.TabH
		fakeStartWidth += me.FakeW
	}
	if (logInfo) {
		n.showView()
	}
}

// set the widget start width & height
func (n *node) addWidget() {
	nw := n.tk
	log(logInfo, "setStartWH() w.id =", n.WidgetId, "n.name", n.Name)
	switch n.WidgetType {
	case widget.Root:
		log(logInfo, "setStartWH() rootNode w.id =", n.WidgetId, "w.name", n.Name)
		nw.color = &colorRoot
		n.setFake()
		return
	case widget.Flag:
		nw.color = &colorFlag
		n.setFake()
		return
	case widget.Window:
		nw.frame = false
		nw.color = &colorWindow
		// redoWindows(0,0)
		return
	case widget.Tab:
		nw.color = &colorTab
		// redoWindows(0,0)
		return
	case widget.Button:
		nw.color = &colorButton
	case widget.Box:
		nw.color = &colorBox
		nw.isFake = true
		n.setFake()
		return
	case widget.Grid:
		nw.color = &colorGrid
		nw.isFake = true
		n.setFake()
		return
	case widget.Group:
		nw.color = &colorGroup
		nw.frame = false
		return
	case widget.Label:
		nw.color = &colorLabel
		nw.frame = false
		return
	default:
		/*
		if n.IsCurrent() {
			n.updateCurrent()
		}
		*/
	}
	n.showWidgetPlacement(logInfo, "addWidget()")
}
