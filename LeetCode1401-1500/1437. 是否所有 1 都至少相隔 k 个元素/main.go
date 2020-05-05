package main

// https://leetcode-cn.com/problems/check-if-all-1s-are-at-least-length-k-places-away/

func kLengthApart(nums []int, k int) bool {
	if k == 0 {
		return true
	}

	count := 0
	start := false
	for _, num := range nums {
		if start {
			if num == 1 {
				if count < k {
					return false
				}
				count = 0
			} else {
				count++
			}
		} else {
			if num == 1 {
				start = true
			}
		}
	}
	return true
}

func main() {

}
