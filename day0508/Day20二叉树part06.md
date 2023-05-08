# Day20二叉树part06

## 654.最大二叉树 
1. 递归
```go
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
```
##  617.合并二叉树 
1. 递归
```go
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
```
## 700.二叉搜索树中的搜索
1. 递归
```go
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
```
##  98.验证二叉搜索树
1. 迭代
```go
func isValidBST(root *TreeNode) bool {
	nodeStack := list.New()
	var resS []int
	nodeStack.PushBack(root)
	for nodeStack.Len() > 0 {
		tmpN := nodeStack.Back()
		nodeStack.Remove(tmpN)
		if tmpN.Value == nil {
			tmpN = nodeStack.Back()
			nodeStack.Remove(tmpN)
			resS = append(resS, tmpN.Value.(*TreeNode).Val)
			continue
		}
		if tmpN.Value.(*TreeNode).Right != nil {
			nodeStack.PushBack(tmpN.Value.(*TreeNode).Right)
		}
		nodeStack.PushBack(tmpN.Value.(*TreeNode))
		nodeStack.PushBack(nil)
		if tmpN.Value.(*TreeNode).Left != nil {
			nodeStack.PushBack(tmpN.Value.(*TreeNode).Left)
		}
	}
	for i := 0; i < len(resS)-1; i++ {
		if resS[i] >= resS[i+1] {
			return false
		}
	}
	return true
}
```
2. 递归
```go
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
```