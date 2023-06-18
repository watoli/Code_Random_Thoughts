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

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	numS := Stack{nums: []int{}}
	for i := 0; i < len(temperatures); i++ {
		if numS.IsEmpty() {
			numS.Push(i)
		} else {
			if temperatures[numS.Top()] >= temperatures[i] {
				numS.Push(i)
			} else {
				for !numS.IsEmpty() && temperatures[numS.Top()] < temperatures[i] {
					res[numS.Top()] = i - numS.Top()
					fmt.Println("res:", numS.Top(), " ", i-numS.Top())
					numS.Pop()
				}
				numS.Push(i)
			}
		}
		fmt.Println("i:", temperatures[i], " ", numS.nums)
	}
	for !numS.IsEmpty() {
		res[numS.Top()] = 0
		numS.Pop()
	}
	return res
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	numM := make(map[int]int)
	numS := Stack{nums: []int{}}
	for i := 0; i < len(nums2); i++ {
		if numS.IsEmpty() {
			numS.Push(i)
		} else {
			if nums2[numS.Top()] >= nums2[i] {
				numS.Push(i)
			} else {
				for !numS.IsEmpty() && nums2[numS.Top()] < nums2[i] {
					numM[nums2[numS.Top()]] = nums2[i]
					numS.Pop()
				}
				numS.Push(i)
			}
		}
	}
	for !numS.IsEmpty() {
		numM[nums2[numS.Top()]] = -1
		numS.Pop()
	}
	for i := 0; i < len(nums1); i++ {
		nums1[i] = numM[nums1[i]]
	}
	return nums1
}
