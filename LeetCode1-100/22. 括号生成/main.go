package main

// https://leetcode-cn.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	var result []string

	if n == 0 {
		return result
	}
	dfs("", n, n, &result)
	return result
}

func dfs(str string, left, right int, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, str)
		return
	}

	// 剪枝
	if left > right {
		return
	}

	if left > 0 {
		dfs(str+"(", left-1, right, result)
	}

	if right > 0 {
		dfs(str+")", left, right-1, result)
	}
}

func main() {

}
