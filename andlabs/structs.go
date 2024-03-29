package main

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// var andlabs map[int]*andlabsT
// var callback func(int) bool
// var callback chan toolkit.Action

// It's probably a terrible idea to call this 'me'
var me config

type config struct {
	rootNode *node // the base of the binary tree. it should have id == 0
}

// stores the raw toolkit internals
type guiWidget struct {
	Width  int
	Height int

	// tw	*toolkit.Widget
	parent	*guiWidget
	children []*guiWidget

	// used to track if a tab has a child widget yet
	child bool

	uiControl ui.Control

	uiBox     *ui.Box
	uiButton  *ui.Button
	uiCombobox *ui.Combobox
	uiCheckbox *ui.Checkbox
	uiEntry   *ui.Entry
	uiGroup   *ui.Group
	uiLabel   *ui.Label
	uiSlider  *ui.Slider
	uiSpinbox *ui.Spinbox
	uiTab     *ui.Tab
	uiWindow  *ui.Window
	uiMultilineEntry   *ui.MultilineEntry
	uiEditableCombobox    *ui.EditableCombobox
	uiImage  *ui.Image

	uiGrid    *ui.Grid
	gridX	int
	gridY	int

	// used as a counter to work around limitations of widgets like combobox
	// this is probably fucked up and in many ways wrong because of unsafe goroutine threading
	// but it's working for now due to the need for need for a correct interaction layer betten toolkits
	c int
	val map[int]string

	// andlabs/ui only accesses widget id numbers
	boxC int	// how many things on in a box or how many tabs
}
