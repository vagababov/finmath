package finamth

import "math"

func CompoundInterest(start, periods, rate float64) float64 {
	return start * math.Pow(1.+rate, periods)
}

// PV (annuity) returns the required amount to put down in order to receive
// rent for period times with given rate.
// if start == 1, then interest is accrued at the start of the period.
// E.g. to receive 1200 monthly over 15 years on an instrument that earns 4.5%
// annually, you have to pay now:
// PV(1200, 15*12, 0.045/12)
func PV(rent, periods, rate float64) float64 {
	return rent / rate * (1. - math.Pow(1.+rate, -periods))
}

// FV returns the future value of an instrument earning rate over periods
// paying rent per period.
func FV(rent, periods, rate float64) float64 {
	return (math.Pow(1.+rate, periods) - 1) / rate * rent
}

// PMT (periodic payment) returns the payment amount to repay "sum" over
// "periods" with "rate".
// For example, PMT is the payment you have to make for a fixed-rate
// mortgage loan.
// For 30-year loan, paid monthly with 4.5% rate: periods = 30*12 = 360
// rate = 0.045/12
func PMT(sum, periods, rate float64) float64 {
	return sum * rate / (1. - math.Pow(1.+rate, -periods))
}