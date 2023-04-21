# Day4 链表part2

## 24. 两两交换链表中的节点
1. 使用虚拟头节点加别名命名，真好做~
```go
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


```
##  19.删除链表的倒数第N个节点
```go
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
```

## 面试题 02.07. 链表相交
1. 哈希存储节点
```go
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]bool)
	for headA != nil {
		nodeMap[headA] = true
		headA = headA.Next
	}
	for headB != nil {

		if _, ok := nodeMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
```
2. 双指针
```go
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
```

## 142. 环形链表 II
1. 哈希
```go
func detectCycle(head *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := nodeMap[head]; ok {
			return head
		}
		nodeMap[head] = true
		head = head.Next
	}
	return nil

}
```
2. 快慢指针
```go
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
```