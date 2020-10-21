package main

import "fmt"

// https://leetcode-cn.com/problems/long-pressed-name/

func isLongPressedName(name string, typed string) bool {
	p1, p2 := 0, 0
	for p1 < len(name) && p2 < len(typed) {
		if name[p1] == typed[p2] {
			p1++
			p2++
		} else {
			if p1-1 >= 0 && name[p1-1] == typed[p2] {
				p2++
			} else {
				return false
			}
		}
	}

	if p1 == len(name) {
		for p2 < len(typed) {
			if typed[p2] != name[p1-1] {
				return false
			}
			p2++
		}
		return true
	}

	return false
}

func main() {
	fmt.Println(isLongPressedName("alex", "alexxr"))
}
