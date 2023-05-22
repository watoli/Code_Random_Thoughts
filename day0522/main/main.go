package main

import (
	"fmt"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] < 0 {
			if nums[j] < 0 {
				return -nums[i] < -nums[j]
			} else {
				return -nums[i] < nums[j]
			}
		} else if nums[j] < 0 {
			return nums[i] < -nums[j]
		}
		return nums[i] < nums[j]
	})
	sum := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		}
		sum += nums[i]
	}
	fmt.Println(nums)
	if k > 0 {
		if k%2 != 0 {
			sum -= nums[0] * 2
		}
	}
	return sum

}

/*
gas = [1,2,3,4,5],
cost = [3,4,5,1,2]
[-2,-2,-2,3,3]
[-1,-1,1]
*/
func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	rGas := 0
	sGas := 0
	for i := 0; i < len(gas); i++ {
		sGas += gas[i] - cost[i]
		rGas += gas[i] - cost[i]
		if rGas < 0 {
			start = i + 1
			rGas = 0
		}
	}
	if sGas < 0 {
		return -1
	}
	return start
}

func candy(ratings []int) int {
	lenR := len(ratings)
	if lenR == 1 {
		return 1
	}
	candyS := make([]int, lenR)
	candyS[0] = 1
	for i := 1; i < lenR; i++ {
		if ratings[i] > ratings[i-1] {
			candyS[i] = candyS[i-1] + 1
		}
	}
	if candyS[lenR-1] == 0 {
		candyS[lenR-1] = 1
	}
	sumC := candyS[lenR-1]
	for i := lenR - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candyS[i] < candyS[i+1]+1 {
				candyS[i] = candyS[i+1] + 1
			}
		}
		sumC += candyS[i]
	}
	return sumC
}
