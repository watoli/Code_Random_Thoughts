package main

import (
	. "Code_Random_Thoughts/mypkg"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode
	var bigger *TreeNode
	var smaller *TreeNode
	if p.Val > q.Val {
		bigger = p
		smaller = q
	} else {
		bigger = q
		smaller = p
	}

	var dfs func(root, bigger, smaller *TreeNode)
	dfs = func(root, bigger, smaller *TreeNode) {
		if root.Val > bigger.Val {
			dfs(root.Left, bigger, smaller)
		} else if root.Val < smaller.Val {
			dfs(root.Right, bigger, smaller)
		} else {
			res = root
			return
		}
	}
	dfs(root, bigger, smaller)
	return res
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func getRight(root *TreeNode) *TreeNode {
	if root == nil || root.Right == nil {
		return root
	}
	return getRight(root.Right)
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		} else {
			left := getRight(root.Left)
			left.Right = root.Right
		}
		root = root.Left
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}
