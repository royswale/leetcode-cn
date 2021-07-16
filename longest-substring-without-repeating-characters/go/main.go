package main

import "fmt"

func main() {
	fmt.Println("hello golang")

	var len = lengthOfLongestSubstring("hello")
	fmt.Println(len)
}
func lengthOfLongestSubstring(s string) int {
	var i, j, max int
	var m = make(map[byte]int)
	for j = 0; j < len(s); j++ {
		var c = s[j]
		if _, ok := m[c]; !ok {
			m[c] = 0
		}
		m[c] += 1

		for m[c] > 1 {
			m[s[i]] -= 1
			i++
		}

		if j-i+1 > max {
			max = j - i + 1
		}
	}
	return max
}
