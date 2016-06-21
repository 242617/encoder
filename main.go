package main

import (
	"./local"
	"fmt"
	"image"
	"log"
	"os"
)

const (
	HELP = `Image Encoder v0.1
Returns array of image pixels
Usage:
	-m, --mode          Output mode: bw (default), gs, rgb)
	-d, --dimensions    Dimensions of result array
	-i, --input         Input file (*.png, *.jpg)
	-o, --output        Output file. If ommitted no files saved
	-h, --help          Help page
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
	result := encoder.Process(img, config.Dimensions)
	fmt.Println(result)
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
