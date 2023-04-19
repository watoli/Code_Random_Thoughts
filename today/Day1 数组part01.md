# Day1

## 704.二分查找

1. 解题代码想到的是左闭右闭

```go
package main

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for {
		if end < start {
			return -1
		} else if nums[(end-start)/2+start] == target {
			return (end-start)/2 + start
		}
		if nums[(end-start)/2+start] < target {
			start = (end-start)/2 + start + 1
		} else {
			end = (end-start)/2 + start - 1
		}
	}
}
```

2. 利用ChatGPT生成了测试用例

```go
package main

import "testing"

func TestSearch(t *testing.T) {
	// Test case 1
	nums1 := []int{-1, 0, 3, 5, 9, 12}
	target1 := 9
	expected1 := 4
	result1 := search(nums1, target1)
	if result1 != expected1 {
		t.Errorf("search(%v, %d) = %d; expected %d", nums1, target1, result1, expected1)
	}

	// Test case 2
	nums2 := []int{-1, 0, 3, 5, 9, 12}
	target2 := 2
	expected2 := -1
	result2 := search(nums2, target2)
	if result2 != expected2 {
		t.Errorf("search(%v, %d) = %d; expected %d", nums2, target2, result2, expected2)
	}

	// Test case 3
	nums3 := []int{1}
	target3 := 1
	expected3 := 0
	result3 := search(nums3, target3)
	if result3 != expected3 {
		t.Errorf("search(%v, %d) = %d; expected %d", nums3, target3, result3, expected3)
	}

	// Test case 4
	nums4 := []int{5, 7, 8, 9, 10}
	target4 := 6
	expected4 := -1
	result4 := search(nums4, target4)
	if result4 != expected4 {
		t.Errorf("search(%v, %d) = %d; expected %d", nums4, target4, result4, expected4)
	}
}
```

## 27.移除元素

1. 暴力解法
```go
package main

func removeElement(nums []int, val int) int {
	sum := len(nums)
	for i := 0; i < sum; i++ {
		if nums[i] == val {
			sum--
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return sum
}
```
2. 双指针解法
```go
package main
func removeElement(nums []int, val int) int {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[slow] = nums[i]
			slow++
		}
	}
	return slow
}
```

3. 相向双指针解法
```go
package main
func removeElement(nums []int, val int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		for left <= right && nums[left] != val {
			left++
		}
		for left <= right && nums[right] == val {
			right--
		}
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	return left
}
```