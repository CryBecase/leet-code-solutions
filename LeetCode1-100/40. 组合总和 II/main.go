package main

import "sort"

// https://leetcode-cn.com/problems/combination-sum-ii/

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var path []int

	sort.Ints(candidates)

	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if target == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if target-candidates[i] < 0 {
				break
			}
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			dfs(target-candidates[i], i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(target, 0)

	return res
}

func minFunc(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {

}
