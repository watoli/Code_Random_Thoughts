package main

import (
	. "Code_Random_Thoughts/mypkg"
	"fmt"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftD := maxDepth(root.Left)
	rightD := maxDepth(root.Right)
	if leftD > rightD {
		return 1 + leftD
	} else {
		return 1 + rightD
	}
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftD := minDepth(root.Left)
	rightD := minDepth(root.Right)
	if leftD == 0 && rightD == 0 {
		return 1
	} else if leftD == 0 {
		return 1 + rightD
	} else if rightD == 0 {
		return 1 + leftD
	} else if rightD < leftD {
		return 1 + rightD
	} else {
		return 1 + leftD
	}
}

func getMin(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getMin(root.Left) + 1
}

func getMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getMax(root.Right) + 1
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min := getMin(root)
	max := getMax(root)
	fmt.Println(min, max)
	if max == min {
		return 1<<min - 1
	} else {
		return countNodes(root.Left) + countNodes(root.Right) + 1
	}
}

func main() {
	fmt.Println((1 << 1) - 1)
}
