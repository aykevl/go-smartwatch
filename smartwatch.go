// Package smartwatch provides a common interface for multiple smartwatches.
//
// It was originally developed to be used with the PineTime64 smartwatch.
package smartwatch

// ChargeStatus returns the state of the battery.
type ChargeStatus uint8

// Charge status of the battery: discharging, charging, and fully charged.
const (
	Discharging ChargeStatus = iota + 1
	Charging
	FullyCharged
)

type voltagePercentPosition struct {
	millivolts int
	percent    int
}

// voltageToPercent calculates the percentage the battery is full based on a
// linear approximation with multiple points on the graph. The points must be in
// order, from full to empty. The first entry in the slice must have
// percent==100, the last entry must have percent==0.
func voltageToPercent(millivolts int, pointsOnGraph []voltagePercentPosition) int {
	if millivolts >= pointsOnGraph[0].millivolts {
		return 100
	}
	if millivolts <= pointsOnGraph[len(pointsOnGraph)-1].millivolts {
		return 0
	}
	for i := 0; i < len(pointsOnGraph)-1; i++ {
		if millivolts < pointsOnGraph[i+1].millivolts {
			continue
		}
		// Voltage is between pointsOnGraph[i] and pointsOnGraph[i+1].
		high := pointsOnGraph[i]
		low := pointsOnGraph[i+1]
		return high.percent + (high.percent-low.percent)*(millivolts-high.millivolts)/(high.millivolts-low.millivolts)
	}
	// unreachable
	return 0
}
