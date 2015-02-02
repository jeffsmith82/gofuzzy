package gofuzzy

import (
	"errors"
	"unicode"
)

var code = []byte("01230127022455012623017202")

func Ld(s, t string) int {
	d := make([][]int, len(s)+1)
	for i := range d {
		d[i] = make([]int, len(t)+1)
	}
	for i := range d {
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}
	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				min := d[i-1][j]
				if d[i][j-1] < min {
					min = d[i][j-1]
				}
				if d[i-1][j-1] < min {
					min = d[i-1][j-1]
				}
				d[i][j] = min + 1
			}
		}

	}
	return d[len(s)][len(t)]
}

func Soundex(s string) (string, error) {
	var sx [4]byte
	var sxi int
	var cx, lastCode byte
	for i, c := range s {
		switch {
		case !unicode.IsLetter(c):
			if c < ' ' || c == 127 {
				return "", errors.New("ASCII control characters disallowed")
			}
			if i == 0 {
				return "", errors.New("initial character must be a letter")
			}
			lastCode = '0'
			continue
		case c >= 'A' && c <= 'Z':
			cx = byte(c - 'A')
		case c >= 'a' && c <= 'z':
			cx = byte(c - 'a')
		default:
			return "", errors.New("non-ASCII letters unsupported")
		}
		// cx is valid letter index at this point
		if i == 0 {
			sx[0] = cx + 'A'
			sxi = 1
			continue
		}
		switch x := code[cx]; x {
		case '7', lastCode:
		case '0':
			lastCode = '0'
		default:
			sx[sxi] = x
			if sxi == 3 {
				return string(sx[:]), nil
			}
			sxi++
			lastCode = x
		}
	}
	if sxi == 0 {
		return "", errors.New("no letters present")
	}
	for ; sxi < 4; sxi++ {
		sx[sxi] = '0'
	}
	return string(sx[:]), nil
}
