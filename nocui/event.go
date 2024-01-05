package main

import (
	"go.wit.com/gui/widget"
)

func (n *node) doWidgetClick() {
	switch n.WidgetType {
	case widget.Root:
		// THIS IS THE BEGINING OF THE LAYOUT
		// rootNode.nextW = 0
		// rootNode.nextH = 0
		// rootNode.redoTabs(true)
	case widget.Flag:
		// me.rootNode.redoColor(true)
		// rootNode.dumpTree(true)
	case widget.Window:
		// setCurrentWindow(w)
		n.doUserEvent()
	case widget.Tab:
		// setCurrentTab(w)
	case widget.Group:
		// n.placeWidgets()
		// n.toggleTree()
	case widget.Checkbox:
		if (n.B) {
			// n.setCheckbox(false)
		} else {
			// n.setCheckbox(true)
		}
		n.doUserEvent()
	case widget.Grid:
		// rootNode.hideWidgets()
		// n.placeGrid()
		// n.showWidgets()
	case widget.Box:
		// n.showWidgetPlacement(logNow, "drawTree()")
		if (n.B) {
			log(true, "BOX IS HORIZONTAL", n.Name)
		} else {
			log(true, "BOX IS VERTICAL", n.Name)
		}
	case widget.Button:
		n.doUserEvent()
	default:
	}
}
