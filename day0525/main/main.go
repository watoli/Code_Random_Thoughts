package main

import (
	. "Code_Random_Thoughts/mypkg"
	"fmt"
	"strconv"
)

/*
2335
*/
func monotoneIncreasingDigits(n int) int {
	if n < 10 {
		return n
	}
	numS := make([]int, 0)
	for n > 0 {
		numS = append([]int{n % 10}, numS...)
		n /= 10
	}

	idx := len(numS)
	for i := len(numS) - 2; i >= 0; i-- {
		if numS[i] > numS[i+1] {
			numS[i] -= 1
			for j := i + 1; j < idx; j++ {
				numS[j] = 9
			}
			idx = i + 1
		}
	}
	res := 0
	for i := 0; i < len(numS); i++ {
		res = res*10 + numS[i]
	}
	return res
}

func minCameraCover(root *TreeNode) int {
	res := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 2
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if left == 0 || right == 0 {
			res++
			return 1
		} else if left == 2 && right == 2 {
			return 0
		} else {
			return 2
		}
	}
	rootRes := dfs(root)
	if rootRes == 0 {
		return res + 1
	}
	return res

}

func main() {
	num := 1234567890
	fmt.Println(num)
	strNum := strconv.Itoa(num) // 将整数转换为字符串
	chars := []byte(strNum)     // 将字符串转换为字符切片
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	reverseStrNum := string(chars) // 反转字符切片并转换为字符串
	for _, ch := range reverseStrNum {
		fmt.Printf("%c\n", ch) // 逐个字符地输出反转后的字符串
	}
}
