package main

// https://leetcode-cn.com/problems/qiu-12n-lcof/

func sumNums(n int) int {
	ans := 0
	var sum func(n int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}

func main() {

}
