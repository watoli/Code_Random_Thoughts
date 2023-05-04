# Day16 二叉树part03
##  104.二叉树的最大深度 
1. 递归
```go
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
```
## 111.二叉树的最小深度 
1. 递归
```go
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
```

## 222.完全二叉树的节点个数
1. 递归
```go
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
```