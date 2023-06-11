# 第九章 动态算法part08
## 139.单词拆分 
```go
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
```
## 多重背包
转化为01背包问题求解
```go
// 多重背包可以化解为 01 背包
func multiplePack(weight, value, nums []int, bagWeight int) int {

	for i := 0; i < len(nums); i++ {
		for nums[i] > 1 {
			weight = append(weight, weight[i])
			value = append(value, value[i])
			nums[i]--
		}
	}
	log.Println(weight)
	log.Println(value)

	res := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[i]; j-- {
			res[j] = getMax(res[j], res[j-weight[i]]+value[i])
		}
		log.Println(res)
	}

	return res[bagWeight]
}
```