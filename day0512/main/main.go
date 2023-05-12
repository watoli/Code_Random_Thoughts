package main

import "fmt"

/*
[0,1,2,3,4,5,6,7,8]
[1,3,] k=5 l=2
4~8 6 k-l=3  n-3=5+1
*/
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
