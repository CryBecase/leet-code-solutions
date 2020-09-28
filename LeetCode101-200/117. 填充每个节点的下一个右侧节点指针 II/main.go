package main

// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	head := root
	for head != nil {
		var nextLayerStart, node1 *Node
		findNotNilNode := func(cur *Node) {
			if cur == nil {
				return
			}
			if node1 != nil {
				node1.Next = cur
			}
			if nextLayerStart == nil {
				nextLayerStart = cur
			}
			node1 = cur
		}
		for node := head; node != nil; node = node.Next {
			findNotNilNode(node.Left)
			findNotNilNode(node.Right)
		}
		head = nextLayerStart
	}

	return root
}

func main() {

}
