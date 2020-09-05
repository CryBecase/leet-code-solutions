package main

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/permutation-sequence/

// 0 ~ 9 的阶乘
var factorials = []int{
	1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880,
}

func getPermutation(n int, k int) string {
	path := strings.Builder{}
	used := make(map[int]bool)

	dfs := func() {}
	dfs = func() {
		if path.Len() == n {
			return
		}

		// 未来要进入 path 的数字的全排列的个数
		cnt := factorials[n-path.Len()-1]
		for i := 1; i <= n; i++ {
			if used[i] {
				continue
			}

			if cnt < k {
				k -= cnt
				continue
			}

			used[i] = true
			path.WriteString(strconv.Itoa(i))
			dfs()
			return
		}
	}
	dfs()

	return path.String()
}

func main() {

}
