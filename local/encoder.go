package local

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type Encoder struct {
}

func (encoder *Encoder) Process(img image.Image, mode Mode, dimensions Dimensions) []uint8 {
	fmt.Println("Process")

	bounds := Dimensions{img.Bounds().Max.X, img.Bounds().Max.Y}
	// if bounds.width < dimensions.width || bound.height < dimensions.height {
	// TODO: Проверка размера исходного изображения - должно быть больше
	// }

	coords := make([]Point, dimensions.width*dimensions.height)
	var index int
	for i := 0; i < dimensions.height; i++ {
		y := (float32(i) + .5) * (float32(bounds.height) / float32(dimensions.height))
		for j := 0; j < dimensions.width; j++ {
			x := (float32(j) + .5) * (float32(bounds.width) / float32(dimensions.width))
			coords[index] = Point{int(x), int(y)}
			index++
		}
	}

	pixels := make([]uint8, len(coords))
	for index, value := range coords {
		r, g, b, _ := img.At(value.X, value.Y).RGBA()
		grayscale := uint8((299*r + 587*g + 114*b + 500) / 1000)
		if mode == "bw" {
			if grayscale >= 0xb0 {
				grayscale = 0xff
			} else {
				grayscale = 0x00
			}
		}
		pixels[index] = grayscale
	}
	return pixels
}
