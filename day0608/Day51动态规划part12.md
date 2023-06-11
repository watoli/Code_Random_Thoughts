# 第九章 动态规划part12
## 309.最佳买卖股票时机含冷冻期
```go
func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		tmp := make([]int, 3)
		dp[i] = tmp
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = dp[i-1][1] + prices[i]
	}
	return max(dp[len(prices)-1][0], dp[len(prices)-1][2])

}
```
## 714.买卖股票的最佳时机含手续费
```go
func maxProfitII(prices []int, fee int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		tmp := make([]int, 2)
		dp[i] = tmp
	}
	dp[0][1] = -prices[0]
	dp[0][0] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}
```