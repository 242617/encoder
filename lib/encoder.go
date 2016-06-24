package local

import (
	// "fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
)

type Encoder struct {
	// TODO: Провести всё к uint
	// TODO: Убрать структуру и оставить одну функцию
}

func (encoder *Encoder) Process(img image.Image, dimensions Dimensions, threshold uint) ([]uint8, []uint8) {
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

	pix := make([]byte, len(coords))
	for i, value := range coords {
		r, g, b, _ := img.At(value.X, value.Y).RGBA()
		grayscale := (299*r + 587*g + 114*b + 500) / 1000
		// pix[i] = grayscale
		if byte(grayscale) >= byte(threshold) {
			pix[i] = 1
		}
	}

	length := 8
	res := make([]byte, len(pix)/length)
	for i, value := range pix {
		// x := i % dimensions.Width
		y := i / dimensions.Width
		r := y / length
		n := y % length
		c := i%dimensions.Width + dimensions.Width*r
		res[c] = byte(res[c]) | byte(value<<byte(n))
	}

	return res, pix
}
