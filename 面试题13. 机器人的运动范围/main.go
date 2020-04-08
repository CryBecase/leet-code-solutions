package main

import (
	"container/list"
	"fmt"
)

// https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/

//  0   ~   n-1
//  ~
// m-1

type position struct {
	x, y int
}

func movingCount(m int, n int, k int) int {
	if k < 0 {
		return 0
	}
	result := 1

	visited := make([][]int, m, m)
	for i := range visited {
		visited[i] = make([]int, n, n)
	}

	queue := list.New()
	queue.PushBack(position{
		x: 0,
		y: 0,
	})
	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		p := e.Value.(position)
		if p.x+1 < m && p.y < n && visited[p.x+1][p.y] == 0 && sumFunc(p.x+1, p.y) <= k {
			queue.PushBack(position{
				x: p.x + 1,
				y: p.y,
			})
			visited[p.x+1][p.y] = 1
			result++
		}
		if p.x < m && p.y+1 < n && visited[p.x][p.y+1] == 0 && sumFunc(p.x, p.y+1) <= k {
			queue.PushBack(position{
				x: p.x,
				y: p.y + 1,
			})
			visited[p.x][p.y+1] = 1
			result++
		}
	}

	return result
}

func sumFunc(x, y int) int {
	return x%10 + x/10 + y%10 + y/10
}

func main() {
	fmt.Println(movingCount(16, 8, 4))
}
