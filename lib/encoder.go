package local

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
)

type Encoder struct {
}

func (encoder *Encoder) Process(img image.Image, dimensions Dimensions, threshold uint) []uint8 {
	bounds := Dimensions{img.Bounds().Max.X, img.Bounds().Max.Y}
	if bounds.Width < dimensions.Width || bounds.Height < dimensions.Height {
		log.Fatal("Result image must be larger than input one!")
	}

	coords := make([]Point, dimensions.Width*dimensions.Height)
	var index int
	for i := 0; i < dimensions.Height; i++ {
		y := (float32(i) + .5) * (float32(bounds.Height) / float32(dimensions.Height))
		for j := 0; j < dimensions.Width; j++ {
			x := (float32(j) + .5) * (float32(bounds.Width) / float32(dimensions.Width))
			coords[index] = Point{int(x), int(y)}
			index++
		}
	}

	pixels := make([]uint8, len(coords))
	for index, value := range coords {
		r, g, b, _ := img.At(value.X, value.Y).RGBA()
		grayscale := uint8((299*r + 587*g + 114*b + 500) / 1000)
		if grayscale >= uint8(threshold) {
			grayscale = 0xff
		} else {
			grayscale = 0x00
		}
		pixels[index] = grayscale
	}
	return pixels
}
