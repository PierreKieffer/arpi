package netscan

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

type Scanner struct {
	Network string
	Time    string
	Summary string
	Devices []Device
	LogChan chan string
	SigChan chan string
}

type Device struct {
	IP   string
	MAC  string
	Name string
}

func InitScanner(network string) *Scanner {
	var scanner = &Scanner{}
	scanner.Network = network
	scanner.InitChan()
	return scanner
}

func (scanner *Scanner) InitChan() {
	if scanner.LogChan == nil {
		scanner.LogChan = make(chan string)
	}

	if scanner.SigChan == nil {
		scanner.SigChan = make(chan string)
	}
}

func (scanner *Scanner) Scan() {

	scanner.LogChan <- "Network scan in progress ... "

	cmd := exec.Command("sudo", "nmap", "-sP", scanner.Network, "-oX", "/tmp/arpi.xml")

	var out bytes.Buffer

	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	scanner.ProcessScanOutput()

	fmt.Println(scanner)

	scanner.LogChan <- "Network scan completed"
}

func (scanner *Scanner) SigHandler() {

	for {
		select {
		case <-scanner.SigChan:
			scanner.Scan()
		}
	}
}

func (scanner *Scanner) LogHandler() {

	for {
		select {
		case status := <-scanner.LogChan:
			fmt.Println(status)
		}
	}
}

func (scanner *Scanner) ProcessScanOutput() {
	scanOutputData, err := ExtractOutputData()
	if err != nil {
		log.Fatal(err)
	}

	scanner.Time = scanOutputData.Startstr
	scanner.Summary = scanOutputData.Runstats.Finished.Summary

	var devices []Device
	for _, host := range scanOutputData.Host {

		var device Device
		device.Name = host.Hostnames.Hostname.Name

		for _, add := range host.Address {
			if add.Addrtype == "ipv4" || add.Addrtype == "ipv6" {
				device.IP = add.Addr
			}
			if add.Addrtype == "mac" {
				device.MAC = add.Addr
			}
		}

		devices = append(devices, device)
	}
	scanner.Devices = devices
}
