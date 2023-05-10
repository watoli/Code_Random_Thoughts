## Day22二叉树
## 235. 二叉搜索树的最近公共祖先
1. 递归
```go
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
```
2. 答案比较简练地递归
```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
        return nil
    }
    for {
         if root.Val > p.Val && root.Val > q.Val {
            root = root.Left
    }
        if root.Val < p.Val && root.Val < q.Val {
            root = root.Right
        }
        if (root.Val - p.Val) * (root.Val - q.Val) <= 0 {
            return root
        }
    }
        return root
}
```

##  701.二叉搜索树中的插入操作
1. 愚蠢的从头遍历
```go
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{Val: val}
	if root == nil {
		return newNode
	}
	var pre *TreeNode
	var dfs func(root *TreeNode, val int)
	dfs = func(root *TreeNode, val int) {
		if root == nil {
			return
		}
		if root.Left != nil {
			dfs(root.Left, val)
		}
		if pre != nil {
			if val < pre.Val {
				pre.Left = newNode
				pre = newNode
				return
			} else if val > pre.Val && val < root.Val {
				//fmt.Println("pre:", pre.Val, " root:", root.Val)
				newNode.Left = root.Left
				//newNode.Right = root.Right
				root.Left = newNode
				pre = newNode
				return
			} else if pre.Val >= val {
				return
			}
		}
		pre = root
		if root.Right != nil {
			dfs(root.Right, val)
		}
	}
	dfs(root, val)
	if pre.Val != val {
		newNode.Left = root
		return newNode
	}
	return root
}
```
2. 大佬的机智简明递归，恐怖如斯
```go
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
```
## 450.删除二叉搜索树中的节点
1. 递归
```go
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
```