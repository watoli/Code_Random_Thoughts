package main

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		tmp := make([]int, 5)
		dp[i] = tmp
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return max(dp[len(prices)-1][0], max(dp[len(prices)-1][2], dp[len(prices)-1][4]))
}

func maxProfitIV(k int, prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		tmp := make([]int, 2*k+1)
		dp[i] = tmp
	}
	for i := 0; i < 2*k+1; i++ {
		if i%2 == 0 {
			dp[0][i] = 0
		} else {
			dp[0][i] = -prices[0]
		}
	}

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		for j := 1; j < 2*k+1; j++ {
			if j%2 != 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			}
		}
	}
	return dp[len(prices)-1][2*k]
}
