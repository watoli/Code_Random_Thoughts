# 数组part02

## 977.有序数组的平方 
1. 双指针
```go
package main

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	res := make([]int, len(nums), len(nums))

	left, right := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[left] <= nums[right] {
			res[i] = nums[right]
			right--
		} else {
			res[i] = nums[left]
			left++
		}

	}

	return res
}

```
2. 待更新排序解法

##  209.长度最小的子数组
1. 没审题，第一遍写成检测等于target子数组了
```go
package main
func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	res := -1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == target {
			if res == -1 || (i-left) < res {
				res = i - left
			}
		} else if sum > target {
			for sum > target {
				sum -= nums[left]
				left++
			}
			if sum == target {
				if res == -1 || (i-left) < res {
					res = i - left
				}
			}
		}
	}
	return res + 1
}
```
2. 滑动窗口一遍通过
```go
package main
func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	res := -1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum >= target {
			for sum-nums[left] >= target {
				sum -= nums[left]
				left++
			}
			if res == -1 || (i-left) < res {
				res = i - left
			}
		}
	}
	return res + 1
}
```

## 59. 螺旋矩阵

1. 模拟法
```go
package main
import "fmt"
func generateMatrix(n int) [][]int {
	/*
	*****
	*****
	*****
	*****
	*****
	 */
	res := make([][]int, n, n)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n, n)
	}
	num := 1
	x, y := 0, 0
	length := n
	squ := n * n
	for num <= squ {
		for i := 0; i < length; i++ {
			res[x][y] = num
			num++
			y++
		}
		if num > squ {
			return res
		}
		y--
		x++
		length--
		for i := 0; i < length; i++ {
			res[x][y] = num
			num++
			x++
		}
		if num > squ {
			return res
		}
		x--
		y--
		for i := 0; i < length; i++ {
			res[x][y] = num
			num++
			y--
		}
		if num > squ {
			return res
		}
		x--
		y++
		length--
		fmt.Println(res)
		for i := 0; i < length; i++ {
			res[x][y] = num
			num++
			x--
		}
		if num > squ {
			return res
		}
		x++
		y++
	}
	return res

}
```