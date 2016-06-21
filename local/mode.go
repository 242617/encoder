package local

import (
	"errors"
	"fmt"
	"regexp"
)

type Mode string

func (mode *Mode) Set(value string) error {
	matched, err := regexp.MatchString(`^(bw|gs|rgb)$`, value)
	if err != nil || !matched {
		return errors.New("Error while dimensions set: incorrect value.")
	}

	*mode = Mode(value)
	return nil
}

func (mode *Mode) String() string {
	return fmt.Sprintf("Mode {value:%s}", "mode")
}
