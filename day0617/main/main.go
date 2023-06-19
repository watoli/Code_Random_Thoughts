package main

// Stack 是用于存放 int 的 栈
type Stack struct {
	nums []int
}

// NewStack 返回 *kit.Stack
func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

// Push 把 n 放入 栈
func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

// Pop 从 s 中取出最后放入 栈 的值
func (s *Stack) Pop() int {
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

// Len 返回 s 的长度
func (s *Stack) Len() int {
	return len(s.nums)
}

// IsEmpty 反馈 s 是否为空
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Top() int {
	return s.nums[s.Len()-1]
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func largestRectangleArea(heights []int) int {
	lenH := len(heights)
	numS := Stack{nums: []int{}}
	left := 0
	mid := 0
	res := 0
	numS.Push(0)
	for i := 0; i < lenH; i++ {
		mid = heights[i]
		if heights[i] < heights[numS.Top()] {
			for !numS.IsEmpty() && heights[i] < heights[numS.Top()] {
				left = heights[numS.Top()]
				res = max(res, min())
			}
		}
		numS.Push(i)
	}
}
