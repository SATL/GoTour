package main

import (
	"image"
	"image/color"
	// "code.google.com/p/go-tour/pic"
)

type Image struct {
	width  int
	height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	img_func := func(x, y int) uint8 {
		return uint8(x * y)
	}

	v := img_func(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main_images() {
	// m := Image{256, 64}
	// pic.ShowImage(m)
}
