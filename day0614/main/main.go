package main

func isSubString(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

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

func findMinimum(numbers ...int) int {
	if len(numbers) == 0 {
		panic("At least one number must be provided")
	}

	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}

	return min
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

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
