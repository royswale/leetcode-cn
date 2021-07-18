package main

import "fmt"

func main() {
	fmt.Println("rotate-image")
	// var matrix = make([][]int, 2)
	// for _, v := range v {
	// }
	var matrix = [][]int{{1, 2}, {3, 4}}
	rotate(matrix)

	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
	// Output: [[7,4,1],[8,5,2],[9,6,3]]

	rotate([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}})
	// Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
	// Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

}

func rotate1(matrix [][]int) {
	fmt.Println(matrix)

	var n = len(matrix)
	// fmt.Println(n)

	if n == 1 {
		return
	}

	var s = make([]int, n)
	// fmt.Println(s)

	// var left int = 0
	// var right int = n - 1
	// fmt.Println(left, right)

	print(matrix)
	// return

	// tmp store the most top line in s
	// s = matrix[0]
	copy(s, matrix[0])

	// left most line rotate to top most line
	for i := 0; i < n; i++ {
		matrix[0][n-1-i] = matrix[i][0]
	}
	print(matrix)
	// down most line rotate to left most line
	for i := 0; i < n; i++ {
		matrix[i][0] = matrix[n-1][i]
	}
	print(matrix)
	// right most line rotate to down most line
	for i := 0; i < n; i++ {
		matrix[n-1][i] = matrix[n-1-i][n-1]
	}
	print(matrix)
	// restore tmp to the most right line
	for i := 0; i < n; i++ {
		matrix[i][n-1] = s[i]
	}
	print(matrix)

	// TODO: recursively rotate
	// pass slice of inner matrix
}
func print(matrix [][]int) {
	fmt.Println("--------------")
	for _, v := range matrix {
		fmt.Println(v)
	}
	fmt.Println("--------------")
}
func rotate(matrix [][]int) {
	// fmt.Println(matrix)

	var n = len(matrix)
	// fmt.Println(n)
	if n == 1 {
		return
	}

	var s = make([]int, n)
	// fmt.Println(s)

	rotatex(matrix, 0, n-1, s)

}

// right is include
func rotatex(matrix [][]int, left int, right int, s []int) {
	// println("------rotatex------")
	// print(matrix)
	// fmt.Printf("left: %d, right: %d\n", left, right)

	if left >= right {
		return
	}

	var n = len(matrix)

	// temporarily store the most top line in s
	// CAUTION: this is reference
	// s = matrix[0]
	// CAUTION: use copy instead
	// copy(s, matrix[0])
	// copy the whole left'th line of matrix
	copy(s, matrix[left])
	// fmt.Printf("copy top to tmp: %v\n", s)

	// left most line rotate to top most line
	for i := left; i <= right; i++ {
		matrix[left][n-1-i] = matrix[i][left]
	}
	// println("after left to top...")
	// print(matrix)

	// down most line rotate to left most line
	for i := left; i <= right; i++ {
		matrix[i][left] = matrix[right][i]
	}
	// println("after down to left...")
	// print(matrix)

	// right most line rotate to down most line
	for i := left; i <= right; i++ {
		matrix[right][i] = matrix[n-1-i][right]
	}
	// println("after right to down...")
	// print(matrix)

	// restore tmp to the most right line
	for i := left; i <= right; i++ {
		// matrix[n-i][right] = s[n-i]
		matrix[i][right] = s[i]
	}
	// println("after [tmp] top to right...")
	// print(matrix)

	// TODO: recursively rotate
	// pass slice of inner matrix
	if left+1 < right-1 {
		rotatex(matrix, left+1, right-1, s)
	}
}
