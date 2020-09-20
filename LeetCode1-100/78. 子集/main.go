package main

// https://leetcode-cn.com/problems/subsets/

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	dfs := func(int) {}
	dfs = func(i int) {
		if i == len(nums) {
			res = append(res, append([]int(nil), path...))
			return
		}

		path = append(path, nums[i])
		dfs(i + 1) // 选当前这个数
		path = path[:len(path)-1]
		dfs(i + 1) // 不选当前这个数
	}
	dfs(0)
	return res
}

func main() {

}
