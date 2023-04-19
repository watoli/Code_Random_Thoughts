package main

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for {
		if end < start {
			return -1
		} else if nums[(end-start)/2+start] == target {
			return (end-start)/2 + start
		}
		if nums[(end-start)/2+start] < target {
			start = (end-start)/2 + start + 1
		} else {
			end = (end-start)/2 + start - 1
		}
	}
}

func removeElement(nums []int, val int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		for left <= right && nums[left] != val {
			left++
		}
		for left <= right && nums[right] == val {
			right--
		}
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	return left
}
