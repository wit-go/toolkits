package main

import (
	"go.wit.com/gui/widget"
)

func (n *node) setText(a *widget.Action) {
	log(debugChange, "setText() START with a.S =", a.S)
	t := n.tk
	if (t == nil) {
		log(debugError, "setText error. tk == nil", n.Name, n.WidgetId)
		actionDump(debugError, a)
		return
	}
	log(debugChange, "setText() Attempt on", n.WidgetType, "with", a.S)

	switch n.WidgetType {
	case widget.Window:
		t.uiWindow.SetTitle(a.S)
	case widget.Tab:
	case widget.Group:
		t.uiGroup.SetTitle(a.S)
	case widget.Checkbox:
		switch a.ActionType {
		case widget.SetText:
			t.uiCheckbox.SetText(a.S)
		case widget.Get:
			n.B = t.uiCheckbox.Checked()
		case widget.Set:
			// TODO: commented out while working on chan
			t.uiCheckbox.SetChecked(a.B)
		default:
			log(debugError, "setText() unknown", a.ActionType, "on checkbox", n.Name)
		}
	case widget.Textbox:
		switch a.ActionType {
		case widget.Set:
			if (t.uiEntry != nil) {
				t.uiEntry.SetText(a.S)
			}
			if (t.uiMultilineEntry != nil) {
				t.uiMultilineEntry.SetText(a.S)
			}
		case widget.SetText:
			if (t.uiEntry != nil) {
				t.uiEntry.SetText(a.S)
			}
			if (t.uiMultilineEntry != nil) {
				t.uiMultilineEntry.SetText(a.S)
			}
		default:
			log(debugError, "setText() unknown", a.ActionType, "on checkbox", n.Name)
		}
	case widget.Label:
		t.uiLabel.SetText(a.S)
	case widget.Button:
		t.uiButton.SetText(a.S)
	case widget.Slider:
		switch a.ActionType {
		case widget.Get:
			n.I = t.uiSlider.Value()
		case widget.Set:
			t.uiSlider.SetValue(a.I)
		default:
			log(debugError, "setText() unknown", a.ActionType, "on checkbox", n.Name)
		}
	case widget.Spinner:
		switch a.ActionType {
		case widget.Get:
			n.I = t.uiSpinbox.Value()
		case widget.Set:
			t.uiSpinbox.SetValue(a.I)
		default:
			log(debugError, "setText() unknown", a.ActionType, "on checkbox", n.Name)
		}
	case widget.Dropdown:
		switch a.ActionType {
		case widget.AddText:
			n.AddDropdownName(a.S)
		case widget.Set:
			var orig int
			var i int = -1
			var s string
			orig = t.uiCombobox.Selected()
			log(debugChange, "try to set the Dropdown to", a.S, "from", orig)
			// try to find the string
			for i, s = range t.val {
				log(debugChange, "i, s", i, s)
				if (a.S == s) {
					t.uiCombobox.SetSelected(i)
					log(debugChange, "setText() Dropdown worked.", n.S)
					return
				}
			}
			log(debugError, "setText() Dropdown did not find:", a.S)
			// if i == -1, then there are not any things in the menu to select
			if (i == -1) {
				return
			}
			// if the string was never set, then set the dropdown to the last thing added to the menu
			if (orig == -1) {
				t.uiCombobox.SetSelected(i)
			}
		case widget.Get:
			// t.S = t.s
		case widget.GetText:
			// t.S = t.s
		default:
			log(debugError, "setText() unknown", a.ActionType, "on checkbox", n.Name)
		}
	case widget.Combobox:
		switch a.ActionType {
		case widget.AddText:
			t.AddComboboxName(a.S)
		case widget.Set:
			t.uiEditableCombobox.SetText(a.S)
			n.S = a.S
		case widget.SetText:
			t.uiEditableCombobox.SetText(a.S)
			n.S = a.S
		default:
			log(debugError, "setText() unknown", a.ActionType, "on checkbox", n.Name)
		}
	default:
		log(debugError, "plugin Send() Don't know how to setText on", n.WidgetType, "yet", a.ActionType)
	}
	log(debugChange, "setText() END with a.S =", a.S)
}
