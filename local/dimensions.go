package local

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Dimensions struct {
	width, height int
}

func (dimensions *Dimensions) Set(value string) error {
	matched, err := regexp.MatchString(`^\d{1,4}:\d{1,4}$`, value)
	if err != nil || !matched {
		return errors.New("Error while dimensions set: incorrect value.")
	}

	args := strings.SplitN(value, ":", 2)
	dimensions.width, _ = strconv.Atoi(args[0])
	dimensions.height, _ = strconv.Atoi(args[1])
	return nil
}

func (dimensions *Dimensions) String() string {
	return fmt.Sprintf("Dimensions {width:%d, height:%d}", dimensions.width, dimensions.height)
}
