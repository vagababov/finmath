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
		fmt.Printf("PMTF: %f RTC(PMTF): %f\n", pmt, finmath.RTC(pmt))
	} else {
		pmt := finmath.PMT(pv, periods, rate)
		fmt.Printf("PMT: %f RTC(PMT): %f\n", pmt, finmath.RTC(pmt))
	}
}

func help() {
	fmt.Println("Executes various financial math functions. Enter value separated by space")
	fmt.Println("\tquit: to exit\n\tPMT for periodic payment")
	fmt.Println("\tPMTF for periodic payment for sinking fund\n")
}

func main() {
	for {
		fmt.Println("enter function to execute or help for list or quit to exit")
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
			return
		case "help":
			help()
		case "pmt":
			exPMT(false)
		case "pmtf":
			exPMT(true)
		default:
			fmt.Println("unknown command: ", cmd)
		}
	}
}
