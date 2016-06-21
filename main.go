package main

import (
	"./local"
	"fmt"
	"image"
	"image/color"
	png "image/png"
	"log"
	"os"
)

const (
	HELP = `Image Encoder v0.1
Returns array of image pixels
Usage:
	-m, --mode         string   Output mode: "bw" (default), "gs"
	-d, --dimensions   string   Dimensions of result array (e.g. 64:48)
	-i, --input        string   Input file (*.png, *.gif, *.jpg) (e.g. "input.gif")
	-p, --preview      bool     Preview result, saves "output.png" if true
	-h, --help         bool     Help page
	`
)

func main() {
	config := local.Config{}
	config.Init()
	fmt.Println(config)

	if config.Help() {
		fmt.Printf(fmt.Sprint(HELP))
		return
	}

	encoder := local.Encoder{}
	img := load(config.Input())
	pixels := encoder.Process(img, config.Mode, config.Dimensions)
	fmt.Println(pixels)

	if config.Preview() {
		target := image.NewGray(image.Rect(0, 0, config.Dimensions.Width(), config.Dimensions.Height()))
		for index, value := range pixels {
			x := index % config.Dimensions.Width()
			y := int(float64(index) / float64(config.Dimensions.Width()))
			target.Set(x, y, color.Gray{value})
		}
		save(target, "preview.png")
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
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
}
