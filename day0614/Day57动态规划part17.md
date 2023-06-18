# 第九章 动态规划part17
## 647. 回文子串
1. 动态规划
```go
func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	res := 0
	for i := 0; i < len(dp); i++ {
		tmp := make([]bool, len(s))
		dp[i] = tmp
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if (j - i) <= 1 {
					dp[i][j] = true
					res++
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					res++
				}
			}
		}
	}
	return res
}
```
2. 中心扩散
```go
func countSubstrings(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result += extend(s, i, i, len(s))   // 以i为中心
		result += extend(s, i, i+1, len(s)) // 以i和i+1为中心
	}
	return result
}

func extend(s string, i int, j int, n int) int {
	res := 0
	for i >= 0 && j < n && s[i] == s[j] {
		i--
		j++
		res++
	}
	return res
}
```
## 516.最长回文子序列
```go
func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s)+1)
	res := 0
	for i := 0; i < len(dp); i++ {
		tmp := make([]int, len(s)+1)
		//for j := 0; j < len(tmp); j++ {
		//	tmp[j] = -1
		//}
		dp[i] = tmp
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if (j - i) <= 1 {
					dp[i][j] = j - i + 1
				} else {
					dp[i][j] = dp[i+1][j-1] + 2
				}
				if dp[i][j] > res {
					res = dp[i][j]
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return res
}
```