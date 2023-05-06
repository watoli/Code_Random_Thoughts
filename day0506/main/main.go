package main

import (
	. "Code_Random_Thoughts/mypkg"
)

func findBottomLeftValue(root *TreeNode) int {
	nodeStack := []*TreeNode{root}
	tmpL := 0
	tmpN := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	for len(nodeStack) != 0 {
		tmpL = len(nodeStack)
		head := nodeStack[0]
		for i := 0; i < tmpL; i++ {
			tmpN = nodeStack[0]
			nodeStack = nodeStack[1:]
			if tmpN.Left != nil {
				nodeStack = append(nodeStack, tmpN.Left)
			}
			if tmpN.Right != nil {
				nodeStack = append(nodeStack, tmpN.Right)
			}
		}
		if len(nodeStack) == 0 {
			return head.Val
		}
	}
	return 0
}

func dfs(root *TreeNode, targetSum, curSum int) bool {
	if root == nil {
		return false
	}
	curSum = root.Val + curSum
	if root.Left == nil && root.Right == nil {
		if curSum != targetSum {
			return false
		} else {
			return true
		}
	}
	return dfs(root.Left, targetSum, curSum) || dfs(root.Right, targetSum, curSum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs(root, targetSum, 0)
}

func dfsII(root *TreeNode, targetSum, curSum int, curS *[]int, curSS *[][]int) {
	curSum = curSum + root.Val
	*curS = append(*curS, root.Val)
	if root.Left == nil && root.Right == nil {
		if curSum == targetSum {
			copyRes := make([]int, len(*curS))
			copy(copyRes, *curS)
			*curSS = append(*curSS, copyRes)
		} else {
			return
		}
	} else {
		if root.Left != nil {
			dfsII(root.Left, targetSum, curSum, curS, curSS)
			*curS = (*curS)[:len(*curS)-1]
		}
		if root.Right != nil {
			dfsII(root.Right, targetSum, curSum, curS, curSS)
			*curS = (*curS)[:len(*curS)-1]
		}
	}

}

func pathSum(root *TreeNode, targetSum int) [][]int {

	curSS := make([][]int, 0)
	curS := make([]int, 0)
	if root == nil {
		return curSS
	}
	dfsII(root, targetSum, 0, &curS, &curSS)
	return curSS
}

func buildTreePI(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	mid := preorder[0]
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == mid {
			root.Left = buildTreePI(preorder[1:i+1], inorder[:i])
			root.Right = buildTreePI(preorder[i+1:], inorder[i+1:])
		}
	}
	return root
}

func buildTreeIP(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  nil,
		Right: nil,
	}
	mid := postorder[len(postorder)-1]
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == mid {
			root.Left = buildTreePI(inorder[:i], postorder[:i])
			root.Right = buildTreePI(inorder[i+1:], postorder[i:len(postorder)-1])
		}
	}
	return root
}
