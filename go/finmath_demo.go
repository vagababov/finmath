package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/vagababov/finmath/go/finmath"
)

func read4Values() ([4]float64, error) {
	var v [4]float64
	n, err := fmt.Scanf("%f %f %f %f", &v[0], &v[1], &v[2], &v[3])
	if n != 4 {
		return v, fmt.Errorf("read %d need %d values", n, 4)
	}
	if err != nil {
		return v, err
	}
	return v, nil
}

func read3Values() ([3]float64, error) {
	var v [3]float64
	n, err := fmt.Scanf("%f %f %f", &v[0], &v[1], &v[2])
	if n != 3 {
		return v, fmt.Errorf("read %d need %d values", n, 3)
	}
	if err != nil {
		return v, err
	}
	return v, nil
}

func exPMT(fv bool) {
	// If fv is true, we'll use PMTF instead of PMT.
	if fv {
		fmt.Println("Enter FV, # of periods and rate per period")
	} else {
		fmt.Println("Enter PV, # of periods and rate per period")
	}
	vals, err := read3Values()
	if err != nil {
		fmt.Println("Error reading values: ", err)
		return
	}
	if fv {
		pmt := finmath.PMTF(vals[0], vals[1], vals[2])
		fmt.Printf("PMTF: %f RTC(PMTF): %.2f\n", pmt, finmath.RTC(pmt))
	} else {
		pmt := finmath.PMT(vals[0], vals[1], vals[2])
		fmt.Printf("PMT: %f RTC(PMT): %.2f\n", pmt, finmath.RTC(pmt))
	}
}

func exPMTFS() {
	fmt.Println("Enter FV, SV, # of periods and rate per period")
	vals, err := read4Values()
	if err != nil {
		fmt.Println("Error reading values: ", err)
		return
	}
	pmt := finmath.PMTFS(vals[0], vals[1], vals[2], vals[3])
	fmt.Printf("PMTFS: %f RTC(PMTFS): %.2f\n", pmt, finmath.RTC(pmt))
}

func exCI() {
	fmt.Println("Enter PV/FV, # of periods and rate per period:")
	vals, err := read3Values()
	if err != nil {
		fmt.Println("Error reading values: ", err)
		return
	}
	ci := finmath.CompoundInterest(vals[0], vals[1], vals[2])
	fmt.Printf("CI: %f RTC(CI): %.2f\n", ci, finmath.RTC(ci))
}

func exPMTG(pmt bool) {
	fmt.Println("Enter gradient amount, # of periods, rate")
	vals, err := read3Values()
	if err != nil {
		fmt.Println("Error reading values: ", err)
		return
	}
	if pmt {
		v := finmath.PMTG(vals[0], vals[1], vals[2])
		fmt.Printf("PMTG: %f RTC(PMTG): %.2f\n", v, finmath.RTC(v))
		return
	}
	v := finmath.PVG(vals[0], vals[1], vals[2])
	fmt.Printf("PVG: %f RTC(PVG): %.2f\n", v, finmath.RTC(v))
}

func exPFV(fv bool) {
	fmt.Println("Enter payment amount, # of periods, rate")
	vals, err := read3Values()
	if err != nil {
		fmt.Println("Error reading values: ", err)
		return
	}
	if fv {
		v := finmath.FV(vals[0], vals[1], vals[2])
		fmt.Printf("FV: %f RTC(FV): %.2f\n", v, finmath.RTC(v))
		return
	}
	v := finmath.PV(vals[0], vals[1], vals[2])
	fmt.Printf("PV: %f RTC(PV): %.2f\n", v, finmath.RTC(v))
}

func help() {
	fmt.Println("Executes various financial math functions. Enter value separated by spaces.")
	fmt.Println("\tquit: to exit")
	fmt.Println("\tPMT: capital recovery (A/P, i, n)")
	fmt.Println("\tPMTFS: sinking fund with starting value (A/F, i, n)")
	fmt.Println("\tPMTF: sinking fund (A/F, i, n)")
	fmt.Println("\tPMTG: arithmetic gradient uniform series (A/G, i, n)")
	fmt.Println("\tCI: compound interest (F/P, i, n) or (P/F, i, n), if n is negative")
	fmt.Println("\tFV: series compound amount (F/A, i, n)")
	fmt.Println("\tPV: series present worth (P/A, i, n)")
	fmt.Println("\tPVG: arithmetic gradient present worth (P/G, i, n)")
	fmt.Println()
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
			fallthrough
		case "exit":
			return
		case "help":
			help()
		case "ci":
			exCI()
		case "pvg":
			exPMTG(false)
		case "pmtg":
			exPMTG(true)
		case "pmt":
			exPMT(false)
		case "pmtfs":
			exPMTFS()
		case "pmtf":
			exPMT(true)
		case "pv":
			exPFV(false)
		case "fv":
			exPFV(true)
		default:
			fmt.Println("unknown command: ", cmd)
		}
	}
}

