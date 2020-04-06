package main

// https://leetcode-cn.com/problems/permutations/

var result [][]int

func permute(nums []int) [][]int {
	result = make([][]int, 0, 2*len(nums))

	if len(nums) == 1 {
		result = append(result, nums)
		return result
	}
	visited := make([]int, len(nums), len(nums))

	backtrack(nums, make([]int, 0), visited)

	return result
}

func backtrack(nums []int, tmp, visited []int) {
	if len(tmp) == len(nums) {
		r := make([]int, len(nums))
		copy(r, tmp)
		result = append(result, r)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == 1 {
			continue
		}
		visited[i] = 1
		tmp = append(tmp, nums[i])
		backtrack(nums, tmp, visited)
		visited[i] = 0
		tmp = tmp[:len(tmp)-1]
	}
}

func main() {

}
