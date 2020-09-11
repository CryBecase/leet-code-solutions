package main

// https://leetcode-cn.com/problems/combination-sum-iii/

func combinationSum3(k int, n int) [][]int {
	if k == 0 {
		return nil
	}

	path := make([]int, 0)
	res := make([][]int, 0)
	dfs := func(int, int) {}
	dfs = func(target, idx int) {
		if target == 0 && len(path) == k {
			res = append(res, append([]int(nil), path...))
			return
		}

		for i := idx; i <= 9; i++ {
			if target-i < 0 {
				break
			}
			path = append(path, i)
			dfs(target-i, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(n, 1)
	return res
}

func main() {

}
