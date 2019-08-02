package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	"math"
)

func line(y, x, thickness float64) bool {
	return math.Abs(y-x) <= thickness
}

type Image struct{}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 200)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	red := color.RGBA{255, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}
	if ( line(float64(y), -0.02*math.Pow(float64(x)-100, 2)+190, 5) ) {
		return red
	} else {
		return white
	}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
