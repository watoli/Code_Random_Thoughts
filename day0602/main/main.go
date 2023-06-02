package main

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= 2; j++ {
			if i-j >= 0 {
				dp[i] += dp[i-j]
			}
		}
	}
	return dp[n]
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

func numSquares(n int) int {
	numS := make([]int, 101)
	for i := 0; i < len(numS); i++ {
		numS[i] = i * i
	}
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = n + 1
	}
	for i := 1; i < len(numS); i++ {
		for j := numS[i]; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-numS[i]]+1)
		}
	}
	return dp[n]
}
