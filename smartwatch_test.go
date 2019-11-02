package smartwatch

import "testing"

func TestVoltageToPercent(t *testing.T) {
	t.Parallel()

	// The voltage/percentage mappings we'll use for this test. Originally from
	// a PineTime discharge curve.
	var voltagePercentPositions = []voltagePercentPosition{
		{3880, 100},
		{3780, 80},
		{3690, 60},
		{3640, 40},
		{3610, 20},
		{3520, 0},
	}

	// Test whether the values are contiguous and don't have weird bumps in the chart.
	lastPercent := 100
	for millivolts := 3900; millivolts >= 3500; millivolts-- {
		percent := voltageToPercent(millivolts, voltagePercentPositions)
		if percent < 0 || percent > 100 {
			t.Errorf("%dmV maps to %d%% which is outside the range 0-100", millivolts, percent)
		}
		if percent > lastPercent {
			t.Errorf("%d%% is higher than the last percent %d%% while the voltage is lower (%dmV -> %dmV)", percent, lastPercent, millivolts+1, millivolts)
		}
		lastPercent = percent

		// For outputting JSON (for debugging):
		// Can be useful to put them in a chart.
		//print("[", i, ", ", voltageToPercent(i, voltagePercentPositions), "],\n")
	}

	// Test whether the calculated values are as expected.
	for millivolts, percent := range map[int]int{
		3881: 100,
		3880: 100,
		3879: 100,
		3691: 61,
		3690: 60,
		3689: 60,
		3630: 34,
		3521: 1,
		3520: 0,
		3519: 0,
		3503: 0,
	} {
		calculatedPercent := voltageToPercent(millivolts, voltagePercentPositions)
		if percent != calculatedPercent {
			t.Errorf("expected %d%% for voltage %dmV, got %d%%", percent, millivolts, calculatedPercent)
		}
	}
}
