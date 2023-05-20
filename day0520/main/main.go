package main

import "fmt"

func maxProfit(prices []int) int {
	pricesL := len(prices)
	if pricesL <= 1 {
		return 0
	}
	res := 0
	for i := 1; i < pricesL; i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

func canJumpOld(nums []int) bool {
	numsL := len(nums)
	if numsL == 1 {
		return true
	}
	nums[numsL-1] = -1
	for i := numsL - 2; i >= 0; i-- {
		fmt.Println(nums, i)
		for j := 1; j <= nums[i] && (j+i) < numsL; j++ {
			if nums[j+i] == -1 {
				nums[i] = -1
				break
			}
		}
	}
	if nums[0] == -1 {
		return true
	}
	return false
}

func canJump(nums []int) bool {
	cover := nums[0]
	for i := 1; i < len(nums); i++ {
		if cover < i {
			return false
		}
		if i+nums[i] > cover {
			cover = i + nums[i]
		}
	}
	return true
}

func jumpDP(nums []int) int {
	numsL := len(nums)
	dp := make([]int, numsL)
	for i := 0; i < numsL; i++ {
		dp[i] = numsL
	}
	dp[numsL-1] = 0
	for i := numsL - 2; i >= 0; i-- {
		min := numsL
		for j := 0; j <= nums[i] && i+j <= (numsL-1); j++ {
			if 1+dp[i+j] < min {
				min = 1 + dp[i+j]
			}
		}
		dp[i] = min
	}
	return dp[0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func jump(nums []int) int {
	numsL := len(nums)
	if numsL == 1 {
		return 0
	}
	res := 0
	cur := 0
	next := 0
	for i := 0; i < numsL-1; i++ {
		next = max(nums[i]+1, next)
		if i == cur {
			if next >= (numsL - 1) {
				res++
				break
			}
			res++
			cur = next
		}
	}
	return res
}
