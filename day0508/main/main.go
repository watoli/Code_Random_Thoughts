package main

import (
	. "Code_Random_Thoughts/mypkg"
	"container/ring"
	"fmt"
)

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	max := -1
	loc := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			loc = i
		}
	}
	root.Val = max
	root.Left = constructMaximumBinaryTree(nums[:loc])
	root.Right = constructMaximumBinaryTree(nums[loc+1:])
	return root
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 != nil && root2 != nil {
		root := &TreeNode{
			Val:   root1.Val + root2.Val,
			Left:  mergeTrees(root1.Left, root2.Left),
			Right: mergeTrees(root1.Right, root2.Right),
		}
		return root
	} else if root1 == nil {
		return root2
	} else {
		return root1
	}
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Left != nil {
		res := searchBST(root.Left, val)
		if res != nil {
			return res
		}
	}
	if root.Right != nil {
		res := searchBST(root.Right, val)
		if res != nil {
			return res
		}
	}
	return nil
}

func BST(pre, cur **TreeNode) bool {
	if *cur == nil {
		return true
	}
	left := BST(pre, &(*cur).Left)
	if (*pre) != nil && (*cur).Val <= (*pre).Val {
		return false
	}
	*pre = *cur
	right := BST(pre, &(*cur).Right)
	return left && right
}

func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	if root == nil {
		return true
	}
	return BST(&pre, &root)
}

/*
[5,5,2,1,0]
*/
func jump(nums []int) int {
	numsL := len(nums)
	dp := make([]int, numsL)
	for i := 0; i < numsL; i++ {
		dp[i] = numsL
	}
	dp[numsL-1] = 0
	for i := numsL - 2; i >= 0; i-- {
		min := numsL
		for j := 0; j <= nums[i] && i+j <= (numsL-1); j++ {
			if 1+dp[i+j] < min {
				min = 1 + dp[i+j]
			}
		}
		dp[i] = min
	}
	return dp[0]
}

func main() {
	// 创建两个容量为3的Ring
	r1 := ring.New(3)
	for i := 1; i <= r1.Len(); i++ {

		r1.Value = i
		fmt.Println(r1)
		r1 = r1.Next()
	}

	r2 := ring.New(3)
	for i := 4; i <= r2.Len()+3; i++ {
		fmt.Println(r2)
		r2.Value = i
		r2 = r2.Next()
	}

	// 输出原始的Ring
	fmt.Println("Original rings:")
	printRing(r1)
	printRing(r2)

	// 将r2链接到r1的后面
	r1.Link(r2)
	fmt.Println("After linking r2 to r1:")
	printRing(r1)
}

// 打印Ring中的所有元素
func printRing(r *ring.Ring) {
	r.Do(func(x interface{}) {
		fmt.Print(x, " ")
	})
	fmt.Println()
}
