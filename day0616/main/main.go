package main

import "fmt"

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

func nextGreaterElements(temperatures []int) []int {
	lenT := len(temperatures)
	res := make([]int, lenT)
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	numS := Stack{nums: []int{}}
	numS.Push(0)
	temperatures = append(temperatures, temperatures...)
	for i := 0; i < len(temperatures); i++ {
		if temperatures[numS.Top()] < temperatures[i] {
			for !numS.IsEmpty() && temperatures[numS.Top()] < temperatures[i] {
				res[numS.Top()] = temperatures[i]
				fmt.Println("res:", numS.Top(), " ", i%lenT-numS.Top())
				numS.Pop()
			}
		}
		if i < lenT {
			numS.Push(i % lenT)
		} else if numS.IsEmpty() {
			break
		}
		fmt.Println("i:", temperatures[i], " ", numS.nums)
	}
	for !numS.IsEmpty() {
		res[numS.Top()] = -1
		numS.Pop()
	}
	return res
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func trap(height []int) int {
	lenT := len(height)
	right := 0
	mid := 0
	res := 0
	numS := Stack{nums: []int{}}
	numS.Push(0)
	for i := 0; i < lenT; i++ {
		if height[numS.Top()] < height[i] {
			right = height[i]
			for !numS.IsEmpty() && height[numS.Top()] < height[i] {
				mid = height[numS.Pop()]
				if !numS.IsEmpty() {
					res += (min(right, height[numS.Top()]) - mid) * (i - numS.Top() - 1)
				}
			}
		} else if height[numS.Top()] == height[i] {
			numS.Pop()
			//numS.Push(i)
		}
		numS.Push(i)
	}
	return res
}
