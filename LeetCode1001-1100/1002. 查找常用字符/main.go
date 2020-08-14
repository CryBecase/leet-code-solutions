package main

import "math"

// https://leetcode-cn.com/problems/find-common-characters/

func commonChars(A []string) []string {
	indice := make(map[rune][][]int)
	for i, str := range A {
		for _, c := range str {
			if _, ok := indice[c]; !ok {
				indice[c] = make([][]int, len(A))
			}
			indice[c][i] = append(indice[c][i], 1)
		}
	}
	result := make([]string, 0)
	for k, v := range indice {
		min := math.MaxInt32
		for _, arr := range v {
			if len(arr) < min {
				min = len(arr)
			}
		}
		for i := min; i > 0; i-- {
			result = append(result, string(k))
		}
	}
	return result
}

func main() {

}
