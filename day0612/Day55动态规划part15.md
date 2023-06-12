#  392.判断子序列 
##  392.判断子序列 
1. 一维动态规划
```go
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	dp := make([]int, len(t)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	for i := 1; i <= len(t); i++ {
		if s[dp[i-1]+1] == t[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}
		if dp[i]+1 == len(s) {
			return true
		}
		fmt.Println(dp)
	}
	if dp[len(t)]+1 == len(s) {
		return true
	}
	return false
}
```
2. 二维动态规划
```go
func isSubsequence(s string, t string) bool {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		tmp := make([]int, len(t)+1)
		dp[i] = tmp
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	if dp[len(s)][len(t)] == len(s) {
		return true
	}
	return false
}
```
## 115.不同的子序列
```go
func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		tmp := make([]int, len(t)+1)
		tmp[0] = 1
		dp[i] = tmp
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(s)][len(t)]
}
```
