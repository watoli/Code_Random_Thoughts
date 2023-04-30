package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(x, y int) bool {
	return h[x] > h[y]
}

func (h *MaxHeap) Swap(x, y int) {
	(*h)[x], (*h)[y] = (*h)[y], (*h)[x]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	lenH := len(*h)
	res := (*h)[lenH-1]
	*h = (*h)[:lenH-1]
	return res
}

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

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	h := &MinHeap{}
	heap.Init(h)
	fmt.Println(h)
	heap.Push(h, 3)
	heap.Push(h, 2)
	heap.Push(h, 1)
	fmt.Println(h)
	fmt.Printf("minimum: %d\n", (*h)[0]) // minimum: 1
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h)) // 1 2 3 5
	}
}

//func main() {
//	heepA := &MinHeap{4, 5, 6}
//
//	heepA.Push(1)
//	heepA.Push(2)
//	heepA.Push(3)
//	heap.Init(heepA)
//	fmt.Println(heepA)
//	fmt.Println(heepA.Pop())
//	fmt.Println(heepA.Pop())
//	fmt.Println(heepA.Pop())
//
//	heepA.Push(1)
//	heepA.Push(2)
//	heepA.Push(3)
//	heap.Init(heepA)
//	fmt.Println(heepA)
//	fmt.Println(heepA.Pop())
//	fmt.Println(heepA.Pop())
//	fmt.Println(heepA.Pop())
//}
