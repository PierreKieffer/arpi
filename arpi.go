package main

import (
	"flag"
	// "fmt"
	// "github.com/PierreKieffer/arpi/pkg/netscan"
	"github.com/PierreKieffer/arpi/pkg/ui"
	"log"
	"os"
)

// var exit = make(chan bool)

var banner = `

	╔═╗╦═╗╔═╗╦
	╠═╣╠╦╝╠═╝║
	╩ ╩╩╚═╩  ╩
`

func main() {

	if os.Geteuid() != 0 {
		log.Fatal("arpi must run as root")
	}
	net := flag.String("net", "192.168.1.0/24", "Network")
	ui.App(*net)

	// fmt.Println(banner)

	// scanner := netscan.InitScanner(*net)

	// go scanner.LogHandler()
	// go scanner.SigHandler()
	// scanner.SigChan <- "scan"

	// <-exit

}
