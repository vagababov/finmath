// Contains unit tests for the finmath package.
package finmath

import (
	"math"
	"testing"
)

func TestCompoundInterest(t *testing.T) {
	if got, want := CompoundInterest(1000, 2, 0.2), 1440.; got != want {
		t.Errorf("CompoundInterest: got: %v want: %v", got, want)
	}
}

func TestFV(t *testing.T) {
	if got, want := RTC(FV(1596.18, 360, 0.03625/12)), 1036696.15; got != want {
		t.Errorf("RTC(FV): got: %f want: %f", got, want)
	}
	// Verify PV and FV relationship holds.
	if got, want := RTC(FV(1596.18, 360, 0.03625/12)), RTC(PV(1596.18, 360, 0.03625/12)*math.Pow(1+0.03625/12, 360)); got != want {
		t.Errorf("FV != PV*(1+i)^n: got: %f want: %f", got, want)
	}
}

func TestPV(t *testing.T) {
	if got, want := RTC(PV(1200, 15*12, 0.045/12)), 156864.12; got != want {
		t.Errorf("RTC(PV): got: %f want: %f", got, want)
	}
}

func TestPMT(t *testing.T) {
	if got, want := RTC(PMT(350000, 360, 0.03625/12)), 1596.18; got != want {
		t.Errorf("RTV(PMT): got: %f want: %f", got, want)
	}
}

func TestRoundting(t *testing.T) {
	tests := []struct {
		v                float64
		wantRTC, wantFTC float64
	}{
		{1.0, 1.0, 1.0},
		{1.1, 1.1, 1.1},
		{1.11, 1.11, 1.11},
		{1.111, 1.11, 1.11},
		{1.113, 1.11, 1.11},
		{1.115, 1.12, 1.11},
		{1.119, 1.12, 1.11},
		{1.119999, 1.12, 1.11},
	}

	for _, test := range tests {
		if got, want := RTC(test.v), test.wantRTC; got != want {
			t.Errorf("RTC(%f): got: %f want: %f", test.v, got, want)
		}
		if got, want := FTC(test.v), test.wantFTC; got != want {
			t.Errorf("FTC(%f): got: %f want: %f", test.v, got, want)
		}
	}
}
