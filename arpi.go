package main

import (
	"flag"
	"fmt"
	"github.com/PierreKieffer/arpi/pkg/netscan"
)

var exit = make(chan bool)

var banner = `

	╔═╗╦═╗╔═╗╦
	╠═╣╠╦╝╠═╝║
	╩ ╩╩╚═╩  ╩
`

func main() {

	fmt.Println(banner)

	net := flag.String("net", "192.168.1.0/24", "Network")

	scanner := netscan.InitScanner(*net)

	go scanner.LogHandler()
	go scanner.SigHandler()
	scanner.SigChan <- "scan"

	<-exit

}
