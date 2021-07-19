package main

import "fmt"

func main() {
	// fmt.Println("hello")
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
	fmt.Println(findMedianSortedArrays([]int{2}, []int{}))
}

// Example 1:

// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
// Example 2:

// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
// Example 3:

// Input: nums1 = [0,0], nums2 = [0,0]
// Output: 0.00000
// Example 4:

// Input: nums1 = [], nums2 = [1]
// Output: 1.00000
// Example 5:

// Input: nums1 = [2], nums2 = []
// Output: 2.00000

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var len1 = len(nums1)
	var len2 = len(nums2)

	var len = len1 + len2

	if len == 1 {
		if len1 == 1 {
			return float64(nums1[0])
		}
		if len2 == 1 {
			return float64(nums2[0])
		}
	}

	var median_index = len / 2

	var p1 = 0
	var p2 = 0

	var even = len%2 == 0
	var sum = 0

	for i := 0; i <= median_index; i++ {
		if p1 < len1 && (p2 >= len2 || nums1[p1] <= nums2[p2]) {
			if (even && i == median_index-1) || i == median_index {
				sum += nums1[p1]
			}
			// fmt.Printf("i: %d, from nums1 index: %d, value: %d\n", i, p1, nums1[p1])
			p1++
			continue
		}
		if p2 < len2 && (p1 >= len1 || nums2[p2] <= nums1[p1]) {
			if (even && i == median_index-1) || i == median_index {
				sum += nums2[p2]
			}
			// fmt.Printf("i: %d, from nums2 index: %d, value: %d\n", i, p2, nums2[p2])
			p2++
			continue
		}

	}

	if even {
		return float64(sum) / 2
	} else {
		return float64(sum)
	}

	// return 0
}
