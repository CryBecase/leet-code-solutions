package main

import "strings"

func convert(s string, numRows int) string {
	numRows = min(len(s), numRows)

	if numRows == 1 {
		return s
	}

	if numRows == 0 {
		return ""
	}

	builderArr := make([]strings.Builder, numRows)
	curRow, goDown := 0, false
	for _, c := range s {
		builderArr[curRow].WriteRune(c)
		if curRow == 0 || curRow >= numRows-1 {
			goDown = !goDown
		}
		if goDown {
			curRow++
		} else {
			curRow--
		}
	}

	var resultBuilder strings.Builder
	for _, builder := range builderArr {
		resultBuilder.WriteString(builder.String())
	}
	result := resultBuilder.String()

	return result
}

func min(n, m int) int {
	if n <= m {
		return n
	} else {
		return m
	}
}

func main() {

}
