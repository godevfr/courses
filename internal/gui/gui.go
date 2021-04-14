package gui

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func StartGUI(title string, data []string) {
	w := app.NewWindow()
	if err := loop(w, title, data); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

type C = layout.Context
type D = layout.Dimensions

func loop(w *app.Window, title string, data []string) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops

	list := &layout.List{
		Axis: layout.Vertical,
	}

	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			in := layout.UniformInset(unit.Dp(8))
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return in.Layout(gtx, func(gtx C) D {
						t := material.H1(th, title)

						return t.Layout(gtx)
					})
				}),

				layout.Rigid(func(gtx C) D {
					return list.Layout(gtx, len(data), func(gtx C, i int) D {
						return in.Layout(gtx, func(gtx C) D {
							d := material.Body1(th, data[i])
							return d.Layout(gtx)
						})
					})
				}),
			)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}

	return nil
}
