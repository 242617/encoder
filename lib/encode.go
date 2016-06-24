package lib

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
)

const LENGTH int = 8

func Encode(img image.Image, dimensions Dimensions, threshold uint) ([]uint8, []uint8) {
	bounds := Dimensions{img.Bounds().Max.X, img.Bounds().Max.Y}
	if bounds.Width < dimensions.Width || bounds.Height < dimensions.Height {
		log.Fatal("Result image must be larger than input one!")
	}

	coords := make([]Point, dimensions.Width*dimensions.Height)
	for index, i := 0, 0; i < dimensions.Height; i++ {
		y := (float32(i) + .5) * (float32(bounds.Height) / float32(dimensions.Height))
		for j := 0; j < dimensions.Width; j++ {
			x := (float32(j) + .5) * (float32(bounds.Width) / float32(dimensions.Width))
			coords[index] = Point{int(x), int(y)}
			index++
		}
	}

	preview := make([]byte, len(coords))
	result := make([]byte, len(coords)/LENGTH)
	for i, value := range coords {
		r, g, b, _ := img.At(value.X, value.Y).RGBA()
		grayscale := (299*r + 587*g + 114*b + 500) / 1000

		var pixel uint
		if byte(grayscale) >= byte(threshold) {
			pixel = 1
			preview[i] = 0xff
		}

		y := i / dimensions.Width
		m := y / LENGTH
		n := y % LENGTH
		c := i%dimensions.Width + dimensions.Width*m
		result[c] = byte(result[c]) | byte(pixel<<byte(n))
	}

	return result, preview
}
