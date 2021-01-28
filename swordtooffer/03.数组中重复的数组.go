package main

//找出数组中重复的数字。

// 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复
// 了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

func findRepeatNumber(nums []int) int {
	showed := map[int]int{}
	for _, i := range nums {
		if _, ok := showed[i]; ok {
			return i
		}
		showed[i] = 0
	}
	return -1
}
