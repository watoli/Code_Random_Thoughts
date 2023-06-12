package main

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
