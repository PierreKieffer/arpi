package main

import (
	"flag"
	// "fmt"
	// "github.com/PierreKieffer/arpi/pkg/netscan"
	"github.com/PierreKieffer/arpi/pkg/ui"
)

// var exit = make(chan bool)

var banner = `

	╔═╗╦═╗╔═╗╦
	╠═╣╠╦╝╠═╝║
	╩ ╩╩╚═╩  ╩
`

func main() {

	net := flag.String("net", "192.168.1.0/24", "Network")
	ui.App(*net)

	// fmt.Println(banner)

	// scanner := netscan.InitScanner(*net)

	// go scanner.LogHandler()
	// go scanner.SigHandler()
	// scanner.SigChan <- "scan"

	// <-exit

}
