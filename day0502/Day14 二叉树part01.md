# Day14 二叉树part01
## 二叉树遍历
```go
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{root.Val}
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

func preorderTraversalIteration(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	nodeStack := []*TreeNode{root}
	for len(nodeStack) != 0 {
		tmpNode := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		if tmpNode == nil {
			continue
		}
		res = append(res, tmpNode.Val)
		nodeStack = append(nodeStack, tmpNode.Right, tmpNode.Left)
	}
	return res
}

func preorderTraversalIterationUnify(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	nodeStack := []*TreeNode{root}
	for len(nodeStack) != 0 {
		tmpNode := nodeStack[len(nodeStack)-1]
		if tmpNode != nil {
			nodeStack = nodeStack[:len(nodeStack)-1]
			if tmpNode.Right != nil {
				nodeStack = append(nodeStack, tmpNode.Right)
			}
			if tmpNode.Left != nil {
				nodeStack = append(nodeStack, tmpNode.Left)
			}
			nodeStack = append(nodeStack, tmpNode, nil)
		} else {
			nodeStack = nodeStack[:len(nodeStack)-1]
			tmpNode = nodeStack[len(nodeStack)-1]
			nodeStack = nodeStack[:len(nodeStack)-1]
			res = append(res, tmpNode.Val)
		}
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

func inorderTraversalIteration(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var nodeStack []*TreeNode
	curNode := root
	for curNode != nil || len(nodeStack) != 0 {
		if curNode != nil {
			nodeStack = append(nodeStack, curNode)
			curNode = curNode.Left
		} else {
			curNode = nodeStack[len(nodeStack)-1]
			nodeStack = nodeStack[:len(nodeStack)-1]
			res = append(res, curNode.Val)
			curNode = curNode.Right
		}
	}
	return res
}

func inorderTraversalIterationUnify(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	nodeStack := []*TreeNode{root}
	for len(nodeStack) != 0 {
		tmpNode := nodeStack[len(nodeStack)-1]
		if tmpNode != nil {
			nodeStack = nodeStack[:len(nodeStack)-1]
			if tmpNode.Right != nil {
				nodeStack = append(nodeStack, tmpNode.Right)
			}
			nodeStack = append(nodeStack, tmpNode, nil)
			if tmpNode.Left != nil {
				nodeStack = append(nodeStack, tmpNode.Left)
			}
		} else {
			nodeStack = nodeStack[:len(nodeStack)-1]
			tmpNode = nodeStack[len(nodeStack)-1]
			nodeStack = nodeStack[:len(nodeStack)-1]
			res = append(res, tmpNode.Val)
		}
	}
	return res
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

func postorderTraversalIteration(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	nodeStack := []*TreeNode{root}
	for len(nodeStack) != 0 {
		tmpNode := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		if tmpNode == nil {
			continue
		}
		res = append([]int{tmpNode.Val}, res...)
		nodeStack = append(nodeStack, tmpNode.Left, tmpNode.Right)
	}
	return res
}

func postorderTraversalIterationUnify(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	nodeStack := []*TreeNode{root}
	for len(nodeStack) != 0 {
		tmpNode := nodeStack[len(nodeStack)-1]
		if tmpNode != nil {
			nodeStack = nodeStack[:len(nodeStack)-1]
			nodeStack = append(nodeStack, tmpNode, nil)
			if tmpNode.Right != nil {
				nodeStack = append(nodeStack, tmpNode.Right)
			}
			if tmpNode.Left != nil {
				nodeStack = append(nodeStack, tmpNode.Left)
			}
		} else {
			nodeStack = nodeStack[:len(nodeStack)-1]
			tmpNode = nodeStack[len(nodeStack)-1]
			nodeStack = nodeStack[:len(nodeStack)-1]
			res = append(res, tmpNode.Val)
		}
	}
	return res
}

```