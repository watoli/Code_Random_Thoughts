# Day31贪心算法part01

## 455.分发饼干 
1. 贪心
```go
func findContentChildren(g []int, s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	lenS := len(s)
	max := s[lenS-1]
	fast := 0
	slow := 0

	for ; fast < len(g); fast++ {
		if g[fast] > max || slow >= lenS {
			break
		}
		for ; slow < len(s); slow++ {
			if s[slow] >= g[fast] {
				slow++
				break
			}
		}
	}
	return fast
}
```

## 376. 摆动序列
1. 贪心
```go
func wiggleMaxLength(nums []int) int {
	numsL := len(nums)
	if numsL <= 1 {
		return numsL
	}
	slow := 1
	res := 1
	flagN := true

	for ; slow < numsL; slow++ {
		if nums[slow] != nums[slow-1] {
			if nums[slow] > nums[slow-1] {
				flagN = true
			} else {
				flagN = false
			}
			res++
			break
		}
	}
	if slow == numsL-1 {
		return res
	}
	for i := slow + 1; i < numsL; i++ {
		if flagN {
			if nums[i] < nums[i-1] {
				flagN = false
				res++
			}
		} else {
			if nums[i] > nums[i-1] {
				flagN = true
				res++
			}
		}
	}
	return res
}
```

## 53. 最大子序和
1. 贪心
```go
func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
```