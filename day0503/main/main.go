package main

import . "Code_Random_Thoughts/mypkg"

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var tmpRes []int
	nodeStack := []*TreeNode{root}
	var tmpNode *TreeNode
	for len(nodeStack) != 0 {
		lenS := len(nodeStack)
		tmpRes = []int{}
		for i := 0; i < lenS; i++ {
			tmpNode = nodeStack[0]
			nodeStack = nodeStack[1:]
			tmpRes = append(tmpRes, tmpNode.Val)
			if tmpNode.Left != nil {
				nodeStack = append(nodeStack, tmpNode.Left)
			}
			if tmpNode.Right != nil {
				nodeStack = append(nodeStack, tmpNode.Right)
			}
		}
		res = append(res, tmpRes)
	}
	return res
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lenS := 0
	nodeStack := []*TreeNode{root.Left, root.Right}
	for len(nodeStack) != 0 {
		var tmpStack []*TreeNode
		lenS = len(nodeStack)
		for i := 0; i < lenS/2; i++ {
			if nodeStack[i] == nil || nodeStack[lenS-1-i] == nil {
				if nodeStack[i] == nil && nodeStack[lenS-1-i] == nil {
					continue
				}
				return false
			} else if nodeStack[i].Val == nodeStack[lenS-1-i].Val {
				tmpStack = append([]*TreeNode{nodeStack[i].Left, nodeStack[i].Right}, tmpStack...)
				tmpStack = append(tmpStack, nodeStack[lenS-1-i].Left, nodeStack[lenS-1-i].Right)
			} else {
				return false
			}
		}
		nodeStack = tmpStack
	}
	return true
}
