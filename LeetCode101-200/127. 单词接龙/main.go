package main

import "fmt"

// https://leetcode-cn.com/problems/word-ladder/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dic := make(map[string]struct{})
	for _, word := range wordList {
		dic[word] = struct{}{}
	}

	if len(dic) == 0 {
		return 0
	}
	if _, ok := dic[endWord]; !ok {
		return 0
	}

	visited1 := make(map[string]struct{})
	visited1[beginWord] = struct{}{}
	visited2 := make(map[string]struct{})
	visited2[endWord] = struct{}{}

	queue1 := make([]string, 0)
	queue1 = append(queue1, beginWord)
	queue2 := make([]string, 0)
	queue2 = append(queue2, endWord)

	res := 1
	for len(queue1) > 0 {
		if len(queue1) > len(queue2) {
			queue1, queue2 = queue2, queue1
			visited1, visited2 = visited2, visited1
		}
		nextQueue := make([]string, 0)
		for _, word := range queue1 {
			wordBytes := []byte(word)
			for j, c := range wordBytes {
				// 把每个位置的字母都用 a ~ z 替换一下 看看能不能在字典中找到 构成一个边
				for i := 'a'; i <= 'z'; i++ {
					wordBytes[j] = byte(i)
					w := string(wordBytes)
					if _, ok := visited1[w]; !ok {
						// 没有走过
						if _, ok := dic[w]; ok {
							// 标记走过
							nextQueue = append(nextQueue, w)
							visited1[w] = struct{}{}
							if _, ok := visited2[w]; ok {
								return res + 1
							}
						}
					}
					// 恢复原样
					wordBytes[j] = c
				}
			}
		}
		queue1 = nextQueue
		res++
	}
	return 0
}

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
