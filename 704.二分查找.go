/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	left := 0
	right := len(nums)-1
	for  {
		if left > right {
			return -1
		}
		index := (left+right)/2
		mid := nums[index]
		if mid == target {
			return index
		}
		if target < mid {
			right = index - 1
		} else {
			left = index + 1
		}
	}
}
// @lc code=end

