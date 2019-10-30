// +build !baremetal

package smartwatch

// This file implements the watch interface for desktop systems, using SDL2. It
// is meant for quick testing and simulating of watch programs.

import (
	"github.com/aykevl/tilegraphics/sdlscreen"
)

// Watch is an emulation for smartwatches on Linux systems.
type Watch struct {
	*sdlscreen.Screen
}

var watch *Watch

// Open returns a Watch instance. It is a singleton: opening it a second time
// will still return the same object, if opening the first time succeeded.
func Open() (*Watch, error) {
	if watch != nil {
		return watch, nil
	}
	screen, err := sdlscreen.NewScreen("smartwatch", 240, 240)
	if err != nil {
		return nil, err
	}
	watch = &Watch{
		Screen: screen,
	}
	return watch, nil
}

// BatteryStatus reads and returns the current battery status (percent and
// whether it is charging).
func (w *Watch) BatteryStatus() (millivolt, percent int, status ChargeStatus) {
	return 3700, 75, Discharging
}
