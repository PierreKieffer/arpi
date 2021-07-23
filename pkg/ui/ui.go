package ui

import (
	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
)

type BaseScreen struct {
	Screen  string
	Header  *widgets.Paragraph
	Status  *widgets.Paragraph
	UIList  *widgets.List
	Display *widgets.Paragraph

	Previous *BaseScreen
}

func (screen *BaseScreen) Create() {

	x, y := termui.TerminalDimensions()

	if screen.Header == nil {
		screen.Header = BuildHeader()
	}

	if screen.UIList == nil {
		screen.Screen = "home"
		items, details := Home()
		screen.UIList = items
		screen.Display = details
	}

	// header
	h := screen.Header
	h.SetRect(0, 0, x, 7)

	// menu list
	ls := screen.UIList
	ls.SelectedRowStyle = termui.NewStyle(termui.ColorMagenta)
	ls.TitleStyle.Fg = termui.ColorYellow
	ls.WrapText = false

	if screen.Display == nil && screen.Status == nil {
		ls.SetRect(0, 7, x, y)
		termui.Render(h, ls)

	} else if screen.Display != nil && screen.Status == nil {
		d := screen.Display
		d.TitleStyle.Fg = termui.ColorYellow

		ls.SetRect(0, 7, 16, y)
		d.SetRect(16, 7, x, y)
		termui.Render(h, ls, d)

	} else {
		d := screen.Display
		d.TitleStyle.Fg = termui.ColorYellow

		s := screen.Status

		ls.SetRect(0, 7, 16, y)
		s.SetRect(16, 7, x, 14)
		d.SetRect(16, 14, x, y)
		termui.Render(h, ls, s, d)

	}
}

func (screen *BaseScreen) HandleSelectItem() {

	selectedItem := screen.UIList.Rows[screen.UIList.SelectedRow]

	switch selectedItem {
	case " Scan ":
		/*
			Execute scan
		*/
		var previousScreen BaseScreen
		previousScreen = *screen

		items := HelpList()
		screen.Screen = "scan"
		screen.UIList = items
		screen.Display = Help()
		screen.Previous = &previousScreen

	case " Help ":
		/*
		   Go to Help page
		*/
		var previousScreen BaseScreen
		previousScreen = *screen

		items := HelpList()
		screen.Screen = "help"
		screen.UIList = items
		screen.Display = Help()
		screen.Previous = &previousScreen

	case " Home ":
		items, details := Home()
		screen.Screen = "home"
		screen.UIList = items
		screen.Display = details
		screen.Previous = nil
	}
}

var baseScreen BaseScreen

func App(network string) {

	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	baseScreen.Create()

	previousKey := ""

	uiEvents := termui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			baseScreen.UIList.ScrollDown()
		case "k", "<Up>":
			baseScreen.UIList.ScrollUp()
		case "<C-d>":
			baseScreen.UIList.ScrollHalfPageDown()
		case "<C-u>":
			baseScreen.UIList.ScrollHalfPageUp()
		case "<C-f>":
			baseScreen.UIList.ScrollPageDown()
		case "<C-b>":
			baseScreen.UIList.ScrollPageUp()
		case "<Enter>":
			baseScreen.HandleSelectItem()
		case "g":
			if previousKey == "g" {
				baseScreen.UIList.ScrollTop()
			}
		case "G", "<End>":
			baseScreen.UIList.ScrollBottom()
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		baseScreen.Create()
	}
}
