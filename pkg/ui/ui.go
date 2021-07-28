package ui

import (
	"fmt"
	"github.com/PierreKieffer/arpi/pkg/netscan"
	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"time"
)

type BaseScreen struct {
	Screen  string
	Header  *widgets.Paragraph
	Status  *widgets.Paragraph
	UIList  *widgets.List
	Display *widgets.Paragraph

	Previous *BaseScreen
}

var (
	signal     = make(chan bool)
	scanner    = &netscan.Scanner{}
	baseScreen BaseScreen
	uiEvents   = termui.PollEvents()
)

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

	} else if screen.Display == nil && screen.Status != nil {

		s := screen.Status

		ls.SetRect(0, 10, x, y)
		s.SetRect(0, 7, x, 10)
		termui.Render(h, ls, s)

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

func (screen *BaseScreen) Update() {

	switch screen.Screen {
	case "scan":

		go scanner.Scan()

		go func() {
			for {
				select {

				case <-signal:
					return

				case log := <-scanner.LogChan:

					if log == "Network scan in progress ... " {
						screen.Status.Text = log
						screen.Status.TextStyle.Fg = termui.ColorRed
						time.Sleep(1 * time.Second)
						termui.Render(screen.Status)

					}

					if log == "Network scan completed" {

						screen.Status.Text = fmt.Sprintf("%v : %v", log, scanner.Summary)
						screen.Status.TextStyle.Fg = termui.ColorGreen
						screen.UIList.Rows = BuildScanReport(scanner)
						screen.UIList.SelectedRow = 0
						time.Sleep(1 * time.Second)
						termui.Render(screen.Status, screen.UIList)
						signal <- true
					}

				}
			}
		}()

	case "home":
		screen.Create()
	}
}

func (screen *BaseScreen) HandleSelectItem() {

	selectedItem := screen.UIList.Rows[screen.UIList.SelectedRow]

	switch selectedItem {
	case " Scan ", " Scan again ":
		/*
			Execute scan
		*/
		var previousScreen BaseScreen
		previousScreen = *screen

		screen.Screen = "scan"

		screen.UIList.Rows = []string{}
		screen.UIList.Title = " Network scan report "
		screen.Status = widgets.NewParagraph()
		screen.Status.Text = ""
		screen.Display = nil

		screen.Previous = &previousScreen

	case " About ":
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
		screen.Status = nil
		screen.Previous = nil

	default:
		return
	}

	screen.Update()
}

func App(network string) {

	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	scanner.Network = network
	scanner.InitChan()

	baseScreen.Create()

	previousKey := ""

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			if len(baseScreen.UIList.Rows) > 0 {
				baseScreen.UIList.ScrollDown()
				selectedItem := baseScreen.UIList.Rows[baseScreen.UIList.SelectedRow]
				if selectedItem != " Home " && selectedItem != " Scan " && selectedItem != " Scan again " && selectedItem != " About " {

					baseScreen.UIList.ScrollUp()
				}
			}
		case "k", "<Up>":
			if len(baseScreen.UIList.Rows) > 0 {
				baseScreen.UIList.ScrollUp()
			}
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
