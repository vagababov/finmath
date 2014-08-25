// Contains unit tests for the finmath package.
package finmath

import "testing"

func TestCompoundInterest(t *testing.T) {
	if got, want := CompoundInterest(1000, 2, 0.2), 1440.; got != want {
		t.Errorf("CompoundInterest: got: %v want: %v", got, want)
	}
}

func TestPV(t *testing.T) {
	if got, want := RTC(PV(1200, 15*12, 0.045/12)), 156864.12; got != want {
		t.Errorf("RTC(PV): got: %v want: %v", got, want)
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
			t.Errorf("RTC(%V): got: %v want: %v", test.v, got, want)
		}
		if got, want := FTC(test.v), test.wantFTC; got != want {
			t.Errorf("FTC(%V): got: %v want: %v", test.v, got, want)
		}
	}
}
