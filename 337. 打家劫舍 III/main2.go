package main

func rob2(root *TreeNode) int {
	r := robInternal2(root)
	return maxFunc(r[0], r[1])
}

// [0]不偷
// [1]偷
func robInternal2(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{0, 0}
	}

	result := [2]int{}

	left := robInternal2(root.Left)
	right := robInternal2(root.Right)

	result[0] = maxFunc(left[0], left[1]) + maxFunc(right[0], right[1])
	result[1] = root.Val + left[0] + right[0]
	return result
}

func main() {

}
