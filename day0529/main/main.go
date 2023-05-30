package main

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func integerBreak(n int) int {
	if n <= 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 0; i <= n; i++ {
		for j := 0; j < i/2; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
