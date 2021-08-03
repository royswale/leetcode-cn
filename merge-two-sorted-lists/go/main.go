package main

import (
	"fmt"
)

func main() {
	var list1 []int = []int{1,2,4}
	var list2 []int = []int{1,3,4}
	fmt.Println(list1)
	fmt.Println(list2)
	printList(makeList(list1))
	printList(makeList(list2))

	printList(mergeTwoLists(makeList(list1), makeList(list2)))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(values []int) *ListNode {
	var head *ListNode
	var p *ListNode
	for _, value := range values {
		var ptr = &ListNode{
			Val: value,
			Next: nil,
		}
		if head == nil {
			head = ptr
			p = ptr
		} else {
			p.Next = ptr
			p = p.Next
		}
	}
	return head
}
func printList(list *ListNode) {
	fmt.Print("[")
	for list != nil {
		fmt.Printf("%d,", list.Val)
		list = list.Next
	}
	fmt.Println("]")
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var p *ListNode

	//var c = make([]int, 0)

	for !( l1 == nil && l2 == nil ) {
		var i int
		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				i = l1.Val
				l1 = l1.Next
			} else {
				i = l2.Val
				l2 = l2.Next
			}
		} else if l1 != nil {
			i = l1.Val
			l1 = l1.Next
		} else {
			i = l2.Val
			l2 = l2.Next
		}
		//c = append(c, i)
		//var ptr = &ListNode{

		//	Val: i,
		//	Next: nil,
		//}
		if head == nil {
			//head = ptr
			//p = ptr

			head = &ListNode{
				Val: i,
				Next: nil,
			}
			p = head
		} else {
			//p.Next = ptr
			//p = p.Next

			p.Next = &ListNode{
				Val: i,
				Next: nil,
			}
			p = p.Next
		}
	}
	//fmt.Println(c)
	return head
}

// Example 1:

// Input: l1 = [1,2,4], l2 = [1,3,4]
// Output: [1,1,2,3,4,4]
// Example 2:

// Input: l1 = [], l2 = []
// Output: []
// Example 3:

// Input: l1 = [], l2 = [0]
// Output: [0]
//

// Constraints:

// The number of nodes in both lists is in the range [0, 50].
// -100 <= Node.val <= 100
// Both l1 and l2 are sorted in non-decreasing order.

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。