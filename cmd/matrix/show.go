package matrix

import (
	"aoc/cmd/common"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (s Strings) Show(window fyne.Window) {
	window.SetPadded(false)

	squares := []fyne.CanvasObject{}

	for _, r := range s {
		for _, v := range r {
			squares = append(squares, newSquare(v))
		}
	}

	window.SetPadded(false)
	c := container.New(common.NewGridLayout(len(s)), squares...)

	window.SetContent(c)
}

func newSquare(s string) fyne.CanvasObject {
	// color.RGBA{234, 239, 44, 100} // yellow
	sqr := canvas.NewRectangle(color.Black)
	sqr.Resize(
		fyne.Size{
			Width:  1,
			Height: 1,
		},
	)
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
