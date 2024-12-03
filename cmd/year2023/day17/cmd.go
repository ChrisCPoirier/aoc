package day17

import (
	"aoc/cmd/grid"
	"container/heap"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day17",
	Short: "day17",
	Long:  `day17`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf(`cmd/year%s/%s/1.txt`, parent, command))

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("score part1: %d", part1(string(b), 0, 3))
	logrus.Infof("score part2: %d", part1(string(b), 4, 10))
	visualize()
}

func part1(s string, min, max int) int64 {
	// var score int = 0

	g := grid.New(s, "").Ints()

	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq,
		&cityBlock{
			heatLoss:        0,
			row:             0,
			column:          0,
			rowDirection:    0,
			columnDirection: 0,
			count:           0,
		})

	visited := map[string]bool{}

	for pq.Len() > 0 {
		cb := heap.Pop(pq).(*cityBlock)

		if cb.row == len(g)-1 && cb.column == len(g[0])-1 && cb.count >= min {
			return int64(cb.heatLoss)
		}

		key := fmt.Sprintf("%d,%d,%d,%d,%d", cb.row, cb.column, cb.rowDirection, cb.columnDirection, cb.count)

		if _, ok := visited[key]; ok {
			continue
		}

		visited[key] = true

		for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			newRowDirection := d[0]
			newColumnDirection := d[1]
			newRow := cb.row + newRowDirection
			newColumn := cb.column + newColumnDirection

			if 0 > newRow || newRow > len(g)-1 || 0 > newColumn || newColumn > len(g[0])-1 {
				continue
			}

			if cb.rowDirection == -newRowDirection && cb.columnDirection == -newColumnDirection {
				continue
			}

			c := 1

			if cb.rowDirection == newRowDirection && cb.columnDirection == newColumnDirection {
				c += cb.count
			} else {
				if cb.count < min && !(cb.row == 0 && cb.column == 0) {
					continue
				}
			}

			if c > max {
				continue
			}

			heap.Push(pq,
				&cityBlock{
					heatLoss:        cb.heatLoss + g[newRow][newColumn],
					row:             newRow,
					column:          newColumn,
					rowDirection:    newRowDirection,
					columnDirection: newColumnDirection,
					count:           c,
				})
		}

	}

	return int64(-99999)
}

func visualize() {

	myApp := app.New()
	w := myApp.NewWindow("Board Layout")
	// w.Resize(fyne.NewSize(1920, 1080))
	// content := container.NewMax()

	// content.Add(Show(w))

	content := container.New(layout.NewHBoxLayout(), Show(w))

	// items := container.NewCenter(content)
	// w.SetContent(items)
	w.SetContent(container.New(layout.NewVBoxLayout(), content))
	w.Resize(fyne.NewSize(1920, 1080))

	w.ShowAndRun()
	// myApp := app.New()
	// myWindow := myApp.NewWindow("Box Layout")

	// text1 := canvas.NewText("Hello", color.White)
	// text2 := canvas.NewText("There", color.White)
	// text3 := canvas.NewText("(right)", color.White)
	// content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	// text4 := canvas.NewText("centered", color.White)
	// centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	// myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	// myWindow.ShowAndRun()

	// type appInfo struct {
	// 	name string
	// 	icon fyne.Resource
	// 	canv bool
	// 	run  func(fyne.Window) fyne.CanvasObject
	// }

	// var apps = []appInfo{
	// 	// {"Bugs", icon.BugBitmap, false, bugs.Show},
	// 	// {"XKCD", icon.XKCDBitmap, false, xkcd.Show},
	// 	// {"Clock", icon.ClockBitmap, true, clock.Show},
	// 	// {"Fractal", icon.FractalBitmap, true, fractal.Show},
	// 	// {"Tic Tac Toe", nil, true, tictactoe.Show},
	// }

	// a := app.New()
	// // a.SetIcon(resourceIconPng)

	// content := container.NewMax()
	// w := a.NewWindow("Examples")

	// // apps[4].icon = theme.RadioButtonIcon() // lazy load Fyne resource to avoid error
	// appList := widget.NewList(
	// 	func() int {
	// 		return len(apps)
	// 	},
	// 	func() fyne.CanvasObject {
	// 		icon := &canvas.Image{}
	// 		label := widget.NewLabel("Text Editor")
	// 		labelHeight := label.MinSize().Height
	// 		icon.SetMinSize(fyne.NewSize(labelHeight, labelHeight))
	// 		return container.NewBorder(nil, nil, icon, nil,
	// 			label)
	// 	},
	// 	func(id widget.ListItemID, obj fyne.CanvasObject) {
	// 		img := obj.(*fyne.Container).Objects[1].(*canvas.Image)
	// 		text := obj.(*fyne.Container).Objects[0].(*widget.Label)
	// 		img.Resource = apps[id].icon
	// 		img.Refresh()
	// 		text.SetText(apps[id].name)
	// 	})
	// appList.OnSelected = func(id widget.ListItemID) {
	// 	content.Objects = []fyne.CanvasObject{apps[id].run(w)}
	// }
	// content.Objects = []fyne.CanvasObject{Show(w)}
	// split := container.NewHSplit(appList, content)
	// split.Offset = 0.1
	// w.SetContent(split)
	// w.Resize(fyne.NewSize(480, 360))
	// w.ShowAndRun()
}
