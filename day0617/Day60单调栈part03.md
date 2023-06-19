# 第十章 单调栈part03
## 84.柱状图中最大的矩形
```go
func largestRectangleArea(heights []int) int {
	heights = append([]int{0}, append(heights, 0)...)
	lenH := len(heights)
	numS := Stack{nums: []int{}}
	mid := 0
	res := 0
	numS.Push(0)
	for i := 0; i < lenH; i++ {
		if heights[i] < heights[numS.Top()] {
			for !numS.IsEmpty() && heights[i] < heights[numS.Top()] {
				mid = heights[numS.Top()]
				numS.Pop()
				res = max(res, mid*(i-numS.Top()-1))
			}
		}
		numS.Push(i)
	}
	return res
}
```
