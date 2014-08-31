package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/vagababov/finmath/go/finmath"
)

func exPMT(fv bool) {
	// If fv is true, we'll use PMTF instead of PMT.
	var pv, periods, rate float64
	if fv {
		fmt.Println("Enter FV, # of periods and rate per period")
	} else {
		fmt.Println("Enter PV, # of periods and rate per period")
	}
	n, err := fmt.Scanf("%f %f %f", &pv, &periods, &rate)
	if n != 3 || err != nil {
		fmt.Println("Incorrect input: ", err)
		return
	}
	if fv {
		pmt := finmath.PMTF(pv, periods, rate)
		fmt.Printf("PMTF: %f RTC(PMTF): %.2f\n", pmt, finmath.RTC(pmt))
	} else {
		pmt := finmath.PMT(pv, periods, rate)
		fmt.Printf("PMT: %f RTC(PMT): %.2f\n", pmt, finmath.RTC(pmt))
	}
}

func exPMTFS() {
	var fv, sv, periods, rate float64
	fmt.Println("Enter FV, SV, # of periods and rate per period")
	n, err := fmt.Scanf("%f %f %f %f", &fv, &sv, &periods, &rate)
	if n != 4 || err != nil {
		fmt.Println("Incorrect input: ", err)
		return
	}
	pmt := finmath.PMTFS(fv, sv, periods, rate)
	fmt.Printf("PMTFS: %f RTC(PMTFS): %.2f\n", pmt, finmath.RTC(pmt))
}

func help() {
	fmt.Println("Executes various financial math functions. Enter value separated by spaces.")
	fmt.Println("\tquit: to exit\n\tPMT for periodic payment")
	fmt.Println("\tPMTFS for periodic payment for sinking fund with starting value")
	fmt.Println("\tPMTF for periodic payment for sinking fund\n")
}

func main() {
	for {
		fmt.Println("Enter function to execute or help for list or quit to exit")
		var cmd string
		_, err := fmt.Scanln(&cmd)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("error reading input")
			continue
		}
		switch strings.ToLower(cmd) {
		case "quit":
		case "exit":
			return
		case "help":
			help()
		case "pmt":
			exPMT(false)
		case "pmtfs":
			exPMTFS()
		case "pmtf":
			exPMT(true)
		default:
			fmt.Println("unknown command: ", cmd)
		}
	}
}
