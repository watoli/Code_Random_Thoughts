package main

import (
	. "Code_Random_Thoughts/mypkg"
)

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	} else if root.Val > high {
		return trimBST(root.Left, low, high)
	} else if root.Val == low {
		root.Left = nil
		root.Right = trimBST(root.Right, low, high)
		return root
	} else if root.Val == high {
		root.Left = trimBST(root.Left, low, high)
		root.Right = nil
		return root
	} else {
		root.Left = trimBST(root.Left, low, high)
		root.Right = trimBST(root.Right, low, high)
		return root
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	lenN := len(nums)
	if lenN == 0 {
		return nil
	}
	res := &TreeNode{
		Val:   nums[lenN/2],
		Left:  nil,
		Right: nil,
	}
	res.Left = sortedArrayToBST(nums[0 : lenN/2])
	res.Right = sortedArrayToBST(nums[lenN/2+1:])
	return res
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	sum := 0
	var sumDFS func(root *TreeNode)
	sumDFS = func(root *TreeNode) {
		if root.Right != nil {
			sumDFS(root.Right)
		}
		sum += root.Val
		root.Val = sum
		if root.Left != nil {
			sumDFS(root.Left)
		}
	}
	sumDFS(root)
	return root
}
