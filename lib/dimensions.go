package lib

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Dimensions struct {
	Width, Height int
}

func (dimensions *Dimensions) Set(value string) error {
	matched, err := regexp.MatchString(`^\d{1,4}:\d{1,4}$`, value)
	if err != nil || !matched {
		return errors.New("Error while dimensions set: incorrect value.")
	}

	args := strings.SplitN(value, ":", 2)
	dimensions.Width, _ = strconv.Atoi(args[0])
	dimensions.Height, _ = strconv.Atoi(args[1])
	return nil
}

func (dimensions *Dimensions) String() string {
	return fmt.Sprintf("Dimensions {Width:%d, Height:%d}", dimensions.Width, dimensions.Height)
}
