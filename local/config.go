package local

import (
	"fmt"
)

type Config struct {
	Dimensions
	help    bool
	version bool
	input   string
	output  string
}

func (config *Config) Init(args []string) {
	for _, value := range args {
		switch value {
		case "-v":
			config.help = false
			config.version = true
		case "--version":
			config.help = false
			config.version = true
		case "-h":
			config.help = true
			config.version = false
		case "--help":
			config.help = true
			config.version = false
		}
	}
	config.Dimensions.width = 64
	config.Dimensions.height = 48
	config.input = "mint.jpg"
}

func (config Config) Echo() {
	fmt.Println("Echo")
	fmt.Println("dimensions:", config.Dimensions)
	fmt.Println("input:", config.input)
	fmt.Println("output:", config.output)
}

func (config Config) Help() bool {
	return config.help
}

func (config Config) Version() bool {
	return config.version
}

func (config Config) Input() string {
	return config.input
}
