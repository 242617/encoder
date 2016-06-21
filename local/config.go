package local

import (
	"flag"
	"fmt"
)

type Config struct {
	Dimensions
	Mode
	help    bool
	input   string
	preview bool
}

func (config *Config) Init() {
	config.Mode = "bw"
	config.Dimensions.width, config.Dimensions.height = 64, 48
	flag.Var(&config.Dimensions, "d", "Dimensions of result array")
	flag.Var(&config.Dimensions, "dimensions", "Dimensions of result array")
	flag.Var(&config.Mode, "m", "Output mode: [bw|gs] (bw default)")
	flag.Var(&config.Mode, "mode", "Output mode: [bw|gs] (bw default)")
	flag.BoolVar(&config.help, "h", false, "Help page")
	flag.BoolVar(&config.help, "help", false, "Help page")
	flag.StringVar(&config.input, "i", "", "Input file")
	flag.StringVar(&config.input, "input", "", "Input file")
	flag.BoolVar(&config.preview, "p", false, "Preview result")
	flag.BoolVar(&config.preview, "preview", false, "Preview result")
	flag.Parse()
}

func (config Config) Help() bool {
	return config.help
}

func (config Config) Input() string {
	return config.input
}

func (config Config) Preview() bool {
	return config.preview
}

func (config Config) String() string {
	return fmt.Sprintf("Config {help: %t, mode: %s, dimensions: %v, input: %s, preview: %t}", config.help, config.Mode, config.Dimensions, config.input, config.preview)
}
