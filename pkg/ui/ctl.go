package ui

import (
	// termui "github.com/gizak/termui/v3"
	"fmt"
	"github.com/PierreKieffer/arpi/pkg/netscan"
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
	options.Title = " Home "
	options.Rows = []string{" Scan ", " About "}

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
	helpList.Title = "About"

	utils := []string{" Home "}
	helpList.Rows = append(helpList.Rows, utils...)

	return helpList
}

func BuildScanReport(scanner *netscan.Scanner) []string {

	var report []string

	report = []string{" Home ",
		" Scan again ",
		"",
		"",
		"        IP       |        MAC        |          NAME          ",
		"---------------------------------------------------------------",
	}

	for _, device := range scanner.Devices {
		reportLine := fmt.Sprintf("%v%v%v", FmtReport("ip", device.IP), FmtReport("mac", device.MAC), FmtReport("name", device.Name))
		report = append(report, reportLine)
	}

	report = []string{" Home ",
		" Scan again ",
		"",
		"",
		"        IP       |        MAC        |          NAME          ",
		"---------------------------------------------------------------",
		" 192.168.1.55    | DC:B7:23:82:7C:2E | MBP-de-Pierre",
		" 192.168.1.56    | DC:C6:43:B0:16:1C | raspberrypi",
		" 192.168.1.92    | A8:D4:89:36:E9:83 | HUAWEI_P20_lite",
		" 192.168.1.254   | 55:D3:35:F8:GH:6A | bbox.lan ",
		" 192.168.1.16    |                   | pierre-xps-13-9380",
	}

	return report
}

func FmtReport(valueType string, value string) string {
	switch valueType {
	case "ip":
		value = fmt.Sprintf(" %v", value)
		for len(value) < 17 {
			value = fmt.Sprintf("%v ", value)
		}
		return fmt.Sprintf("%v|", value)
	case "mac":
		value = fmt.Sprintf(" %v", value)
		for len(value) < 19 {
			value = fmt.Sprintf("%v ", value)
		}
		return fmt.Sprintf("%v|", value)
	case "name":
		value = fmt.Sprintf(" %v", value)
		return value
	}

	return ""
}
