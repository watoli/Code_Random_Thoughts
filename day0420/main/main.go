package main

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	res := make([]int, len(nums), len(nums))

	left, right := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[left] <= nums[right] {
			res[i] = nums[right]
			right--
		} else {
			res[i] = nums[left]
			left++
		}

	}

	return res
}

func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	res := -1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum >= target {
			for sum-nums[left] >= target {
				sum -= nums[left]
				left++
			}
			if res == -1 || (i-left) < res {
				res = i - left
			}
		}
	}
	return res + 1
}
