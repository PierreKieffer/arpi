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
	scanner.PingScan()
	scanner.OSDetection()
}

func (scanner *Scanner) PingScan() {
	scanner.LogChan <- "Network scan in progress ... "

	cmd := exec.Command("sudo", "nmap", "-sP", scanner.Network, "-oX", "/tmp/arpi.xml")

	var out bytes.Buffer

	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	scanner.ProcessPingScan()

	fmt.Println(scanner)

	scanner.LogChan <- "Network scan completed"
}

func (scanner *Scanner) OSDetection() {

	for _, host := range scanner.Devices {

		outputFile := fmt.Sprintf("/tmp/arpi_%v.xml", host.IP)

		cmd := exec.Command("sudo", "nmap", "-O", "-p", "80", host.IP, "-oX", outputFile)

		var out bytes.Buffer

		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		host.ProcessOSDetection()

	}
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

func (scanner *Scanner) ProcessPingScan() {
	outputData, err := ExtractPingScanData()
	if err != nil {
		log.Fatal(err)
	}
	scanner.Time = outputData.Startstr
	scanner.Summary = outputData.Runstats.Finished.Summary

	var devices []Device
	for _, host := range outputData.Host {

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

func (device *Device) ProcessOSDetection() {

	fmt.Println(" -------------------------------- ")

	outputFile := fmt.Sprintf("/tmp/arpi_%v.xml", device.IP)

	outputData, err := ExtractScanOSData(outputFile)

	if err != nil {
		log.Fatal(err)
	}

	for _, os := range outputData.Host.Os.Osmatch {
		fmt.Println(os.Osclass)

	}

}
