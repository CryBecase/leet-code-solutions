package main

import "fmt"

// https://leetcode-cn.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	path := make([]int, 0)
	sum := 0
	res := make([][]int, 0)

	dfs := func(int) {}
	dfs = func(idx int) {
		if idx < 0 || idx >= len(candidates) || sum > target {
			return
		}

		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		num := candidates[idx]

		sum += num
		path = append(path, num)
		dfs(idx)
		sum -= num
		path = path[:len(path)-1]

		for ; idx+1 < len(candidates); idx++ {
			num = candidates[idx+1]
			sum += num
			path = append(path, num)
			dfs(idx + 1)
			sum -= num
			path = path[:len(path)-1]
		}

	}
	for i, num := range candidates {
		sum = num
		path = []int{num}
		dfs(i)
	}

	return res
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
