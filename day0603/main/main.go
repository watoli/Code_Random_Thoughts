package main

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]struct{})
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = struct{}{}
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i <= len(s); i++ {
		for j := 0; j <= i; j++ {
			if _, ok := wordMap[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
