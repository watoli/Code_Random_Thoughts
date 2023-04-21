package main

import (
	. "Code_Random_Thoughts/mypkg"
	"fmt"
)

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

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	ptr := &Node{}
	fmt.Println(ptr)
}
