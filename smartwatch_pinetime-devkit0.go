// +build pinetime_devkit0

package smartwatch

// This file implements the watch interface for the PineTime:
// https://wiki.pine64.org/index.php/PineTime

import (
	"machine"

	"tinygo.org/x/drivers/st7789"
)

// Watch implements the PineTime64 developer kit.
type Watch struct {
	*st7789.Device
}

var watch *Watch

// The two pins connected to the power regulator chip, that indicate charging
// status and power presence.
const (
	pinBatteryCharging machine.Pin = 12
	pinPowerConnected  machine.Pin = 19
)

// Voltage to percent mappings. Values in between can be linearly approximated.
// This is a rough fitting, better fits are likely possible.
var voltagePercentPositions = []voltagePercentPosition{
	{3880, 100},
	{3780, 80},
	{3690, 60},
	{3640, 40},
	{3610, 20},
	{3520, 0},
}

// Open returns a Watch instance. It is a singleton: opening it a second time
// will still return the same object, if opening the first time succeeded.
func Open() (*Watch, error) {
	if watch != nil {
		return watch, nil
	}

	// Configure the screen.
	spi := machine.SPI0
	spi.Configure(machine.SPIConfig{
		MOSI:      machine.SPI0_MOSI_PIN,
		MISO:      machine.SPI0_MISO_PIN,
		SCK:       machine.SPI0_SCK_PIN,
		Frequency: 8000000,
		Mode:      3,
	})
	const (
		resetPin = machine.LCD_RESET
		dcPin    = machine.LCD_RS
		blPin    = machine.LCD_BACKLIGHT_MID
	)
	screen := st7789.New(spi, resetPin, dcPin, blPin)
	watch = &Watch{
		Device: &screen,
	}
	machine.LCD_CS.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.LCD_CS.Low()
	screen.Configure(st7789.Config{
		Width:  240,
		Height: 240,
	})
	screen.EnableBacklight(false) // enables the backlight

	// Configure the battery status pins.
	pinBatteryCharging.Configure(machine.PinConfig{Mode: machine.PinInput})
	pinPowerConnected.Configure(machine.PinConfig{Mode: machine.PinInput})

	return watch, nil
}

// BatteryStatusRaw reads and returns the current battery voltage in millivolts
// (mV) and returns the current charging status of the battery (discharging,
// charging, full).
func (w *Watch) BatteryStatusRaw() (millivolt int, status ChargeStatus) {
	if !pinPowerConnected.Get() {
		// Power is connected.
		if !pinBatteryCharging.Get() {
			// Battery is charging.
			status = Charging
		} else {
			status = FullyCharged
		}
	} else {
		status = Discharging
	}
	value := machine.ADC{31}.Get()
	return int(value) * 2000 / (65535 / 3), status
}

// BatteryStatus reads and returns the current battery status (percent and
// charge status).
func (w *Watch) BatteryStatus() (millivolts, percent int, status ChargeStatus) {
	millivolts, status = w.BatteryStatusRaw()
	percent = voltageToPercent(millivolts, voltagePercentPositions)
	return
}
