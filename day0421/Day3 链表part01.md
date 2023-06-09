# Day3 链表part01

## 203. 移除链表元素

1. 迭代方法，写的比较繁琐
```go
package main

import . "Code_Random_Thoughts/mypkg"

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	//res := new(ListNode)
	curNode := head
	for curNode.Val == val {
		curNode = curNode.Next
		if curNode == nil {
			return nil
		}

	}
	newHead := curNode
	for curNode.Next != nil {
		preNode := curNode
		curNode = curNode.Next
		for curNode.Val == val {
			curNode = curNode.Next
			if curNode == nil {
				preNode.Next = nil
				return newHead
			}
		}
		preNode.Next = curNode
	}

	return newHead
}

```

2. 精简后的迭代写法
```go
package main

import . "Code_Random_Thoughts/mypkg"

func removeElements(head *ListNode, val int) *ListNode {
	res := new(ListNode)
	res.Next = head
	tmp := res
	for tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
			//tmp = tmp.Next
		} else {
			tmp = tmp.Next
		}
	}
	return res.Next
}
```

## 707. 设计链表
> 上一道题没意识到使用了虚拟头节点，结果这道题吃大亏了  
> 虚拟头节点YY DS!
```go
package main
type MyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	val  int
	pre  *Node
	next *Node
}

func Constructor() MyLinkedList {
	dummyHead := &Node{}
	dummyTail := &Node{}
	dummyHead.next = dummyTail
	dummyTail.pre = dummyHead

	return MyLinkedList{head: dummyHead, tail: dummyTail, length: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.length || index < 0 {
		return -1
	}
	var tmp *Node
	if this.length-index > index {
		tmp = this.head
		for i := 0; i <= index; i++ {
			tmp = tmp.next
		}
	} else {
		tmp = this.tail
		for i := 0; i < this.length-index; i++ {
			tmp = tmp.pre
		}
	}
	return tmp.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.length, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.length || index < 0 {
		return
	}
	var pre, next *Node
	if this.length-index > index {
		pre = this.head
		for i := 0; i < index; i++ {
			pre = pre.next
		}
		next = pre.next
	} else {
		next = this.tail
		for i := 0; i < this.length-index; i++ {
			next = next.pre
		}
		pre = next.pre

	}
	addNode := &Node{
		val:  val,
		pre:  pre,
		next: next,
	}
	this.length++
	pre.next = addNode
	next.pre = addNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.length || index < 0 {
		return
	}
	var pre *Node
	if this.length-index > index {
		pre = this.head
		for i := 0; i < index; i++ {
			pre = pre.next
		}
		//pre = next.pre
		pre.next.next.pre = pre
		pre.next = pre.next.next
	} else {
		pre = this.tail
		for i := 0; i < this.length-index; i++ {
			pre = pre.pre
		}

		pre.pre.next = pre.next
		pre.next.pre = pre.pre

	}
	this.length--

}
```
## 206. 反转链表
1. 迭代解法
```go
package main
import . "Code_Random_Thoughts/mypkg"
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head.Next
	head.Next = nil
	pre := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
```

2. 递归解法

```go
package main
import . "Code_Random_Thoughts/mypkg"
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
```