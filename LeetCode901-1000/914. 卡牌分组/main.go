package main

import "sort"

// https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/

func hasGroupsSizeX(deck []int) bool {
	if len(deck) == 2 {
		if deck[0] == deck[1] {
			return true
		}
	}
	sort.Ints(deck)
X:
	for size := 2; size <= len(deck)/2; size++ {
		if len(deck)%size == 0 {
			for i := 0; i < len(deck)/size; i++ {
				for j := 0; j < size; j++ {
					if deck[size*i+j] != deck[size*i] {
						continue X
					}
				}
			}
			return true
		}
	}
	return false
}

func main() {

}
