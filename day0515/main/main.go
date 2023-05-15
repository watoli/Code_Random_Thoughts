package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	//fmt.Println(candidates)
	var res [][]int
	if candidates[0] > target {
		return res
	}
	var tmpS []int
	var BT func(candidates []int, target, idx, sum int)
	BT = func(candidates []int, target, idx, sum int) {
		if sum == target {
			copyS := make([]int, len(tmpS))
			copy(copyS, tmpS)
			res = append(res, copyS)
		} else {
			if sum > target {
				return
			}
		}
		for i := idx; i < len(candidates) && (sum+candidates[i]) <= target; i++ {
			tmpS = append(tmpS, candidates[i])
			sum += candidates[i]
			BT(candidates, target, i, sum)
			tmpS = tmpS[:len(tmpS)-1]
			sum -= candidates[i]
		}

	}
	BT(candidates, target, 0, 0)
	return res
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	//fmt.Println(candidates)
	var res [][]int
	if candidates[0] > target {
		return res
	}
	var tmpS []int
	//limS := make([]int, len(candidates))
	//limS[len(candidates)-1] = candidates[len(candidates)-1]
	//for i := len(candidates) - 2; i >= 0; i-- {
	//	limS[i] = candidates[i] + limS[i+1]
	//}
	var BT func(candidates []int, target, idx, sum int)
	BT = func(candidates []int, target, idx, sum int) {
		fmt.Println("")
		fmt.Println("tmpS:", tmpS, " sum:", sum)
		if sum == target {
			fmt.Println("append:", tmpS)
			copyS := make([]int, len(tmpS))
			copy(copyS, tmpS)
			res = append(res, copyS)
			return
		} else {
			if sum > target {
				return
			}
		}
		for i := idx; i < len(candidates) && (sum+candidates[i]) <= target; {
			fmt.Println("i:", i, " sum:", sum)
			if i == (len(candidates) - 1) {
				if (target-sum)%(candidates[len(candidates)-1]) != 0 {
					break
				}
			}
			tmpS = append(tmpS, candidates[i])
			sum += candidates[i]
			BT(candidates, target, i+1, sum)
			tmpS = tmpS[:len(tmpS)-1]
			sum -= candidates[i]
			i++
			for i < len(candidates) {
				fmt.Println(i)
				if candidates[i] == candidates[i-1] {
					i++
					continue
				}
				break
			}
		}

	}
	BT(candidates, target, 0, 0)
	return res
}

func partition(s string) [][]string {
	var ifPalindrome func(subS string) bool
	ifPalindrome = func(subS string) bool {
		lS := len(subS)
		for i := 0; i < lS/2; i++ {
			if subS[i] != subS[lS-i-1] {
				return false
			}
		}
		return true
	}
	lenS := len(s)
	var res [][]string
	var tmpS []string
	tmpL := 0
	var BT func(s string, idx int)
	BT = func(s string, idx int) {
		fmt.Printf("lenS: %d, res: %v, tmpS: %v, tmpL: %d, idx: %d\n", lenS, res, tmpS, tmpL, idx)
		if tmpL == lenS {
			copyS := make([]string, len(tmpS))
			copy(copyS, tmpS)
			res = append(res, copyS)
			return
		}
		if idx <= lenS-1 {
			tmpS = append(tmpS, string(s[idx]))
			tmpL++
			BT(s, idx+1)
			tmpS = tmpS[:len(tmpS)-1]
			tmpL--
		}
		for i := idx + 2; i <= lenS; i++ {
			ttS := s[idx:i]
			if ifPalindrome(ttS) {
				tmpS = append(tmpS, ttS)
				tmpL += i - idx
				BT(s, i)
				tmpS = tmpS[:len(tmpS)-1]
				tmpL -= i - idx
			}
		}
	}
	BT(s, 0)
	return res
}
