package main

// https://leetcode-cn.com/problems/increasing-subsequences/

func findSubsequences(nums []int) [][]int {
	var res [][]int
	var path []int
	// cur是当前下标，length是路径长度，pre是前一个数
	dfs := func(cur, length, pre int) {}
	dfs = func(cur, length, pre int) {
		if length >= 2 {
			// 保存答案
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		// 保证下一个选出来的数不会重复
		cache := make(map[int]int)

		// 考虑其他数字
		for i := cur; i < len(nums); i++ {
			_, ok := cache[nums[i]]

			if nums[i] < pre || ok {
				continue
			}

			path = append(path, nums[i])
			cache[nums[i]] = 1
			dfs(i+1, length+1, nums[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0, -101)
	return res
}

func main() {

}
