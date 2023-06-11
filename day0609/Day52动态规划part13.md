# 第九章 动态规划part13
## 300.最长递增子序列
```go
func lengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			fmt.Println("i:", i, " j:", j)
			fmt.Println(dp)
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
```
## 674. 最长连续递增序列
```go
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, count := 1, 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			count++
		} else {
			count = 1
		}
		if count > res {
			res = count
		}
	}
	return res
}
```
## 718. 最长重复子数组
```go
func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		tmp := make([]int, len(nums2))
		dp[i] = tmp
	}
	res := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				if i-1 < 0 || j-1 < 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}
	return res
}
```