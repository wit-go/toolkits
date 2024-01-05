package main

import (
	"strconv"
	"github.com/andlabs/ui"
	"go.wit.com/gui/widget"
)

func (n *node) show(b bool) {
	if n.tk == nil {
		return
	}
	if n.tk.uiControl == nil {
		return
	}
	if (b) {
		n.tk.uiControl.Show()
	} else {
		n.tk.uiControl.Hide()
	}
}

func (n *node) enable(b bool) {
	if n == nil {
		panic("WHAT? enable was passed nil. How does this even happen?")
	}
	if n.tk == nil {
		return
	}
	if n.tk.uiControl == nil {
		return
	}
	if (b) {
		n.tk.uiControl.Enable()
	} else {
		n.tk.uiControl.Disable()
	}
}

func (n *node) pad(at widget.ActionType) {
	log(logInfo, "pad() on WidgetId =", n.WidgetId)

	t := n.tk
	if (t == nil) {
		log(logError, "pad() toolkit struct == nil. for", n.WidgetId)
		return
	}

	switch n.WidgetType {
	case widget.Group:
		switch at {
		case widget.Margin:
			t.uiGroup.SetMargined(true)
		case widget.Unmargin:
			t.uiGroup.SetMargined(false)
		case widget.Pad:
			t.uiGroup.SetMargined(true)
		case widget.Unpad:
			t.uiGroup.SetMargined(false)
		}
	case widget.Tab:
		switch at {
		case widget.Margin:
			tabSetMargined(t.uiTab, true)
		case widget.Unmargin:
			tabSetMargined(t.uiTab, false)
		case widget.Pad:
			tabSetMargined(t.uiTab, true)
		case widget.Unpad:
			tabSetMargined(t.uiTab, false)
		}
	case widget.Window:
		switch at {
		case widget.Margin:
			t.uiWindow.SetMargined(true)
		case widget.Unmargin:
			t.uiWindow.SetMargined(false)
		case widget.Pad:
			t.uiWindow.SetBorderless(false)
		case widget.Unpad:
			t.uiWindow.SetBorderless(true)
		}
	case widget.Grid:
		switch at {
		case widget.Margin:
			t.uiGrid.SetPadded(true)
		case widget.Unmargin:
			t.uiGrid.SetPadded(false)
		case widget.Pad:
			t.uiGrid.SetPadded(true)
		case widget.Unpad:
			t.uiGrid.SetPadded(false)
		}
	case widget.Box:
		switch at {
		case widget.Margin:
			t.uiBox.SetPadded(true)
		case widget.Unmargin:
			t.uiBox.SetPadded(false)
		case widget.Pad:
			t.uiBox.SetPadded(true)
		case widget.Unpad:
			t.uiBox.SetPadded(false)
		}
	case widget.Textbox:
		log(debugError, "TODO: implement ActionType =", at)
	default:
		log(debugError, "TODO: implement pad() for", at)
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

		stretchy = true
		if (p.tk.uiBox != nil) {
			p.tk.uiBox.Append(n.tk.uiControl, stretchy)
		}
		// log(debugNow, "is there a tParent parent? =", tParent.parent)
		// tParent.uiBox.Delete(0)

		// this didn't work:
		// tWidget.uiControl.Disable()
		// sleep(.8)
	default:
		log(logError, "TODO: need to implement move() for type =", n.WidgetType)
		log(logError, "TODO: need to implement move() for where =", p.ParentId)
		log(logError, "TODO: need to implement move() for widget =", n.WidgetId)
	}
}

func (n *node) Delete() {
	p := n.parent
	log(debugNow, "uiDelete()", n.WidgetId, "to", p.WidgetId)

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
		log(debugNow, "tWidget.boxC =", p.Name)
		log(debugNow, "is there a tParent parent? =", p.parent)
		if (p.tk.boxC < 1) {
			log(debugNow, "Can not delete from Box. already empty. tWidget.boxC =", p.tk.boxC)
			return
		}
		p.tk.uiBox.Delete(0)
		p.tk.boxC -= 1

		// this didn't work:
		// tWidget.uiControl.Disable()
		// sleep(.8)
		// tParent.uiBox.Append(tWidget.uiControl, stretchy)
	default:
		log(debugError, "TODO: need to implement uiDelete() for widget =", n.WidgetId, n.WidgetType)
		log(debugError, "TODO: need to implement uiDelete() for parent =", p.WidgetId, p.WidgetType)
	}
}

func rawAction(a *widget.Action) {
	log(logInfo, "rawAction() START a.ActionType =", a.ActionType)
	log(logInfo, "rawAction() START a.S =", a.S)

	if (a.ActionType == widget.InitToolkit) {
		// TODO: make sure to only do this once
		// go uiMain.Do(func() {
		// 	ui.Main(demoUI)
			// go catchActionChannel()
		// })
		// try doing this on toolkit load in init()
		return
	}

	log(logInfo, "rawAction() START a.WidgetId =", a.WidgetId, "a.ParentId =", a.ParentId)
	switch a.WidgetType {
	case widget.Flag:
		flag(a)
		return
	}

	n := me.rootNode.findWidgetId(a.WidgetId)

	if (a.ActionType == widget.Add) {
		ui.QueueMain(func() {
			add(a)
		})
		// TODO: remove this artificial delay
		// sleep(.001)
		return
	}

	if (a.ActionType == widget.Dump) {
		log(debugNow, "rawAction() Dump =", a.ActionType, a.WidgetType, n.Name)
		me.rootNode.listChildren(true)
		return
	}

	if (n == nil) {
		me.rootNode.listChildren(true)
		log(true, "rawAction() ERROR findWidgetId found nil", a.ActionType, a.WidgetType)
		log(true, "rawAction() ERROR findWidgetId found nil for id =", a.WidgetId)
		log(true, "rawAction() ERROR findWidgetId found nil", a.ActionType, a.WidgetType)
		log(true, "rawAction() ERROR findWidgetId found nil for id =", a.WidgetId)
		return
		panic("findWidgetId found nil for id = " + strconv.Itoa(a.WidgetId))
	}

	switch a.ActionType {
	case widget.Show:
		n.show(true)
	case widget.Hide:
		n.show(false)
	case widget.Enable:
		n.enable(true)
	case widget.Disable:
		n.enable(false)
	case widget.Get:
		n.setText(a)
	case widget.GetText:
		switch a.WidgetType {
		case widget.Textbox:
			a.S = n.S
		}
	case widget.Set:
		n.setText(a)
	case widget.SetText:
		n.setText(a)
	case widget.AddText:
		n.setText(a)
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
		log(debugNow, "rawAction() attempt to move() =", a.ActionType, a.WidgetType)
		newParent := me.rootNode.findWidgetId(a.ParentId)
		n.move(newParent)
	default:
		log(debugError, "rawAction() Unknown =", a.ActionType, a.WidgetType)
	}
	log(logInfo, "rawAction() END =", a.ActionType, a.WidgetType)
}
