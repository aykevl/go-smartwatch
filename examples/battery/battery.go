// This program displays some simple graphics on the screen, as a "hello world"
// style program.
package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/aykevl/go-smartwatch"
	"github.com/conejoninja/tinyfont"
)

var watch *smartwatch.Watch

func main() {
	watch, _ = smartwatch.Open()

	// Draw a square in the middle of the screen.
	for {
		watch.FillScreen(color.RGBA{0, 0, 0, 255})

		millivolts, batteryPercent, chargingStatus := watch.BatteryStatus()
		chargingString := ""
		switch chargingStatus {
		case smartwatch.FullyCharged:
			chargingString = " fully charged"
		case smartwatch.Charging:
			chargingString = " charging"
		}
		msg := fmt.Sprintf("battery: %d%% (%.3fV)%s", batteryPercent, float32(millivolts)/1000, chargingString)
		tinyfont.WriteLine(watch, &tinyfont.Org01, 10, 20, []byte(msg), color.RGBA{255, 0, 0, 255})

		width, height := watch.Size()

		// Background
		watch.FillRectangle(width/2-54, height/2-29, 108, 58, color.RGBA{100, 100, 100, 255})
		watch.FillRectangle(width/2+54, height/2-16, 10, 32, color.RGBA{100, 100, 100, 255})

		// Fuel
		fuelColor := color.RGBA{0, 255, 0, 255} // green
		if batteryPercent <= 25 {
			fuelColor = color.RGBA{255, 127, 0, 255} // orange
		}
		if batteryPercent <= 10 {
			fuelColor = color.RGBA{255, 0, 0, 255} // red
		}
		watch.FillRectangle(width/2-50, height/2-25, int16(batteryPercent), 50, fuelColor)

		watch.Display()

		//semihosting.Stdout.Write([]byte(msg + "\n"))

		time.Sleep(time.Second * 10)
	}
}
