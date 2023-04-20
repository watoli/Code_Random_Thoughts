package main

import "fmt"

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, num int) int {
	if root == nil {
		return num
	}
	num++
	left := dfs(root.Left, num)
	right := dfs(root.Right, num)
	if right < left {
		right = left
	}
	if num < right {
		num = right
	}
	return num
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	num := 1
	left := dfs(root.Left, num)
	right := dfs(root.Right, num)
	if right < left {
		right = left
	}
	if num < right {
		num = right
	}
	return num
}
