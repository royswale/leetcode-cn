package main

import "fmt"

func main() {
	fmt.Println(isValid("[{}]"))
	fmt.Println(isValid("[{}"))
	fmt.Println(isValid("{}]"))
}

func isValid(s string) bool {
	// var stack []uint8
	// var bytesArray = []byte(s)
	// for i := 0; i < len(bytesArray); i++ {
	// }

	var stack []rune
	for _, c := range s {
		if c == '[' || c == '(' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			switch stack[len(stack)-1] {
			case '[':
				if c != ']' {
					return false
				}
			case '(':
				if c != ')' {
					return false
				}
			case '{':
				if c != '}' {
					return false
				}
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

/*
Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
Example 4:

Input: s = "([)]"
Output: false
Example 5:

Input: s = "{[]}"
Output: true


Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
