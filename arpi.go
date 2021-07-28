package main

import (
	"flag"
	"github.com/PierreKieffer/arpi/pkg/ui"
	"log"
	"os"
)

var banner = `

	╔═╗╦═╗╔═╗╦
	╠═╣╠╦╝╠═╝║
	╩ ╩╩╚═╩  ╩
`

func main() {

	if os.Geteuid() != 0 {
		log.Fatal("ERROR : arpi must run as root")
	}
	net := flag.String("net", "192.168.1.0/24", "Network")
	flag.Parse()
	ui.App(*net)
}
