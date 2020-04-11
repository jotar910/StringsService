package enters

import (
	"regexp"
	"strings"
)

const endlineStr = `.*[.!?:…]$`

func RemoveEnters(s string) (string, error) {
	if len(s) < 1 {
		return s, nil
	}

	endline := regexp.MustCompile(endlineStr)

	var res []string
	for _, p := range strings.Split(s, "\n") {
		if len(p) > 0 {
			if endline.MatchString(p) {
				p = p + "\n"
			} else {
				p = p + " "
			}
			res = append(res, p)
		}
	}
	return strings.Join(res, ""), nil
}
