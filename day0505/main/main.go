package main

import (
	. "Code_Random_Thoughts/mypkg"
	"strconv"
)

func BalancedDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		leftD := BalancedDepth(root.Left)
		rightD := BalancedDepth(root.Right)
		if leftD == -1 || rightD == -1 {
			return -1
		}
		if leftD > rightD {
			if leftD-rightD > 1 {
				return -1
			}
			return 1 + leftD
		} else {
			if rightD-leftD > 1 {
				return -1
			}
			return 1 + rightD
		}
	}
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		res := BalancedDepth(root)
		if res == -1 {
			return false
		} else {
			return true
		}
	}
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		return append(res, strconv.Itoa(root.Val))
	}
	if root.Left != nil {
		Left := binaryTreePaths(root.Left)
		for i := 0; i < len(Left); i++ {
			Left[i] = strconv.Itoa(root.Val) + "->" + Left[i]
		}
		res = append(res, Left...)
	}
	if root.Right != nil {
		Right := binaryTreePaths(root.Right)
		for i := 0; i < len(Right); i++ {
			Right[i] = strconv.Itoa(root.Val) + "->" + Right[i]
		}
		res = append(res, Right...)
	}
	return res
}

func getSum(root *TreeNode, isLeft bool) int {
	if root.Left == nil && root.Right == nil {
		if isLeft == true {
			return root.Val
		} else {
			return 0
		}
	} else {
		sum := 0
		if root.Left != nil {
			sum += getSum(root.Left, true)
		}
		if root.Right != nil {
			sum += getSum(root.Right, false)
		}
		return sum
	}
}

func sumOfLeftLeaves(root *TreeNode) int {
	return getSum(root, false)
}
