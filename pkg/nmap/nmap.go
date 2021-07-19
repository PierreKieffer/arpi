package nmap

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type ScanData struct {
	Time    string
	Devices []Device
}

type Device struct {
	IP   string
	Name string
}

func Scan(net string) {
	cmd := exec.Command("nmap", "-sn", net)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	ParseScannerOutput(out.String())
}

func ParseScannerOutput(scannerOutput string) (*ScanData, error) {
	outArray := strings.Split(scannerOutput, "\n")
	outArray = outArray[1 : len(outArray)-1]

	var scanData ScanData

	header := outArray[0]
	footer := outArray[len(outArray)-1]

	time := ExtractTimestamp(header)

	fmt.Println(time)
	fmt.Println(footer)

	var netData []string
	for _, v := range outArray[1 : len(outArray)-1] {
		if strings.Contains(v, "Nmap scan report") {
			netData = append(netData, v)
		}

	}

	fmt.Println(netData)

	return &scanData, nil

}

func ProcessNetData(netData []string) {

	var devices []Device

	for _, v := range netData {
		/*
			Nmap scan report for raspberrypi (192.168.1.68)
		*/
		rawDeviceScan := strings.Replace(v, "Nmap scan report for ", "", -1)
		deviceName := strings.Split(rawDeviceScan, " ")[0]
		deviceIP := strings.Split(rawDeviceScan, " ")[1]
	}

}

func ExtractTimestamp(scanHeader string) string {
	splitHeader := strings.Split(scanHeader, "at ")
	time := splitHeader[len(splitHeader)-1]
	return time
}

func ExtractDeviceData(rawDeviceScan string) Device {

}
