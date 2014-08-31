// Package finmath implements routines for engineering economics computations.
// Those include mortgate payments, money value change in time, etc.
// Note, that it uses float64 so it should be used with reservation for computations that require
// arbitrary precision.
package finmath

import "math"

// CompoundInterest returns the value of the investment with start value after periods accruing
// interest with rate.
func CompoundInterest(start, periods, rate float64) float64 {
	return start * math.Pow(1.+rate, periods)
}

// PV (annuity) returns the required amount to put down in order to receive
// rent for period times with given rate. A.k.a Series present worth.
// if start == 1, then interest is accrued at the start of the period.
// E.g. to receive 1200 monthly over 15 years on an instrument that earns 4.5%
// annually, you have to pay now:
// PV(1200, 15*12, 0.045/12)
func PV(rent, periods, rate float64) float64 {
	return rent / rate * (1. - math.Pow(1.+rate, -periods))
}

// FV returns the future value of an instrument earning rate over periods
// paying rent per period. A.k.a series compound amount.
func FV(rent, periods, rate float64) float64 {
	return (math.Pow(1.+rate, periods) - 1) / rate * rent
}

// PVG returns the required
func PVG(gradient, periods, rate float64) float64 {
	return 0
}

// PMT (periodic payment) returns the payment amount to repay "pv" over
// "periods" with "rate". A.k.a capital recovery.
// For example, PMT is the payment you have to make for a fixed-rate
// mortgage loan.
// For 30-year loan, paid monthly with 4.5% rate: periods = 30*12 = 360
// rate = 0.045/12
func PMT(pv, periods, rate float64) float64 {
	return pv * rate / (1. - math.Pow(1.+rate, -periods))
}

// PMTF returns the required payments to make for given number of periods earning
// rate each period  to reach fv after periods elapsed. A.k.a Sinking fund.
func PMTF(fv, periods, rate float64) float64 {
	return fv * rate / (math.Pow(1.+rate, periods) - 1)
}

// RTC rounds the given amount to the nearest cent.
func RTC(v float64) float64 {
	return math.Floor(v*100+0.5) / 100.
}

// FTC rounds down the given amount to the cent.
func FTC(v float64) float64 {
	return math.Trunc(v*100) / 100
}
