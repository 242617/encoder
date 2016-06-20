package local

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	png "image/png"
	"log"
	"os"
)

type Manager struct {
}

func (manager *Manager) Process(path string, dimensions Dimensions) []uint {
	fmt.Println("Process")
	img := load(path)
	bounds := Dimensions{img.Bounds().Max.X, img.Bounds().Max.Y}
	// if bounds.width < dimensions.width || bound.height < dimensions.height {
	// TODO: Проверка размера исходного изображения - должно быть больше
	// }
	// fmt.Println(bounds)
	// fmt.Println(dimensions)

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

	// result := image.NewRGBA(image.Rect(0, 0, dimensions.width, dimensions.height))
	new := image.NewGray(image.Rect(0, 0, dimensions.width, dimensions.height))
	for index, value := range coords {
		x := index % dimensions.width
		y := int(float64(index) / float64(dimensions.width))
		r, g, b, _ := img.At(value.X, value.Y).RGBA()
		// fmt.Printf("(%d, %d, %d) byte\n", byte(r), byte(g), byte(b))
		// result.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), 255})

		grayscale := uint8((299*r + 587*g + 114*b + 500) / 1000)
		/*if grayscale >= 0xa0 {
			grayscale = 0xff
		} else {
			grayscale = 0x00
		}*/
		new.Set(x, y, color.Gray{uint8(grayscale)})
	}

	save(new)

	return []uint{0xff, 0xff, 0xff}
}

func load(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// info, err := file.Stat()
	// fmt.Println("File stat:", info.Size(), err)

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func save(img image.Image) {

	file, err := os.OpenFile("_result.png", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}

}
