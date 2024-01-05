package main

import (
	"go.wit.com/gui/widget"
)

func (n *node) show(b bool) {
}

func (n *node) enable(b bool) {
}

func (n *node) pad(at widget.ActionType) {
	switch n.WidgetType {
	case widget.Group:
		switch at {
		case widget.Margin:
			// SetMargined(true)
		case widget.Unmargin:
			// SetMargined(false)
		case widget.Pad:
			// SetMargined(true)
		case widget.Unpad:
			// SetMargined(false)
		}
	case widget.Tab:
	case widget.Window:
	case widget.Grid:
	case widget.Box:
	case widget.Textbox:
		log(logError, "TODO: implement ActionType =", at)
	default:
		log(logError, "TODO: implement pad() for", at)
	}
}

func (n *node) move(newParent *node) {
	p := n.parent

	switch p.WidgetType {
	case widget.Group:
	case widget.Tab:
		// tabSetMargined(tParent.uiTab, true)
	case widget.Window:
		// t.uiWindow.SetBorderless(false)
	case widget.Grid:
		// t.uiGrid.SetPadded(true)
	case widget.Box:
		log(logInfo, "TODO: move() where =", p.ParentId)
		log(logInfo, "TODO: move() for widget =", n.WidgetId)
	default:
		log(logError, "TODO: need to implement move() for type =", n.WidgetType)
		log(logError, "TODO: need to implement move() for where =", p.ParentId)
		log(logError, "TODO: need to implement move() for widget =", n.WidgetId)
	}
}

func (n *node) Delete() {
	p := n.parent
	log(logNow, "uiDelete()", n.WidgetId, "to", p.WidgetId)

	switch p.WidgetType {
	case widget.Group:
		// tParent.uiGroup.SetMargined(true)
	case widget.Tab:
		// tabSetMargined(tParent.uiTab, true)
	case widget.Window:
		// t.uiWindow.SetBorderless(false)
	case widget.Grid:
		// t.uiGrid.SetPadded(true)
	case widget.Box:
		log(logNow, "tWidget.boxC =", p.Name)
		log(logNow, "is there a tParent parent? =", p.parent)
		// this didn't work:
		// tWidget.uiControl.Disable()
		// sleep(.8)
		// tParent.uiBox.Append(tWidget.uiControl, stretchy)
	default:
		log(logError, "TODO: need to implement uiDelete() for widget =", n.WidgetId, n.WidgetType)
		log(logError, "TODO: need to implement uiDelete() for parent =", p.WidgetId, p.WidgetType)
	}
}

func doAction(a *widget.Action) {
	log(logNow, "doAction() START a.ActionType =", a.ActionType)
	log(logNow, "doAction() START a.S =", a.S)

	if (a.ActionType == widget.InitToolkit) {
		// TODO: make sure to only do this once
		// go uiMain.Do(func() {
		// 	ui.Main(demoUI)
			// go catchActionChannel()
		// })
		// try doing this on toolkit load in init()
		return
	}

	log(logNow, "doAction() START a.WidgetId =", a.WidgetId, "a.ParentId =", a.ParentId)
	switch a.WidgetType {
	case widget.Root:
		me.rootNode = addNode(a)
		log(logNow, "doAction() found rootNode")
		return
	case widget.Flag:
		// flag(&a)
		return
	}

	n := me.rootNode.findWidgetId(a.WidgetId)

	switch a.ActionType {
	case widget.Add:
		addNode(a)
	case widget.Show:
		n.show(true)
	case widget.Hide:
		n.show(false)
	case widget.Enable:
		n.enable(true)
	case widget.Disable:
		n.enable(false)
	case widget.Get:
		// n.setText(a.S)
	case widget.GetText:
		switch a.WidgetType {
		case widget.Textbox:
			a.S = n.S
		}
	case widget.Set:
		// n.setText(a.S)
	case widget.SetText:
		// n.setText(a.S)
	case widget.AddText:
		// n.setText(a.S)
	case widget.Margin:
		n.pad(widget.Unmargin)
	case widget.Unmargin:
		n.pad(widget.Margin)
	case widget.Pad:
		n.pad(widget.Pad)
	case widget.Unpad:
		n.pad(widget.Unpad)
	case widget.Delete:
		n.Delete()
	case widget.Move:
		log(logNow, "doAction() attempt to move() =", a.ActionType, a.WidgetType)
		newParent := me.rootNode.findWidgetId(a.ParentId)
		n.move(newParent)
	default:
		log(logError, "doAction() Unknown =", a.ActionType, a.WidgetType)
	}
	log(logInfo, "doAction() END =", a.ActionType, a.WidgetType)
}
