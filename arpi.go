package main

import (
	"flag"
	"github.com/PierreKieffer/arpi/pkg/nmap"
)

var exit = make(chan bool)

func main() {
	net := flag.String("net", "192.168.1.0/24", "Network")
	var scanner = &nmap.Scanner{}
	scanner.Network = *net

	go scanner.ListenScanStatus()
	go scanner.Scan()

	<-exit

}
