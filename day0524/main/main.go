package main

import (
	"fmt"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := 0
	end := intervals[0][1]
	fmt.Println(intervals)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			end = min(end, intervals[i][1])
			res++
		} else {
			end = intervals[i][1]
		}
	}
	return res
}

func partitionLabels(s string) []int {
	alS := [26]int{}
	for i := 0; i < len(s); i++ {
		alS[int(s[i]-'a')] = i
	}
	left, right := 0, 0
	res := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if alS[int(s[i]-'a')] > right {
			right = alS[int(s[i]-'a')]
		}
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}
	return res
}
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[i-1][1] {
			intervals[i][0] = intervals[i-1][0]
			if intervals[i-1][1] > intervals[i][1] {
				intervals[i][1] = intervals[i-1][1]
			}
		} else {
			res = append(res, intervals[i-1])
		}
	}
	res = append(res, intervals[len(intervals)-1])
	return res
}
