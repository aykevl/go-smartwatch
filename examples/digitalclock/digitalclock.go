// This program displays the current time on the screen.
package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/aykevl/go-smartwatch"
	"github.com/conejoninja/tinyfont"
	"github.com/conejoninja/tinyfont/freemono"
)

var watch *smartwatch.Watch

func main() {
	watch, _ = smartwatch.Open()
	width, height := watch.Size()

	// Pick an appropriate font.
	font := &tinyfont.Org01 // fallback font
	fonts := []*tinyfont.Font{&freemono.Bold9pt7b, &freemono.Bold12pt7b, &freemono.Bold18pt7b, &freemono.Bold24pt7b}
	for _, f := range fonts {
		// If the font fits on this screen, use it.
		lineWidth, _ := tinyfont.LineWidth(f, []byte("00:00"))
		if int16(lineWidth) <= width {
			font = f
		}
	}
	fontHeight := int16(font.Glyphs['0'-font.First].Height)

	// Draw the current time.
	for {
		// Clear the screen.
		watch.FillScreen(color.RGBA{0, 0, 0, 255})

		// Draw the current time (with second precision).
		now := time.Now()
		// Quick-and-dirty hack to get the current time (roughly) without
		// relying on locale support, which is not yet supported in TinyGo.
		hour := (now.Unix() / 60 / 60) % 24
		minute := now.Unix() / 60 % 60
		msg := []byte(fmt.Sprintf("%02d:%02d", hour, minute))
		textWidth, _ := tinyfont.LineWidth(font, msg)
		tinyfont.WriteLine(watch, font, width/2-int16(textWidth/2), height/2+fontHeight/2, msg, color.RGBA{255, 255, 255, 255})

		watch.Display()

		// Sleep until the next minute.
		time.Sleep(time.Minute - time.Duration(now.Nanosecond())*time.Nanosecond)
	}
}
