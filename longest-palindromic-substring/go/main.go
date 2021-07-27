package main

import "fmt"

func main() {

	fmt.Println(longestPalindrome("cbabba"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("abcd"))
	fmt.Println(longestPalindrome("abcbd"))
}

// https://leetcode-cn.com/problems/longest-palindromic-substring/
// https://leetcode-cn.com/problems/longest-palindromic-substring/solution/zui-chang-hui-wen-zi-chuan-by-leetcode-solution/
// 方法一：动态规划
/*
func longestPalindrome(s string) string {
	var len = len(s)
	if len == 1 {
		return s
	}

	var maxLen int = 1
	var start int = 0

	var dp [][]bool = make([][]bool, len)
	for i := 0; i < len; i++ {
		dp[i] = make([]bool, len)

		// no need
		// dp[i][i] = true
	}

	for L := 2; L <= len; L++ {

		for i := 0; i < len; i++ {
			var j = L + i - 1

			if j >= len {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
				continue
			} else {
				if L <= 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && L > maxLen {
				maxLen = L
				start = i
			}
		}
	}

	return s[start : start+maxLen]
}
*/
// 方法二：中心扩展算法
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}

	return s[start : end+1]
}
func expandAroundCenter(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
