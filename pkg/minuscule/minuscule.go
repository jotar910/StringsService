package minuscule

import (
	"fmt"
	"strings"
)

// ParseAll func parses all the letters to lower case.
func ParseAll(s string) (string, error) {
	return strings.ToLower(s), nil
}

// ParseEachFirst func parses each first letter on a word to lower case.
func ParseEachFirst(s string) (string, error) {
	if len(s) < 1 {
		return s, nil
	}

	var ws []string
	for _, w := range strings.Split(s, " ") {
		if len(w) > 0 {
			completed := fmt.Sprintf("%s%s", strings.ToLower(w[:1]), w[1:])
			ws = append(ws, completed)
		} else {
			ws = append(ws, "")
		}
	}
	return strings.Join(ws, " "), nil
}

// ParseFirst func parses the first letter to lower case.
func ParseFirst(s string) (string, error) {
	if len(s) < 1 {
		return s, nil
	}
	return fmt.Sprintf("%s%s", strings.ToLower(s[:1]), s[1:]), nil
}
