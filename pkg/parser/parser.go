package parser

import (
	"fmt"
	"strings"

	"joaor.dev.com/joao/stringservice/pkg/capital"
	"joaor.dev.com/joao/stringservice/pkg/enters"
	"joaor.dev.com/joao/stringservice/pkg/minuscule"
	"joaor.dev.com/joao/stringservice/pkg/spaces"
)

// Parse func parses a string accordingly to the passed flags
func Parse(t string, flagsStr string) string {
	if len(flagsStr) == 0 {
		return t
	}

	fs := strings.Split(flagsStr, ",")

	var sb strings.Builder
	sb.Grow(len(t))

	fs = sortFs(fs)
	s := t

	for _, f := range fs {
		var err error

		switch {
		case f == "c":
			s, err = capital.ParseFirst(s)
		case f == "ac":
			s, err = capital.ParseAll(s)
		case f == "acf":
			s, err = capital.ParseEachFirst(s)
		case f == "s":
			s, err = spaces.RemoveExtraMiddle(s)
		case f == "as":
			s, err = spaces.RemoveExtra(s)
		case f == "e":
			s, err = enters.RemoveEnters(s)
		case f == "l":
			s, err = minuscule.ParseFirst(s)
		case f == "al":
			s, err = minuscule.ParseAll(s)
		case f == "afl":
			s, err = minuscule.ParseEachFirst(s)
		default:
			err = fmt.Errorf("Unknown command %q", f)
		}
		if err == nil {
			sb.Reset()
			sb.WriteString(s)
		} else {
			fmt.Println(err)
		}
	}

	return sb.String()
}

func same(a, b string) bool {
	return true
}

func sortFs(fs []string) []string {
	var fsSorted []string
	for _, f := range fs {
		fsSorted = append(fsSorted, string(sortM([]rune(f), 0, len(f)-1)))
	}
	return fsSorted
}

func sortM(runes []rune, start, end int) []rune {
	ret := []rune{}
	if start >= end {
		return append(ret, runes[start])
	}

	middle := (end + start + 1) / 2
	chn := make(chan []rune, 2)

	go func(runes []rune, start, end int) {
		chn <- sortM(runes, start, end)
	}(runes, start, middle-1)
	go func(runes []rune, start, end int) {
		chn <- sortM(runes, start, end)
	}(runes, middle, end)

	first := <-chn
	last := <-chn

	return mergeM(first, last)
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func mergeM(arrA, arrB []rune) []rune {
	var i, iA, iB int
	lenA := len(arrA)
	lenB := len(arrB)
	ret := make([]rune, lenA+lenB)

	for iA < lenA && iB < lenB {
		var val rune
		if arrA[iA] < arrB[iB] {
			val = arrA[iA]
			iA++
		} else {
			val = arrB[iB]
			iB++
		}
		ret[i] = val
		i++
	}

	for iA < lenA {
		ret[i] = arrA[iA]
		iA++
		i++
	}

	for iB < lenB {
		ret[i] = arrB[iB]
		iB++
		i++
	}

	return ret
}
