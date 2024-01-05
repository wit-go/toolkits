package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"go.wit.com/gui/widget"
)

func actionDump(b bool, a *widget.Action) {
	log(b, "actionDump() Widget.Type =", a.ActionType)
	log(b, "actionDump() Widget.S =", a.S)
	log(b, "actionDump() Widget.I =", a.I)
	log(b, "actionDump() WidgetId =", a.WidgetId)
	log(b, "actionDump() ParentId =", a.ParentId)
}

func add(a *widget.Action) {
	if (a.WidgetType == widget.Root) {
		me.rootNode = addNode(a)
		return
	}
	n := addNode(a)

	p := n.parent
	switch n.WidgetType {
	case widget.Window:
		newWindow(n)
		return
	case widget.Tab:
		p.newTab(n)
		return
	case widget.Label:
		p.newLabel(n)
		return
	case widget.Button:
		p.newButton(n)
		return
	case widget.Grid:
		p.newGrid(n)
		return
	case widget.Checkbox:
		p.newCheckbox(n)
		return
	case widget.Spinner:
		p.newSpinner(n)
		return
	case widget.Slider:
		p.newSlider(n)
		return
	case widget.Dropdown:
		p.newDropdown(n)
		return
	case widget.Combobox:
		p.newCombobox(n)
		return
	case widget.Textbox:
		p.newTextbox(n)
		return
	case widget.Group:
		p.newGroup(n)
		return
	case widget.Box:
		p.newBox(n)
		return
	case widget.Image:
		p.newImage(n)
		return
	default:
		log(debugError, "add() error TODO: ", n.WidgetType, n.Name)
	}
}

// This routine is very specific to this toolkit
// It's annoying and has to be copied to each widget when there are changes
// it could be 'simplfied' maybe or made to be more generic, but this is as far as I've gotten
// it's probably not worth working much more on this toolkit, the andlabs/ui has been great and got me here!
// but it's time to write direct GTK, QT, macos and windows toolkit plugins
// -- jcarr 2023/03/09

// Grid numbering examples by (X,Y)
// ---------
// -- (1) --
// -- (2) --
// ---------
//
// -----------------------------
// -- (1,1) -- (1,2) -- (1,3) --
// -- (2,1) -- (2,2) -- (2,3) --
// -----------------------------

// internally for andlabs/ui
// (x&y flipped and start at zero) 
// -----------------------------
// -- (0,0) -- (1,0) -- (1,0) --
// -- (0,1) -- (1,1) -- (1,1) --
// -----------------------------
func (p *node) place(n *node) bool {
	log(logInfo, "place() START", n.WidgetType, n.Name)

	if (p.tk == nil) {
		log(logError, "p.tk == nil", p.Name, p.ParentId, p.WidgetType, p.tk)
		log(logError, "n = ", n.Name, n.ParentId, n.WidgetType, n.tk)
		panic("p.tk == nil")
	}

	log(logInfo, "place() switch", p.WidgetType)
	switch p.WidgetType {
	case widget.Grid:
		log(logInfo, "place() Grid try at Parent X,Y =", n.X, n.Y)
		n.tk.gridX = n.AtW - 1
		n.tk.gridY = n.AtH - 1
		log(logInfo, "place() Grid try at gridX,gridY", n.tk.gridX, n.tk.gridY)
		// at the very end, subtract 1 from X & Y since andlabs/ui starts counting at zero
		p.tk.uiGrid.Append(n.tk.uiControl,
			n.tk.gridX, n.tk.gridY, 1, 1,
			false, ui.AlignFill, false, ui.AlignFill)
		return true
	case widget.Group:
		if (p.tk.uiBox == nil) {
			p.tk.uiGroup.SetChild(n.tk.uiControl)
			log(logInfo, "place() hack Group to use this as the box?", n.Name, n.WidgetType)
			p.tk.uiBox  = n.tk.uiBox
		} else {
			p.tk.uiBox.Append(n.tk.uiControl, stretchy)
		}
		return true
	case widget.Tab:
		if (p.tk.uiTab == nil) {
			log(logError, "p.tk.uiTab == nil for n.WidgetId =", n.WidgetId, "p.tk =", p.tk)
			panic("p.tk.uiTab == nil")
		}
		if (n.tk.uiControl == nil) {
			log(logError, "n.tk.uiControl == nil for n.WidgetId =", n.WidgetId, "n.tk =", n.tk)
			panic("n.tk.uiControl == nil")
		}
		log(logError, "CHECK LOGIC ON THIS. APPENDING directly into a window without a tab")
		// log(logError, "THIS SHOULD NEVER HAPPEN ??????? trying to place() node=", n.WidgetId, n.Name, n.Text, n.WidgetType)
		// log(logError, "THIS SHOULD NEVER HAPPEN ??????? trying to place() on parent=", p.WidgetId, p.Name, p.Text, p.WidgetType)
		// panic("n.tk.uiControl == nil")
		p.tk.uiTab.Append(n.Text, n.tk.uiControl)
		p.tk.boxC += 1
		return true
	case widget.Box:
		log(logInfo, "place() uiBox =", p.tk.uiBox)
		log(logInfo, "place() uiControl =", n.tk.uiControl)
		p.tk.uiBox.Append(n.tk.uiControl, stretchy)
		p.tk.boxC += 1
		return true
	case widget.Window:
		p.tk.uiWindow.SetChild(n.tk.uiControl)
		return true
	default:
		log(debugError, "place() how? Parent =", p.WidgetId, p.WidgetType)
	}
	return false
}
