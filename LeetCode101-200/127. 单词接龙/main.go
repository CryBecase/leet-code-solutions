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

	visited := make(map[string]struct{})
	visited[beginWord] = struct{}{}

	queue := make([]string, 0)
	queue = append(queue, beginWord)

	res := 1
	for len(queue) > 0 {
		queue2 := make([]string, 0)
		for _, word := range queue {
			wordBytes := []byte(word)
			for j, c := range wordBytes {
				// 把每个位置的字母都用 a ~ z 替换一下 看看能不能在字典中找到 构成一个边
				for i := 'a'; i <= 'z'; i++ {
					wordBytes[j] = byte(i)
					w := string(wordBytes)
					if _, ok := visited[w]; !ok {
						// 没有走过
						if _, ok := dic[w]; ok {
							// 标记走过
							queue2 = append(queue2, w)
							visited[w] = struct{}{}
							if w == endWord {
								return res + 1
							}
						}
					}
					// 恢复原样
					wordBytes[j] = c
				}
			}
		}
		queue = queue2
		res++
	}
	return 0
}

func main() {
	fmt.Println(ladderLength("hot", "dog", []string{"hot", "cog", "dog", "tot", "hog", "hop", "pot", "dot"}))
}
