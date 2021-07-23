package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	// Input: root = [1,2,2,3,4,4,3]
	// Output: true
	// Input: root = [1,2,2,null,3,null,3]
	// Output: false
	// Constraints:
	// The number of nodes in the tree is in the range [1, 1000].
	// -100 <= Node.val <= 100
	// 来源：力扣（LeetCode）
	// 链接：https://leetcode-cn.com/problems/symmetric-tree

	var root1 *TreeNode = nil
	root1 = insertLevelOrder([]int{1, 2, 2, 3, 4, 4, 3}, root1, 0)
	inOrder(root1)
	fmt.Println()
	fmt.Println(printLevelOrder(root1))
	fmt.Println()
	fmt.Println(isSymmetric(root1))
	fmt.Println()

	var root2 *TreeNode = nil
	root2 = insertLevelOrder([]int{1, 2, 2, -101, 3, -101, 3}, root2, 0)
	inOrder(root2)
	fmt.Println()
	fmt.Println(printLevelOrder(root2))
	fmt.Println()
	fmt.Println(isSymmetric(root2))
	fmt.Println()

	var root3 *TreeNode = nil
	root3 = insertLevelOrder([]int{1, 2, 2, -101, 3, 3, -101}, root3, 0)
	inOrder(root3)
	fmt.Println()
	fmt.Println(printLevelOrder(root3))
	fmt.Println()
	fmt.Println(isSymmetric(root3))
	fmt.Println()

	var root4 *TreeNode = nil
	root4 = insertLevelOrder([]int{1, 2, 2, -101, 3, 3, -101, -101, -101, 5, -101, -101, 5, -101, -101}, root4, 0)
	inOrder(root4)
	fmt.Println()
	fmt.Println(printLevelOrder(root4))
	fmt.Println()
	fmt.Println(isSymmetric(root4))
	fmt.Println()
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}
func compare(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
}

// Construct a complete binary tree from given array in level order fashion - GeeksforGeeks
// https://www.geeksforgeeks.org/construct-complete-binary-tree-given-array/
// Level Order Binary Tree Traversal - GeeksforGeeks
// https://www.geeksforgeeks.org/level-order-tree-traversal/
// Tree Traversals (Inorder, Preorder and Postorder) - GeeksforGeeks
// https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/

func printLevelOrder(root *TreeNode) []int {
	var values []int

	if root == nil {
		return values
	}

	var nodes []*TreeNode
	//var children []*TreeNode = make([]*TreeNode, 0, 0)

	nodes = append(nodes, root)

	for len(nodes) > 0 {
		//var node = nodes[0]
		//nodes = nodes[1:]
		//
		//values = append(values, node.Val)
		//if node.Left != nil {
		//	nodes = append(nodes, node.Left)
		//}
		//if node.Right != nil {
		//	nodes = append(nodes, node.Right)
		//}

		var tmp []*TreeNode
		for _, node := range nodes {
			if node == nil {
				tmp = append(tmp, nil)
				tmp = append(tmp, nil)
			} else {
				tmp = append(tmp, node.Left)
				tmp = append(tmp, node.Right)
			}

			var value int
			if node == nil {
				value = -101
			} else {
				value = node.Val
			}
			values = append(values, value)
		}
		var empty = true
		for _, node := range tmp {
			if node != nil {
				empty = false
				break
			}
		}

		nodes = nodes[:0]
		if !empty {
			for _, node := range tmp {
				nodes = append(nodes, node)
			}
		}

	}

	return values
}

func inOrder(root *TreeNode) {
	if root != nil {
		inOrder(root.Left)
		fmt.Printf("%d ", root.Val)
		inOrder(root.Right)
	}
}

func insertLevelOrder(arr []int, root *TreeNode, i int) *TreeNode {
	if i < len(arr) {
		// val < -100 means null node
		if arr[i] < -100 {
			return root
		}

		root = &TreeNode{
			Val:   arr[i],
			Left:  nil,
			Right: nil,
		}

		root.Left = insertLevelOrder(arr, root.Left, 2*i+1)
		root.Right = insertLevelOrder(arr, root.Right, 2*i+2)
	}
	return root
}
