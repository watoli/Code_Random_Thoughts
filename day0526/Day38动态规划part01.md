# Day38动态规划part01

## 509. 斐波那契数
```go
func fib(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i < len(dp); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
```
## 70. 爬楼梯
```go
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[len(dp)-1] = 0
	dp[len(dp)-2] = 1
	for i := len(dp) - 3; i >= 0; i-- {
		dp[i] = dp[i+1] + dp[i+2]
		fmt.Println(dp)
	}
	return dp[0]
}
```
## 746. 使用最小花费爬楼梯
```go
func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func minCostClimbingStairs(cost []int) int {
	lenC := len(cost)
	if lenC == 2 {
		return min(cost[0], cost[1])
	}
	dp := make([]int, lenC+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= lenC; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
		fmt.Println(dp)
	}
	return dp[lenC]
}
```