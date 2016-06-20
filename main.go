package main

import (
	"./local"
	"fmt"
	"os"
)

const help = `Image Encoder
Returns array of uint8 of image
-d, --dimensions    Dimensions of result array.
-i, --input         Input file (*.png, *.jpg).
-o, --output        Output file. If ommitted returns input stdout.
-h, --help          Help page.
-v, --version       Application version.
`
const version = "Image Encoder v0.1"

func main() {
	config := local.Config{}
	config.Init(os.Args[1:])
	if config.Help() {
		fmt.Println(help)
		return
	}
	if config.Version() {
		fmt.Println(version)
		return
	}

	manager := local.Manager{}
	result := manager.Process(config.Input(), config.Dimensions)
	fmt.Println(result)

}
