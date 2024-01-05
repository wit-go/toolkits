package main

import (
	"go.wit.com/gui/widget"
)

// this is specific to the nocui toolkit
func initWidget(n *node) *guiWidget {
	var w *guiWidget
	w = new(guiWidget)
	// Set(w, "default")

	if n.WidgetType == widget.Root {
		log(logInfo, "setupWidget() FOUND ROOT w.id =", n.WidgetId)
		n.WidgetId = 0
		me.rootNode = n
		return w
	}

	if (n.WidgetType == widget.Box) {
		if (n.B) {
			n.horizontal = true
		} else {
			n.horizontal = false
		}
	}

	return w
}
