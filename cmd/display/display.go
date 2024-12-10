package display

import (
	"aoc/cmd/matrix"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var BLACK = color.RGBA{0, 0, 0, 255}
var GREEN = color.RGBA{0, 255, 0, 255}
var RED = color.RGBA{255, 0, 0, 255}
var BLUE = color.RGBA{0, 0, 255, 255}

type display struct {
	Matrix    matrix.Strings
	app       fyne.App
	window    fyne.Window
	container *fyne.Container
}

func New(m matrix.Strings) display {
	myApp := app.New()
	myWindow := myApp.NewWindow("visualize")
	c := m.Fyne(myWindow)

	return display{app: myApp, window: myWindow, container: c, Matrix: m}
}

func (d display) ShowAndRun() {
	d.window.ShowAndRun()
}

func (d display) ColorCell(r, c int, cl color.RGBA) {
	d.container.Objects[c+(len(d.Matrix[r])*r)] = matrix.NewSquare(d.Matrix[r][c], cl)
}

func (d display) ColorCells(cells [][]int, cl color.RGBA) {
	for _, cell := range cells {
		time.Sleep(time.Millisecond * 5)
		d.ColorCell(cell[0], cell[1], cl)
	}
}

func (d display) Reset() {
	for r, items := range d.Matrix {
		for c := range items {
			d.ColorCell(r, c, BLACK)
		}
	}
}

func (d display) ColorCellsNoWait(cells [][]int, cl color.RGBA) {
	for _, cell := range cells {
		d.ColorCell(cell[0], cell[1], cl)
	}
	d.container.Refresh()
}
