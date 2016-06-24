package lib

import (
	"fmt"
)

type Config struct {
	Dimensions
	Input     string
	Threshold uint
	Preview   bool
	Help      bool
}

func (config Config) String() string {
	return fmt.Sprintf("Config {Dimensions: %v, Input: %s, Threshold: %d, Preview: %t, Help: %t}", config.Dimensions, config.Input, config.Threshold, config.Preview, config.Help)
}
