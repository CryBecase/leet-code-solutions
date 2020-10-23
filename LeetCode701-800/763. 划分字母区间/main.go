package main

// https://leetcode-cn.com/problems/partition-labels/

func partitionLabels(S string) []int {
	counter := make(map[byte]int)
	for i := 0; i < len(S); i++ {
		counter[S[i]]++
	}

	res := make([]int, 0)

	curWindow := make(map[byte]int)
	preLast, curLast := -1, 0
	for curLast < len(S) {
		for curLast = preLast + 1; curLast < len(S); curLast++ {
			curWindow[S[curLast]]++
			valid := true
			for c, count := range curWindow {
				if !valid {
					break
				}
				if counter[c] > count {
					valid = false
				}
			}
			if valid {
				res = append(res, curLast-preLast)
				preLast = curLast
				curWindow = make(map[byte]int)
				break
			}
		}
	}

	return res
}

func main() {

}
