package main

import (
	"fmt"
	"strings"
)

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
