package main

import (
	"fmt"
	"math"
)

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func lastStoneWeightII(stones []int) int {
	sum := 0
	sL := len(stones)
	if sL == 1 {
		return stones[0]
	}
	for i := 0; i < sL; i++ {
		sum += stones[i]
	}
	bag := sum / 2
	dp := make([]int, bag+1)
	for i := 0; i < sL; i++ {
		for j := bag; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - dp[bag]*2
}
func abs(x int) int {
	return int(math.Abs(float64(x)))
}
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	numsL := len(nums)
	for i := 0; i < numsL; i++ {
		sum += nums[i]
	}
	if ((sum+target)%2 != 0) || (sum < abs(target)) {
		return 0
	}
	sum = (target + sum) / 2
	dp := make([]int, sum+1)
	dp[0] = 1
	for i := 0; i < numsL; i++ {
		for j := sum; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
			fmt.Println("i:", i, " j:", j, " ", dp)
		}
	}
	return dp[sum]
}

func getNum(str string) (m, n int) {
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			m++
		} else {
			n++
		}
	}
	return
}
func findMaxForm(strs []string, m int, n int) int {
	strS := make([][2]int, len(strs))
	for i := 0; i < len(strS); i++ {
		strS[i][0], strS[i][1] = getNum(strs[i])
	}
	fmt.Println(strS)
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(strS); i++ {
		for j := m; j >= strS[i][0]; j-- {
			for k := n; k >= strS[i][1]; k-- {
				dp[j][k] = max(dp[j][k], dp[j-strS[i][0]][k-strS[i][1]]+1)
			}
		}
	}
	return dp[m][n]
}
