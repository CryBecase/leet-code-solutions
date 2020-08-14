package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/multiply-strings/

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	p2 := len(num2) - 1
	results := make([]string, 0)
	for p2 >= 0 {
		n2 := int(num2[p2] - '0')
		builder := strings.Builder{}
		arr := make([]int, 0)
		carry := 0
		for p1 := len(num1) - 1; p1 >= 0; p1-- {
			n1 := int(num1[p1] - '0')
			sum := n1*n2 + carry
			n := sum % 10
			carry = sum / 10
			arr = append(arr, n)
		}
		if carry > 0 {
			arr = append(arr, carry)
		}
		for i := len(arr) - 1; i >= 0; i-- {
			builder.WriteString(strconv.Itoa(arr[i]))
		}
		for i := len(num2) - 1 - p2; i > 0; i-- {
			builder.WriteString("0")
		}
		p2--
		results = append(results, builder.String())
	}

	carry := 0
	ptrs := make([]int, len(results))
	arr := make([]int, 0)
	for i := range ptrs {
		ptrs[i] = len(results[i]) - 1
	}
	builder := strings.Builder{}
	for {
		has := false
		for _, p := range ptrs {
			if p >= 0 {
				has = true
			}
		}
		if !has {
			break
		}
		sum := carry
		for i, p := range ptrs {
			if p >= 0 {
				n := int(results[i][p] - '0')
				sum += n
			}
		}
		n := sum % 10
		carry = sum / 10
		arr = append(arr, n)
		for i := range ptrs {
			ptrs[i]--
		}
	}
	if carry > 0 {
		arr = append(arr, carry)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		builder.WriteString(strconv.Itoa(arr[i]))

	}
	return builder.String()
}

func main() {
	fmt.Println(multiply("123", "456"))
}
