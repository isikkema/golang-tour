package main

import (
	"golang.org/x/tour/pic"
	"math"
)

func dist(x1, y1, x2, y2 int) (d float64) {
	d = math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
	return
}

func Pic(dx, dy int) [][]uint8 {
	rv := make([][]uint8, dy)
	for y := range rv {
		rv[y] = make([]uint8, dx)
		for x := range rv[y] {
			mid := len(rv)/2
			if dist(x, y, mid, mid) <= 127 {
				rv[y][x] = uint8(255-2*dist(x, y, mid, mid))
			}
		}
	}
	
	return rv
}

func main() {
	pic.Show(Pic)
}
