package main

// https://leetcode-cn.com/problems/keys-and-rooms/

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	if n == 0 {
		return true
	}

	visited := make([]int, n)
	visited[0] = 1

	dfs := func(int) {}
	dfs = func(roomIdx int) {
		for _, k := range rooms[roomIdx] {
			if visited[k] == 1 {
				continue
			}
			visited[k] = 1
			dfs(k)
		}
	}

	dfs(0)

	for _, v := range visited {
		if v == 0 {
			return false
		}
	}
	return true
}

func main() {

}
