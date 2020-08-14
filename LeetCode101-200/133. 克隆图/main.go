package main

type Node struct {
	Val       int
	Neighbors []*Node
}

var indice map[*Node]*Node

func cloneGraph(node *Node) *Node {
	indice = make(map[*Node]*Node)
	return dfs(node)
}

func dfs(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, ok := indice[node]; ok {
		return n
	}
	result := &Node{node.Val, nil}
	indice[node] = result
	neighbors := make([]*Node, 0)
	for _, n := range node.Neighbors {
		neighbors = append(neighbors, dfs(n))
	}
	result.Neighbors = neighbors
	return result
}
