package main

// https://leetcode-cn.com/problems/daily-temperatures/

func dailyTemperatures(T []int) []int {
	ans := make([]int, len(T))
	stack := make([]int, 0)
	for i, t := range T {
		for len(stack) > 0 && T[stack[len(stack)-1]] < t {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {

}
