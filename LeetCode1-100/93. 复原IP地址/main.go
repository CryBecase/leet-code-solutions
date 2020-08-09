package main

import "strconv"

// https://leetcode-cn.com/problems/restore-ip-addresses/

var (
	ips []int
	ans []string
)

func restoreIpAddresses(s string) []string {
	ips = make([]int, 4)
	ans = make([]string, 0)
	dfs(s, 0, 0)
	return ans
}

func dfs(s string, start, segNum int) {
	if segNum == 4 {
		// 有了 4 个段
		if start == len(s) {
			// 读完 s 了，代表产生一个结果
			ipAddr := ""
			for i := 0; i < 4; i++ {
				ipAddr += strconv.Itoa(ips[i])
				if i != 3 {
					ipAddr += "."
				}
			}
			ans = append(ans, ipAddr)
		}
		// 无论读完没读完都不用继续了
		return
	}

	if start == len(s) {
		// 没有 4 个段，但是读完了，结束
		return
	}

	if s[start] == '0' {
		// 第一个字符是 0，则这一段只能是 0
		ips[segNum] = 0
		dfs(s, start+1, segNum+1)
	}

	addr := 0
	for end := start; end < len(s); end++ {
		addr = addr*10 + int(s[end]-'0')
		if addr > 0 && addr <= 255 {
			ips[segNum] = addr
			dfs(s, end+1, segNum+1)
		} else {
			break
		}
	}

}

func main() {

}
