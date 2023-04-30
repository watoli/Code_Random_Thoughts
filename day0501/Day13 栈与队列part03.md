# Day13 栈与队列part03

# 239. 滑动窗口最大值
1. 单调队列，这里面push的大于号，真tm的关键！
```go
type MonotonicQueue struct {
	queue []int
	len   int
}

func NewMonotonicQueue() *MonotonicQueue {
	NewQueue := make([]int, 0)
	return &MonotonicQueue{queue: NewQueue, len: 0}
}

func (q *MonotonicQueue) Push(x int) {
	for q.len != 0 && q.queue[q.len-1] < x {
		q.queue = q.queue[:q.len-1]
		q.len--
	}
	q.queue = append(q.queue, x)
	q.len++
}

func (q *MonotonicQueue) Pop() {
	q.queue = q.queue[1:]
	q.len--
}

func (q MonotonicQueue) Peek() int {
	return q.queue[0]
}
func (q MonotonicQueue) Size() int {
	return len(q.queue)
}

func maxSlidingWindow(nums []int, k int) []int {
	lenN := len(nums)
	res := make([]int, lenN-k+1)
	MoQ := NewMonotonicQueue()
	for i := 0; i < k; i++ {
		MoQ.Push(nums[i])
	}
	res[0] = MoQ.Peek()
	for i := k; i < lenN; i++ {
		if MoQ.Size() > 0 && MoQ.Peek() == nums[i-k] {
			MoQ.Pop()
		}
		MoQ.Push(nums[i])
		res[i-k+1] = MoQ.Peek()
	}
	return res
}
func main() {
	var a []int
	b := make([]int, 0)
	c := []int{}

	fmt.Println(a, b, c)
}

```

## 347. 前 K 个高频元素
1. 最小堆实现
```go
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(x, y int) bool {
	return h[x] < h[y]
}

func (h MinHeap) Swap(x, y int) {
	(h)[x], (h)[y] = (h)[y], (h)[x]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	lenH := len(*h)
	res := (*h)[lenH-1]
	*h = (*h)[:lenH-1]
	return res
}

func topKFrequent(nums []int, k int) []int {
	numMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := numMap[nums[i]]; !ok {
			numMap[nums[i]] = 1
		} else {
			numMap[nums[i]]++
		}
	}
	res := &MinHeap{}
	heap.Init(res)
	for _, v := range numMap {
		if len(*res) < k {
			heap.Push(res, v)
		} else {
			if v > (*res)[0] {
				heap.Pop(res)
				heap.Push(res, v)
			}
		}
	}

	min := (*res)[0]
	i := 0
	for key, v := range numMap {
		if v >= min {
			(*res)[i] = key
			i++
		}
		if i >= k {
			break
		}
	}
	return *res
}
```