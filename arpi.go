package main

import (
	"flag"
	"github.com/PierreKieffer/arpi/pkg/nmap"
)

func main() {
	net := flag.String("net", "192.168.1.0/24", "Network")

	nmap.Scan(*net)

}
