package main

func pathSum2(root *TreeNode, sum int) int {
	return calc(root, sum, make([]int, 1000, 1000), 0)
}

// arr 保存路经
// p 指向路经终点
func calc(root *TreeNode, sum int, arr []int, p int) int {
	if root == nil {
		return 0
	}
	tmp := root.Val
	n := 0
	if root.Val == sum {
		n++
	}
	for i := p - 1; i >= 0; i-- {
		tmp += arr[i]
		if tmp == sum {
			n++
		}
	}
	arr[p] = root.Val
	return n + calc(root.Left, sum, arr, p+1) + calc(root.Right, sum, arr, p+1)
}

func main() {

}
