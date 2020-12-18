package main

import (
	"github.com/EvilKhaosKat/experiments/games/ttt/core"
	"github.com/rivo/tview"
	"time"
)

func main() {
	game := core.NewGame(5)
	go core.LaunchGameCycle(game)

	time.Sleep(4000 * time.Millisecond)

	/*app := tview.NewApplication()
	flex := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)

	sideBar := tview.NewList().
		AddItem("List item 1", "Some explanatory text", 'a', nil).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("List item 3", "Some explanatory text", 'c', nil).
		AddItem("List item 4", "Some explanatory text", 'd', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	flex.AddItem(sideBar, 0, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(sideBar).Run(); err != nil {
		panic(err)
	}*/
}

func refresh(app *tview.Application) {
	tick := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-tick.C:
			app.Draw()
		}
	}
}
