struct Solution {

}
impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        let mut result: i32 = nums[0];
        let mut sum: i32 = 0;
        for num in nums {
            if sum <= 0 {
                sum = num;
            } else {
                sum += num;
            }
            if sum > result {
                result = sum;
            }
        }

        return result;
    }
}
fn main() {
    let nums = vec![-2,1,-3,4,-1,2,1,-5,4];
    println!("{}", Solution::max_sub_array(nums));
    println!("{}", Solution::max_sub_array(vec![1]));
    println!("{}", Solution::max_sub_array(vec![5,4,-1,7,8]));
}

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
//
// A subarray is a contiguous part of an array.
//
//
// Example 1:
//
// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Example 2:
//
// Input: nums = [1]
// Output: 1
// Example 3:
//
// Input: nums = [5,4,-1,7,8]
// Output: 23
//
// Constraints:
//
// 1 <= nums.length <= 3 * 104
// -105 <= nums[i] <= 105
//
// Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-subarray
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。