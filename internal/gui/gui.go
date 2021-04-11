package gui

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func StartGUI(title, data string) {
	w := app.NewWindow()
	if err := loop(w, title, data); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func loop(w *app.Window, title, data string) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops

	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			t := material.H1(th, title)
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}

			t.Alignment = text.Middle
			t.Color = maroon
			t.Layout(gtx)

			d := material.Body1(th, data)
			d.Layout(gtx)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}

	return nil
}
