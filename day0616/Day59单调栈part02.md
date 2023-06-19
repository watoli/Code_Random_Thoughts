# 第十章 单调栈part02
## 503.下一个更大元素II
```go
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
```
## 42. 接雨水
```go
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
```
