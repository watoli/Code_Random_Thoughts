package main

import (
	"fmt"
	. "github.com/halfrost/LeetCode-Go/structures"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		twoNode := cur.Next
		threeNode := cur.Next.Next
		fourNode := cur.Next.Next.Next
		cur.Next = threeNode
		threeNode.Next = twoNode
		twoNode.Next = fourNode
		cur = twoNode
	}
	return dummyHead.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	fast := dummyHead
	slow := dummyHead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	//fast = fast.Next
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	for curA != curB {
		if curA == nil {
			curA = headB

		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA

		} else {
			curB = curB.Next
		}
	}
	return curA
}

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	slow = slow.Next
	fast = head.Next.Next
	for fast != nil && fast.Next != nil && fast != slow {
		fast = fast.Next.Next
		slow = slow.Next

	}
	if fast == slow {
		res := head
		for res != slow {
			res = res.Next
			slow = slow.Next
		}
		return res
	} else {
		return nil
	}

}

func main() {
	headA := new(ListNode)
	headB := new(ListNode)
	if headA.Next == headB.Next {
		fmt.Println("yes")
	}
}
