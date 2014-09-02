// Package finmath implements routines for engineering economics computations.
// Those include mortgate payments, money value change in time, etc.
// Note, that it uses float64 so it should be used with reservation for computations that require
// arbitrary precision.
// This package is provided as is without warranty and should be used at your own risk.
package finmath

import "math"

// CompoundInterest returns the value of the investment with start value after periods accruing
// interest with rate. A.k.a. Single Payment Compound Amount (F/P, i, n)
// To find reverse, Single Payment Present Worth (P/F, i, n) -- just use -periods.
func CompoundInterest(start, periods, rate float64) float64 {
	return start * math.Pow(1.+rate, periods)
}

// PV (annuity) returns the required amount to put down in order to receive
// rent for period times with given rate. A.k.a Series present worth (P/A, i, n).
// E.g. to receive 1200 monthly over 15 years on an instrument that earns 4.5%
// annually, you have to pay now:
// PV(1200, 15*12, 0.045/12)
func PV(rent, periods, rate float64) float64 {
	return rent / rate * (1. - math.Pow(1.+rate, -periods))
}

// FV returns the future value of an instrument earning rate over periods
// paying rent per period. A.k.a series compound amount (F/A, i, n).
func FV(rent, periods, rate float64) float64 {
	return (math.Pow(1.+rate, periods) - 1) / rate * rent
}

// PMTG returns the required fixed payment to offset constantly growing maintenance costs, a.k.a arithmetic gradient uniform series (A/G, i, n).
// Or, annual worth of an arithmetic gradient.
// E.g. if each period the maintenance costs are growing by 250 and money is earning rate 1%
// each period for 10 years, then payment is PVG(250, 10, 0.01)
func PMTG(gradient, periods, rate float64) float64 {
	x := math.Pow(1.+rate, periods)
	return gradient * ((x - rate*periods - 1) / (rate*x - rate))
}

// PVG returns the present worth of an uniformly increasing/decreasing value, a.k.a. arithmetic gradient present worth (P/G, i, n)
func PVG(gradient, periods, rate float64) float64 {
	x := math.Pow(1.+rate, periods)
	return gradient * ((x - rate*periods - 1) / (rate * rate * x))
}

// PMT (periodic payment) returns the payment amount to repay "pv" over
// "periods" with "rate". A.k.a capital recovery (A/P, i, n).
// For example, PMT is the payment you have to make for a fixed-rate
// mortgage loan.
// For 30-year loan, paid monthly with 4.5% rate: periods = 30*12 = 360
// rate = 0.045/12
func PMT(pv, periods, rate float64) float64 {
	return pv * rate / (1. - math.Pow(1.+rate, -periods))
}

// PMTF returns the required payments to make for given number of periods earning
// rate each period  to reach fv after periods elapsed. A.k.a Sinking fund (A/F, i, n).
func PMTF(fv, periods, rate float64) float64 {
	return fv * rate / (math.Pow(1.+rate, periods) - 1)
}

// PMTFS calculates PMTF where there's also starting amount earning same rate.
func PMTFS(fv, sv, periods, rate float64) float64 {
	// 1. Calculate how much sv will earn in that period of time.
	sve := CompoundInterest(sv, periods, rate)
	// 2. Return amount required to put each year to match the difference.
	return PMTF(fv-sve, periods, rate)
}

// RTC rounds the given amount to the nearest cent.
func RTC(v float64) float64 {
	return math.Floor(v*100+0.5) / 100.
}

// FTC rounds down the given amount to the cent.
func FTC(v float64) float64 {
	return math.Trunc(v*100) / 100
}

