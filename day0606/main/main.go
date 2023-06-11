package main

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
func maxProfit(prices []int) int {
	dp := [2][2]int{}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = max(dp[(i-1)%2][0], -prices[i])
		dp[i%2][1] = max(dp[(i-1)%2][1], dp[(i-1)%2][0]+prices[i])
	}

	return dp[(len(prices)-1)%2][1]
}

func maxProfitII(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		tmp := make([]int, 2)
		dp[i] = tmp
	}
	dp[0][1] = -prices[0]
	dp[0][0] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}
