package main

import (
	"fmt"
)

/*
	k=4,n=15
	[1,2,]
	6,6
	[x+l-1+x]l/2<n-ts
	2xl+l^2-l<2n-2ts
	x=(2n-2ts-l^2+l)2l
	14-9+3

8/2/3
*/

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	if ifExist := func(k, n int) bool {
		min := 0
		max := 0
		for i := 1; i <= k; i++ {
			min += i
			max += 10 - i
		}
		if min > n || max < n {
			return false
		}
		return true
	}(k, n); !ifExist {
		return res
	}
	tmpS := make([]int, 0)
	tmpSum := 0
	var BT func(m, k, n int)
	BT = func(m, k, n int) {
		//fmt.Println(" m:", m, " k:", k, " n:", n, " tmpSum", tmpSum)
		//fmt.Println(tmpS)
		if len(tmpS) == k {
			if tmpSum == n {
				copyS := make([]int, k)
				copy(copyS, tmpS)
				res = append(res, copyS)
			}
			return
		}

		lim := (2*n - 2*tmpSum - (k-len(tmpS))*(k-len(tmpS)) + (k - len(tmpS))) / (2 * (k - len(tmpS)))
		if lim > 9 {
			lim = 9
		}
		//fmt.Println(" lim:", lim, "  (k - len(tmpS)):", k-len(tmpS))
		for i := m; i <= lim; i++ {
			tmpS = append(tmpS, i)
			tmpSum += i
			BT(i+1, k, n)
			tmpS = tmpS[:len(tmpS)-1]
			tmpSum -= i
		}

	}
	BT(1, k, n)
	return res
}

func letterCombinations(digits string) []string {
	var res []string
	lenD := len(digits)
	if lenD == 0 {
		return res
	}
	var tmpS []byte
	var BT func(digits string, idx int)
	BT = func(digits string, idx int) {
		fmt.Println(int(digits[idx]))
		if idx == lenD {
			res = append(res, string(tmpS))
			return
		}
		switch digits[idx] {
		case '7':
			for i := 0; i < 4; i++ {
				tmpS = append(tmpS, byte(int('p')+i))
				BT(digits, idx+1)
				tmpS = tmpS[:len(tmpS)-1]
			}
		case '8':
			for i := 0; i < 3; i++ {
				tmpS = append(tmpS, byte(int('t')+i))
				BT(digits, idx+1)
				tmpS = tmpS[:len(tmpS)-1]
			}
		case '9':
			for i := 0; i < 4; i++ {
				tmpS = append(tmpS, byte(int('w')+i))
				BT(digits, idx+1)
				tmpS = tmpS[:len(tmpS)-1]
			}
		default:
			for i := 0; i < 3; i++ {
				tmpS = append(tmpS, byte((int(digits[idx])-50)*3+97+i))
				BT(digits, idx+1)
				tmpS = tmpS[:len(tmpS)-1]
			}
		}
	}
	BT(digits, 0)
	return res
}

func main() {
	sInt := "0123456789"
	for i := 0; i < len(sInt); i++ {
		fmt.Println(i, ":", sInt[i])
	}
}
