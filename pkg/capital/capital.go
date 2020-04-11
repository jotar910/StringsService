package capital

import (
	"io"
	"strings"
)

// ParseAll func parses all the letter to uppercase.
func ParseAll(s string) (string, error) {
	return strings.ToUpper(s), nil
}

// ParseEachFirst func parses each first letter on a word to uppercase.
func ParseEachFirst(s string) (string, error) {
	return strings.Title(s), nil
}

// ParseFirst func parses only the first first letter to uppercase.
func ParseFirst(s string) (string, error) {
	if len(s) < 1 {
		return s, nil
	}

	var sb strings.Builder
	sb.WriteString(strings.ToUpper(s[:1]))
	sb.WriteString(s[1:])
	return sb.String(), nil
}

func readBytes(r io.Reader) ([]byte, error) {
	var strBytes []byte
	_, err := r.Read(strBytes)
	if err != nil {
		return nil, err
	}
	return strBytes, err
}
