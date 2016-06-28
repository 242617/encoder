package encoder

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	png "image/png"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	HELP = `Image Encoder v0.2
Decodes input_file (*.png, *.gif, *.jpg) and returns uint array for 64x48 OLED display
Usage:
	encoder [options] input_file
Options:
	-threshold    uint     Gray color threshold
	-preview      bool     Preview result, saves "output.jpg" if true
	-help         bool     Help page
`
	LENGTH int = 8
	WIDTH  int = 64
	HEIGHT int = 48
)

func main() {
	config := Config{}
	if len(os.Args) < 2 {
		fmt.Printf(fmt.Sprint(HELP))
		return
	}
	config.Input = os.Args[len(os.Args)-1 : len(os.Args)][0]
	flag.UintVar(&config.Threshold, "threshold", 160, "Gray color threshold")
	flag.BoolVar(&config.Preview, "preview", false, "Preview result")
	flag.BoolVar(&config.Help, "help", false, "Help page")
	flag.Parse()

	if config.Help {
		fmt.Printf(fmt.Sprint(HELP))
		return
	}

	result, pixels := encode(load(config.Input), config.Threshold)
	fmt.Println(Format(result))

	if config.Preview {
		target := image.NewGray(image.Rect(0, 0, WIDTH, HEIGHT))
		for index, value := range pixels {
			x := index % WIDTH
			y := int(float64(index) / float64(WIDTH))
			target.Set(x, y, color.Gray{value})
		}
		save(target, "preview.png")
	}
}

type Config struct {
	Input     string
	Threshold uint
	Preview   bool
	Help      bool
}

type Point struct {
	X, Y int
}

func encode(img image.Image, threshold uint) ([]byte, []byte) {
	if img.Bounds().Max.X < WIDTH || img.Bounds().Max.Y < HEIGHT {
		log.Fatal("Result image must be larger than input one!")
	}

	coords := make([]Point, WIDTH*HEIGHT)
	for index, i := 0, 0; i < HEIGHT; i++ {
		y := (float32(i) + .5) * (float32(img.Bounds().Max.Y) / float32(HEIGHT))
		for j := 0; j < WIDTH; j++ {
			x := (float32(j) + .5) * (float32(img.Bounds().Max.X) / float32(WIDTH))
			coords[index] = Point{int(x), int(y)}
			index++
		}
	}

	result := make([]byte, len(coords)/LENGTH)
	preview := make([]byte, len(coords))
	for i, value := range coords {
		r, g, b, _ := img.At(value.X, value.Y).RGBA()
		pixel, grayscale := 0, (2126*r+7152*g+722*b+500)/10000
		if byte(grayscale) >= byte(threshold) {
			pixel = 1
			preview[i] = 0xff
		}
		y := i / WIDTH
		m := y / LENGTH
		n := y % LENGTH
		c := i%WIDTH + WIDTH*m
		result[c] = byte(result[c]) | byte(pixel<<byte(n))
	}

	return result, preview
}

func load(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func save(img image.Image, path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
}

func Format(src []byte) string {
	res := make([]string, len(src))
	for i, v := range src {
		res[i] = strconv.Itoa(int(v))
	}
	return strings.Join(res, ",")
}
