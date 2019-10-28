// This program displays some simple graphics on the screen, as a "hello world"
// style program.
package main

import (
	"image/color"
	"time"

	"github.com/aykevl/go-smartwatch"
	"github.com/conejoninja/tinyfont"
)

var watch *smartwatch.Watch

func main() {
	watch, _ = smartwatch.Open()

	// Draw a square in the middle of the screen.
	const size = 50
	width, height := watch.Size()
	watch.FillScreen(color.RGBA{0, 0, 0, 255})
	watch.FillRectangle(width/2-size/2, height/2-size/2, size, size, color.RGBA{255, 255, 0, 255})

	// Draw some text on the screen.
	msg := []byte("Hello watch!")
	textWidth, _ := tinyfont.LineWidth(&tinyfont.Org01, msg)
	tinyfont.WriteLine(watch, &tinyfont.Org01, width/2-int16(textWidth/2), 30, msg, color.RGBA{255, 0, 0, 255})

	watch.Display()

	for {
		time.Sleep(time.Hour)
	}
}
