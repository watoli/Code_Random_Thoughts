# 第九章 动态规划part10
## 121. 买卖股票的最佳时机
```go
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		tmp := make([]int, 2)
		dp[i] = tmp
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[len(prices)-1][0]
}
```
## 122.买卖股票的最佳时机II
```go
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
```