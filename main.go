package main

import (
	"strings"
	"os"
	"fmt"
	"bufio"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(false)
	f, err := os.Open("/etc/passwd")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	cols := 7
	table.SetCell(0, 0, tview.NewTableCell("Username").SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 1, tview.NewTableCell("Password").SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 2, tview.NewTableCell("UID").SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 3, tview.NewTableCell("GID").SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 4, tview.NewTableCell("GECOS").SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 5, tview.NewTableCell("Homedir").SetTextColor(tcell.ColorYellow))
	table.SetCell(0, 6, tview.NewTableCell("Shell").SetTextColor(tcell.ColorYellow))
	r := 1
	for s.Scan() {
		split := strings.Split(s.Text(), `:`)
		for c := 0; c < cols; c++ {
			color := tcell.ColorWhite
			table.SetCell(r, c, tview.NewTableCell(split[c]).SetTextColor(color))
		}
		r += 1
	}
	if err := s.Err(); err != nil {
		fmt.Println(err)
	}
	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, false)
		}
	}).SetSelectedFunc(func(row int, column int) {
		for c := 0; c < cols; c++ {
			table.GetCell(row, c).SetTextColor(tcell.ColorRed)
		}
		table.SetSelectable(false, false)
	})
	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
