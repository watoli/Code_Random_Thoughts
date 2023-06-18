# 第十章 单调栈part01
## 739. 每日温度
```go
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
```
## 496.下一个更大元素 I
```go
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
```
