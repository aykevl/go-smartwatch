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

// Open returns a Watch instance. It is a singleton: opening it a second time
// will still return the same object, if opening the first time succeeded.
func Open() (*Watch, error) {
	if watch != nil {
		return watch, nil
	}
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
	return watch, nil
}
