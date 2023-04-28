package main

type Queue struct {
	data []int
	len  int
}

func NewQueue() *Queue {
	return &Queue{
		data: []int{},
		len:  0,
	}
}

func (q *Queue) Push(num int) {
	q.data = append(q.data, num)
	q.len++
}

func (q *Queue) Pop() int {
	res := q.data[0]
	q.data = q.data[1:]
	q.len--
	return res
}

func (q *Queue) Peek() int {
	return q.data[0]
}

func (q *Queue) Size() int {
	return q.len
}

func (q *Queue) Empty() bool {
	if q.len == 0 {
		return true
	} else {
		return false
	}
}

type Stack struct {
	data []int
	len  int
}

func NewStack() *Stack {
	return &Stack{
		data: []int{},
		len:  0,
	}
}

func (s *Stack) Push(num int) {
	s.data = append(s.data, num)
	s.len++
}

func (s *Stack) Pop() int {
	res := s.data[s.len-1]
	s.data = s.data[:s.len-1]
	s.len--
	return res
}

func (s *Stack) Peek() int {
	return s.data[s.len-1]
}

func (s *Stack) Size() int {
	return s.len
}
func (s *Stack) Empty() bool {
	if s.len == 0 {
		return true
	} else {
		return false
	}
}

type MyQueue struct {
	inStack  Stack
	outStack Stack
}

func ConstructorStack() MyQueue {
	inS := *NewStack()
	outS := *NewStack()
	return MyQueue{inStack: inS, outStack: outS}
}

func (q *MyQueue) Push(x int) {
	q.inStack.Push(x)
}

func (q *MyQueue) Pop() int {
	if q.outStack.Empty() {
		for !q.inStack.Empty() {
			q.outStack.Push(q.inStack.Pop())
		}
	}
	return q.outStack.Pop()
}

func (q *MyQueue) Peek() int {
	if q.outStack.Empty() {
		for !q.inStack.Empty() {
			q.outStack.Push(q.inStack.Pop())
		}
	}
	return q.outStack.Peek()
}

func (q *MyQueue) Empty() bool {
	return q.outStack.Empty() && q.inStack.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type MyStack struct {
	cycle Queue
}

func Constructor() MyStack {
	Cq := *NewQueue()
	return MyStack{cycle: Cq}
}

func (s *MyStack) Push(x int) {
	s.cycle.Push(x)
}

func (s *MyStack) Pop() int {
	for i := 0; i < s.cycle.len-1; i++ {
		s.cycle.Push(s.cycle.Pop())
	}
	return s.cycle.Pop()
}

func (s *MyStack) Top() int {
	for i := 0; i < s.cycle.len-1; i++ {
		s.cycle.Push(s.cycle.Pop())
	}
	res := s.cycle.Peek()
	s.cycle.Push(s.cycle.Pop())
	return res
}

func (s *MyStack) Empty() bool {
	return s.cycle.Empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
