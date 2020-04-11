package spaces

import (
	"regexp"
	"strings"
)

// RemoveExtra func trims the string and removes the extra spaces inside the string
func RemoveExtra(s string) (string, error) {
	res, err := RemoveExtraMiddle(s)
	if err != nil {
		return "", nil
	}
	return strings.TrimSpace(res), nil
}

// RemoveExtraMiddle func removes the extra spaces inside the string
func RemoveExtraMiddle(s string) (string, error) {
	space := regexp.MustCompile(`[ ]+`)
	return space.ReplaceAllString(s, " "), nil
}
