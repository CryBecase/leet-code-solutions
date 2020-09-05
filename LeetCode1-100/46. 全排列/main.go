package main

// https://leetcode-cn.com/problems/permutations/

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	used := make(map[int]bool)
	path := make([]int, 0)

	dfs := func() {}
	dfs = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		for _, num := range nums {
			if used[num] {
				continue
			}

			used[num] = true
			path = append(path, num)
			dfs()
			path = path[:len(path)-1]
			used[num] = false
		}
	}
	dfs()

	return result
}

func main() {

}
