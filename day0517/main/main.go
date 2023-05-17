package main

import (
	"fmt"
	"sort"
)

func findSubsequences(nums []int) [][]int {
	var res [][]int
	numsL := len(nums)
	var tmpS []int
	var BT func(nums []int, idx int)

	BT = func(nums []int, idx int) {
		usedS := make([]int, 201)
		for i := idx; i < numsL; i++ {
			if len(tmpS) > 0 && nums[i] < tmpS[len(tmpS)-1] || usedS[nums[i]+100] == 1 {
				continue
			}
			tmpS = append(tmpS, nums[i])
			if len(tmpS) > 1 {
				copyS := make([]int, len(tmpS))
				copy(copyS, tmpS)
				res = append(res, copyS)
			}
			BT(nums, i+1)
			tmpS = tmpS[:len(tmpS)-1]
			usedS[nums[i]+100] = 1
		}
	}
	BT(nums, 0)
	return res
}

/*
[1,2,3,4,5,6]
*/

func permute(nums []int) [][]int {
	var res [][]int
	var tmpS []int
	lenN := len(nums)
	usedS := make([]int, lenN)
	var BT func(nums []int)
	BT = func(nums []int) {
		if len(tmpS) == lenN {
			copyS := make([]int, len(tmpS))
			copy(copyS, tmpS)
			res = append(res, copyS)
			return
		}
		for i := 0; i < lenN; i++ {
			if usedS[i] == 1 {
				continue
			}
			tmpS = append(tmpS, nums[i])
			usedS[i] = 1
			BT(nums)
			usedS[i] = 0
			tmpS = tmpS[:len(tmpS)-1]
		}
	}
	BT(nums)
	return res
}

/*
[1,1,1,2,2,2]
*/
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var tmpS []int
	lenN := len(nums)
	usedS := make([]int, lenN)
	var BT func(nums []int)
	BT = func(nums []int) {
		fmt.Println("tmpS:", tmpS, " usedS: ", usedS)
		if len(tmpS) == lenN {
			copyS := make([]int, len(tmpS))
			copy(copyS, tmpS)
			res = append(res, copyS)
			return
		}
		for i := 0; i < lenN; {
			fmt.Println(i)
			if usedS[i] == 1 {
				i++
				continue
			}
			tmpS = append(tmpS, nums[i])
			usedS[i] = 1
			BT(nums)
			usedS[i] = 0
			tmpS = tmpS[:len(tmpS)-1]
			i++
			for i < lenN {
				fmt.Println(i)
				if nums[i] == nums[i-1] {
					i++
				} else {
					break
				}
			}
		}
	}
	BT(nums)
	return res
}
