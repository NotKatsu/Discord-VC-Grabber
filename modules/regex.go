package modules

import (
	"regexp"
)

func intCheck(number string) bool {
	return regexp.MustCompile(`^[0-9]+$`).MatchString(number)
}
