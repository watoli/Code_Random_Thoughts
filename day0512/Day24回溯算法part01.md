# Day24回溯算法part01

##  77. 组合  
1. 回溯
```go
func combine(n int, k int) [][]int {
	var res [][]int
	var tmpS []int
	var dfs func(m, n, k int, s *[]int)
	dfs = func(m, n, k int, s *[]int) {
		if len(*s) == k {
			fmt.Println(res)
			tmp := make([]int, len(*s))
			copy(tmp, *s)
			res = append(res, tmp)
			return
		}
		for i := m; i <= n-(k-len(*s))+1; i++ {
			fmt.Println("i:", i, " m:", m, " n:", n, " s:", *s)
			*s = append(*s, i)
			dfs(i+1, n, k, s)
			*s = (*s)[:len(*s)-1]
		}
	}
	dfs(1, n, k, &tmpS)
	return res
}
```
2. 不用指针或全局变量
```go
func combine(n int, k int) [][]int {
	var res [][]int
	var tmpS []int
	var dfs func(m, n, k int, s []int)
	dfs = func(m, n, k int, s []int) {
		if len(s) == k {
			fmt.Println(res)
			tmp := make([]int, len(s))
			copy(tmp, s)
			res = append(res, tmp)
			return
		}
		for i := m; i <= n-(k-len(s))+1; i++ {
			fmt.Println("i:", i, " m:", m, " n:", n, " s:", s)
			s = append(s, i)
			dfs(i+1, n, k, s)
			s = (s)[:len(s)-1]
		}
	}
	dfs(1, n, k, tmpS)
	return res
}
```