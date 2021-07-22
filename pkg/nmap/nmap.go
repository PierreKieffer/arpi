package nmap

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Scanner struct {
	Network    string
	Time       string
	StatusLine string
	Devices    []Device
	LogChan    chan string
	SigChan    chan string
}

type Device struct {
	IP   string
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

	cmd := exec.Command("nmap", "-sn", scanner.Network)

	var out bytes.Buffer

	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	scanner.ParseScannerOutput(out.String())

	fmt.Println(scanner)

	scanner.LogChan <- "Network scan completed"
}

func (scanner *Scanner) SigHandler() {

	for {
		signal := <-scanner.SigChan
		switch signal {
		case "scan":
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

func (scanner *Scanner) ParseScannerOutput(scannerOutput string) error {
	outArray := strings.Split(scannerOutput, "\n")
	outArray = outArray[1 : len(outArray)-1]

	header := outArray[0]
	statusLine := outArray[len(outArray)-1]
	time := ExtractTimestamp(header)

	scanner.Time = time
	scanner.StatusLine = statusLine

	var netData []string
	for _, v := range outArray[1 : len(outArray)-1] {
		if strings.Contains(v, "Nmap scan report") {
			netData = append(netData, v)
		}

	}

	scanner.ProcessNetData(netData)

	return nil

}

func (scanner *Scanner) ProcessNetData(netData []string) {

	var devices []Device

	for _, v := range netData {
		device := ExtractDeviceData(v)
		devices = append(devices, device)
	}

	scanner.Devices = devices
}

func ExtractTimestamp(scanHeader string) string {
	splitHeader := strings.Split(scanHeader, "at ")
	time := splitHeader[len(splitHeader)-1]
	return time
}

func ExtractDeviceData(rawDeviceScan string) Device {
	rawDeviceScanFmt := strings.Replace(rawDeviceScan, "Nmap scan report for ", "", -1)
	deviceName := strings.Split(rawDeviceScanFmt, " ")[0]
	deviceIP := strings.Split(rawDeviceScanFmt, " ")[1]
	device := Device{
		IP:   deviceIP,
		Name: deviceName,
	}

	return device
}
