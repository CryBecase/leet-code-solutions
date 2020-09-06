package main

import (
	"fmt"
	"strings"
)

var m = map[byte]int{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
	'i': 8,
	'j': 9,
	'k': 10,
	'l': 11,
	'm': 12,
	'n': 13,
	'o': 14,
	'p': 15,
	'q': 16,
	'r': 17,
	's': 18,
	't': 19,
	'u': 20,
	'v': 21,
	'w': 22,
	'x': 23,
	'y': 24,
	'z': 25,
}

var chs = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

func modifyString(s string) string {
	if s == "" {
		return ""
	}
	if len(s) == 1 && s[0] == '?' {
		return "a"
	}

	builder := strings.Builder{}

	sslice := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] != '?' {
			sslice[i] = s[i]
			builder.WriteByte(s[i])
			continue
		}

		if i-1 >= 0 && i+1 < len(s) {
			for _, b := range chs {
				if b != sslice[i-1] && b != s[i+1] {
					sslice[i] = b
					builder.WriteByte(b)
					break
				}
			}
		} else if i-1 >= 0 {
			for _, b := range chs {
				if b != sslice[i-1] {
					sslice[i] = b
					builder.WriteByte(b)
					break
				}
			}
		} else if i+1 < len(s) {
			for _, b := range chs {
				if b != s[i+1] {
					sslice[i] = b
					builder.WriteByte(b)
					break
				}
			}
		}
	}

	return builder.String()
}

func main() {
	fmt.Println(modifyString("??yw?ipkj?"))
}
