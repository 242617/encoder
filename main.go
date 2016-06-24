package main

import (
	"./lib"
	"flag"
	"fmt"
	"image"
	"image/color"
	jpeg "image/jpeg"
	"log"
	"os"
)

// TODO: Убрать -dimensions - всегда 64x48

const (
	HELP = `Image Encoder v0.1
Returns array of image pixels
Usage:
	-dimensions   string   Target dimensions (e.g. 64:48)
	-input        string   Input file (*.png, *.gif, *.jpg) (e.g. "input.gif")
	-threshold    uint     Gray color threshold
	-preview      bool     Preview result, saves "output.jpg" if true
	-help         bool     Help page
	`
)

func main() {
	config := local.Config{}
	config.Width, config.Height = 64, 48
	flag.Var(&config.Dimensions, "dimensions", "Target dimensions")
	flag.StringVar(&config.Input, "input", "", "Input file")
	flag.UintVar(&config.Threshold, "threshold", 160, "Gray color threshold")
	flag.BoolVar(&config.Preview, "preview", false, "Preview result")
	flag.BoolVar(&config.Help, "help", false, "Help page")
	flag.Parse()
	fmt.Println(config)

	if config.Help {
		fmt.Printf(fmt.Sprint(HELP))
		return
	}

	encoder := local.Encoder{}
	img := load(config.Input)
	result, pixels := encoder.Process(img, config.Dimensions, config.Threshold)
	fmt.Println(result)
	fmt.Println(pixels)

	if config.Preview {
		target := image.NewGray(image.Rect(0, 0, config.Width, config.Height))
		for index, value := range pixels {
			x := index % config.Width
			y := int(float64(index) / float64(config.Width))
			target.Set(x, y, color.Gray{value})
		}
		save(target, "preview.jpg")
	}
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
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		log.Fatal(err)
	}
}
