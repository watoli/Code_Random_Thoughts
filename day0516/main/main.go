package main

import (
	"fmt"
	"sort"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	var res []string
	if len(s) < 4 {
		return res
	}
	var temS []string
	var ifIP func(s string) bool
	ifIP = func(s string) bool {
		if len(s) != 1 && s[0] == '0' {
			return false
		}
		numS, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		if numS > 255 {
			return false
		}
		return true
	}
	var BT func(s string, idx int)
	BT = func(s string, idx int) {
		if len(temS) == 4 {
			ttS := ""
			for i := 0; i < len(temS); i++ {
				if i == 3 {
					ttS += temS[i]
				} else {
					ttS += temS[i] + "."
				}
			}
			res = append(res, ttS)
		}
		if (len(s)-idx) > (3*(4-len(temS))) || (len(s)-idx) < (4-len(temS)) {
			return
		}
		if len(temS) == 3 {
			if ifIP(s[idx:]) {
				temS = append(temS, s[idx:])
				BT(s, idx)
				temS = temS[:len(temS)-1]
			}
		} else {
			for i := idx; i < len(s) && i < idx+3; i++ {
				if ifIP(s[idx : i+1]) {
					temS = append(temS, s[idx:i+1])
					BT(s, i+1)
					temS = temS[:len(temS)-1]
				}
			}
		}
	}
	BT(s, 0)
	return res
}

/*
[1,2,3,4,5,6]
123 124 125
*/
func subsets(nums []int) [][]int {
	res := [][]int{{}}
	numsL := len(nums)
	var tmpS []int
	var BT func(nums []int, idx int)
	BT = func(nums []int, idx int) {
		for i := idx; i < numsL; i++ {
			tmpS = append(tmpS, nums[i])
			copyS := make([]int, len(tmpS))
			copy(copyS, tmpS)
			res = append(res, copyS)
			BT(nums, i+1)
			tmpS = tmpS[:len(tmpS)-1]
		}
	}
	BT(nums, 0)
	return res
}
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{{}}
	numsL := len(nums)
	var tmpS []int
	var BT func(nums []int, idx int)
	BT = func(nums []int, idx int) {
		for i := idx; i < numsL; {
			tmpS = append(tmpS, nums[i])
			copyS := make([]int, len(tmpS))
			copy(copyS, tmpS)
			res = append(res, copyS)
			BT(nums, i+1)
			tmpS = tmpS[:len(tmpS)-1]
			i++
			for i < numsL {
				if nums[i] == nums[i-1] {
					i++
					continue
				}
				break
			}
		}
	}
	BT(nums, 0)
	return res
}

func main() {
	res := [][]int{{}}
	fmt.Println(res)
}
