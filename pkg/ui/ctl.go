package ui

import (
	// termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func BuildHeader() *widgets.Paragraph {
	header := widgets.NewParagraph()
	header.Text = `
  ╔═╗╦═╗╔═╗╦
  ╠═╣╠╦╝╠═╝║
  ╩ ╩╩╚═╩  ╩
`

	return header
}

func Home() (*widgets.List, *widgets.Paragraph) {
	options := widgets.NewList()
	options.Title = "Home"
	options.Rows = []string{" Scan ", " Help "}

	details := widgets.NewParagraph()
	details.Text = `

     -----------------------------
     -        Move around        -
     -----------------------------
     go up               ▲  or 'k'
     go down             ▼  or 'j'
     go to the top       'gg'
     go to the bottom    'G'
     select item         'enter'
     Quit                'q'

`
	return options, details
}

func Help() *widgets.Paragraph {
	help := widgets.NewParagraph()
	help.Text = `
       	     ╔═╗╦═╗╔═╗╦
       	     ╠═╣╠╦╝╠═╝║
       	     ╩ ╩╩╚═╩  ╩

     -----------------------------
     -        Move around        -
     -----------------------------
     go up               ▲  or 'k'
     go down             ▼  or 'j'
     go to the top       'gg'
     go to the bottom    'G'
     select item         'enter'
     Quit                'q'




     -----------------------------
     -          Author           -
     -----------------------------
     https://github.com/PierreKieffer
`

	return help
}

func HelpList() *widgets.List {

	helpList := widgets.NewList()
	helpList.Title = "Help"

	utils := []string{" Home "}
	helpList.Rows = append(helpList.Rows, utils...)

	return helpList
}

func ExecScan() {
}

func ScanLogHandler(scanner *netscan.Scanner, signal chan bool) {

	for {
		select {
		case signal:
			return

		case log := <-scanner.LogChan:

		}
	}
}
