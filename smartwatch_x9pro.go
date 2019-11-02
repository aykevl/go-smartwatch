// +build x9pro

package smartwatch

import (
	"machine"

	"tinygo.org/x/drivers/ssd1331"
)

// Watch implements the PineTime64 developer kit.
type Watch struct {
	*ssd1331.Device
}

var watch *Watch

// Open returns a Watch instance. It is a singleton: opening it a second time
// will still return the same object, if opening the first time succeeded.
func Open() (*Watch, error) {
	if watch != nil {
		return watch, nil
	}

	// Configure the screen.
	machine.OLED_LED_POW.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.OLED_LED_POW.Low()
	spi := machine.SPI0
	spi.Configure(machine.SPIConfig{
		MOSI:      machine.OLED_MOSI,
		MISO:      machine.NoPin,
		SCK:       machine.OLED_SCK,
		Frequency: 8000000,
	})
	const (
		resetPin = machine.OLED_RES
		dcPin    = machine.OLED_DC
		csPin    = machine.OLED_CS
	)
	display := ssd1331.New(spi, resetPin, dcPin, csPin)
	display.Configure(ssd1331.Config{})
	display.SetContrast(0x30, 0x20, 0x30)

	watch = &Watch{
		Device: &display,
	}

	return watch, nil
}
