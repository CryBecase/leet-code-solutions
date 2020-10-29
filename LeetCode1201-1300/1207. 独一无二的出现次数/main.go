package main

// https://leetcode-cn.com/problems/unique-number-of-occurrences/

func uniqueOccurrences(arr []int) bool {
	counter := make(map[int]int)
	for _, n := range arr {
		counter[n]++
	}
	counterCounter := make(map[int]int)
	for _, count := range counter {
		counterCounter[count]++
	}
	for _, v := range counterCounter {
		if v > 1 {
			return false
		}
	}
	return true
}

func main() {

}
