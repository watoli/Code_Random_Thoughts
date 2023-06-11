# 第九章 动态规划part09
## 198.打家劫舍
```go
func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

/*
2 1 1 2
2 1 2
*/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
		fmt.Println(dp)
	}
	return dp[len(nums)-1]
}
```
## 213.打家劫舍II
```go
func robII(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	} else if len(nums) == 3 {
		return max(nums[0], max(nums[1], nums[2]))
	}
	dp := make([]int, len(nums))
	dp1 := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	dp1[0], dp1[1], dp1[2] = 0, nums[1], max(nums[1], nums[2])
	for i := 2; i < len(nums); i++ {
		if i == len(nums)-1 {
			dp[i] = max(dp[i-1], nums[i]+dp1[i-2])
		} else {
			dp[i] = max(dp[i-1], nums[i]+dp[i-2])
			dp1[i] = max(dp1[i-1], nums[i]+dp1[i-2])
		}
	}
	return dp[len(nums)-1]
}
```
## 337.打家劫舍III
```go

```
