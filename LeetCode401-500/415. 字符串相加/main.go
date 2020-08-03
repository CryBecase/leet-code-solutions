package main

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/add-strings/

func addStrings(num1 string, num2 string) string {
	numSlice := make([]int, 0)
	p1, p2 := len(num1)-1, len(num2)-1
	carry := 0
	for p1 >= 0 && p2 >= 0 {
		n1, _ := strconv.Atoi(string(num1[p1]))
		n2, _ := strconv.Atoi(string(num2[p2]))
		n := (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10
		numSlice = append(numSlice, n)
		p1--
		p2--
	}
	for p1 >= 0 {
		n1, _ := strconv.Atoi(string(num1[p1]))
		n := (n1 + carry) % 10
		carry = (n1 + carry) / 10
		numSlice = append(numSlice, n)
		p1--
	}
	for p2 >= 0 {
		n2, _ := strconv.Atoi(string(num2[p2]))
		n := (n2 + carry) % 10
		carry = (n2 + carry) / 10
		numSlice = append(numSlice, n)
		p2--
	}
	if carry > 0 {
		numSlice = append(numSlice, carry)
		carry = 0
	}
	builder := strings.Builder{}
	for i := len(numSlice) - 1; i >= 0; i-- {
		builder.WriteString(strconv.Itoa(numSlice[i]))
	}
	return builder.String()
}

func main() {

}
