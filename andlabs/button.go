package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func (p *node) newButton(n *node) {
	log(debugToolkit, "newButton() START", n.Name)

	t := p.tk
	if (t == nil) {
		log(debugToolkit, "newButton() toolkit struct == nil. name=", n.Name)
		return
	}

	newt := new(guiWidget)

	b := ui.NewButton(n.Text)
	newt.uiButton = b
	newt.uiControl = b
	newt.parent = t

	b.OnClicked(func(*ui.Button) {
		n.doUserEvent()
	})

	n.tk = newt
	p.place(n)
	log(debugToolkit, "newButton() END", n.Name)
}
