package local

import (
	"flag"
	"fmt"
)

const (
	BW = iota
	Gray
	RGB
)

type Config struct {
	Dimensions
	Mode
	help    bool
	version bool
	input   string
	output  string
}

func (config *Config) Init() {
	fmt.Println("Init")

	config.Mode = "bw"
	config.Dimensions.width, config.Dimensions.height = 64, 48
	flag.Var(&config.Dimensions, "d", "Output image dimensions (e.g. 64:48)")
	flag.Var(&config.Dimensions, "dimensions", "Output image dimensions (e.g. 64:48)")
	flag.Var(&config.Mode, "m", "Output mode: [bw|gs|rgb] (bw default)")
	flag.Var(&config.Mode, "mode", "Output mode: [bw|gs|rgb] (bw default)")
	flag.BoolVar(&config.help, "h", false, "Help with the application")
	flag.BoolVar(&config.help, "help", false, "Help with the application")
	flag.StringVar(&config.input, "i", "input.jpg", "Input image")
	flag.StringVar(&config.input, "input", "input.jpg", "Input image")
	flag.StringVar(&config.output, "o", "", "Output image")
	flag.StringVar(&config.output, "output", "", "Output image")
	flag.Parse()
}

func (config Config) Help() bool {
	return config.help
}

func (config Config) Input() string {
	return config.input
}

func (config Config) String() string {
	return fmt.Sprintf("Config {help: %t, mode: %s, dimensions: %v}", config.help, config.Mode, config.Dimensions)
}
