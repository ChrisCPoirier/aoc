package matrix

import (
	"aoc/cmd/common"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (s Strings) Fyne(window fyne.Window) *fyne.Container {

	squares := []fyne.CanvasObject{}

	for _, r := range s {
		for _, v := range r {
			squares = append(squares, NewSquare(v, color.RGBA{0, 0, 0, 100}))
		}
	}

	c := container.New(common.NewGridLayout(len(s)), squares...)

	window.SetContent(c)
	return c
}

func NewSquare(s string, c color.RGBA) fyne.CanvasObject {
	sqr := canvas.NewRectangle(c)
	sqr.Resize(
		fyne.Size{
			Width:  1,
			Height: 1,
		},
	)
	sqr.StrokeColor = c
	sqr.StrokeWidth = 0
	text := canvas.NewText(s, color.White)
	text.Resize(
		fyne.Size{
			Width:  1,
			Height: 1,
		},
	)
	text.TextSize = 5
	stack := container.NewStack(sqr, text)
	stack.Resize(
		fyne.Size{
			Width:  1,
			Height: 1,
		},
	)

	return stack
}
