package main

// https://leetcode-cn.com/problems/combinations/

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	dfs := func(int) {}
	dfs = func(start int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i <= n; i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(1)
	return res
}

func main() {

}
