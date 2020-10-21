package main

import (
	"fmt"
	"image"
	"image/color"
)

// Image image type
type Image struct {
	w int
	h int
}

// ColorModel return the color model
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds return the bounds of the image
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

// At return color at a point
func (i Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	// m := Image{100, 100}
	fmt.Println("Hello!")
}
