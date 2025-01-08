package main

import (
	"fmt"
	"image/color"
	"image/gif"
	"os"
	"time"

	"gobox/internal/utils"
)

// const chars = "@%#*+=-:. "
// const chars = " .,:ilwW"

const chars = "   :danD"

func main() {
	file, err := os.Open("resources/momo.gif")
	utils.CheckPanic(err)
	defer file.Close()

	g, err := gif.DecodeAll(file)
	utils.CheckPanic(err)

	frames := make([]string, 0)

	for _, i := range g.Image {
		frame := ""

		for y := i.Rect.Min.Y; y < i.Rect.Max.Y; y++ {
			for x := i.Rect.Min.X; x < i.Rect.Max.X; x++ {
				frame += pixelToASCII(i.At(x, y))
			}

			frame += "\n"
		}

		frames = append(frames, frame)
	}

	utils.ClearConsole()
	for {
		for _, frame := range frames {
			fmt.Println(frame)
			utils.MoveCursor(0, 0)
			time.Sleep(1000 / 60 * time.Millisecond)
		}
	}
}

func pixelToASCII(pixel color.Color) string {
	r, g, b, _ := pixel.RGBA()

	gray := 0.3*float64(r/256) + 0.6*float64(g/256) + 0.1*float64(b/256)
	scale := gray / 255
	index := int(scale * float64(len(chars)-1))

	return string(chars[index])
}
