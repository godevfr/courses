package gui

import (
	"image"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/image/draw"
)

func StartGUI(title string, photos []image.Image) {
	w := app.NewWindow()
	if err := loop(w, title, photos); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

type C = layout.Context
type D = layout.Dimensions

func loop(w *app.Window, title string, photos []image.Image) error {
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
					return list.Layout(gtx, len(photos), func(gtx C, i int) D {
						return in.Layout(gtx, func(gtx C) D {
							return layoutPhoto(gtx, photos[i])
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

func layoutPhoto(gtx layout.Context, photo image.Image) layout.Dimensions {
	sz := gtx.Constraints.Min.X

	var ratio float64
	if photo.Bounds().Size().Y != 0 {
		ratio = float64(photo.Bounds().Size().X) / float64(photo.Bounds().Size().Y)
	}
	if ratio == 0.0 {
		ratio = 1.0
	}

	img := image.NewRGBA(image.Rectangle{Max: image.Point{X: sz, Y: int(float64(sz) / ratio)}})
	draw.ApproxBiLinear.Scale(img, img.Bounds(), photo, photo.Bounds(), draw.Src, nil)
	photoOp := paint.NewImageOp(img)

	imgWidget := widget.Image{Src: photoOp}
	imgWidget.Scale = float32(sz) / float32(gtx.Px(unit.Dp(float32(sz))))
	return imgWidget.Layout(gtx)
}
