package main

import (
	"fmt"
	"image/color"
	"image/gif"
	"os"
	"time"
)

// const chars = "@%#*+=-:. "
const chars = "   :danD"

// const chars = " .,:ilwW"

func main() {
	file, err := os.Open("resources/momo.gif")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	g, err := gif.DecodeAll(file)
	if err != nil {
		panic(err)
	}

	frames := make([]string, 0)
	for _, i := range g.Image {
		frame := ""

		for y := i.Rect.Min.Y; y < i.Rect.Max.Y; y++ {
			for x := i.Rect.Min.X; x < i.Rect.Max.X; x++ {
				color := i.At(x, y)
				char := pixelToASCII(color)
				frame += char
			}

			frame += "\n"
		}

		frames = append(frames, frame)
	}

	for {
		for _, frame := range frames {
			fmt.Println(frame)
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func pixelToASCII(pixel color.Color) string {
	r, g, b, _ := pixel.RGBA()
	gray := uint8(0.299*float64(r/256) + 0.587*float64(g/256) + 0.114*float64(b/256))
	scale := float64(gray) / 255.0
	index := int(scale * float64(len(chars)-1))
	return string(chars[index])
}
