package main

import (
	"fmt"
	"sort"
	"strings"
)

type airportSlice struct {
	name []string
	flag []bool
	used bool
}

func findItinerary(tickets [][]string) []string {
	fmt.Println("\n", tickets)
	airportMap := make(map[string]*airportSlice)
	for i := 0; i < len(tickets); i++ {
		//fmt.Println(tickets[i][0])
		as, ok := airportMap[tickets[i][0]]
		if !ok {
			airS := &airportSlice{
				name: []string{tickets[i][1]},
				flag: []bool{false},
				used: false,
			}
			airportMap[tickets[i][0]] = airS
		} else {
			//fmt.Println(as)
			as.name = append(as.name, tickets[i][1])
			as.flag = append(as.flag, false)
		}
	}
	//fmt.Println(airportMap)
	for _, v := range airportMap {
		sort.Slice((*v).name, func(i, j int) bool {
			less := strings.Compare((*v).name[i], (*v).name[j])
			if less < 0 {
				return true
			} else {
				return false
			}
		})
	}
	for k, v := range airportMap {
		fmt.Println("k:", k, " v:", *v)
	}
	res := []string{"JFK"}
	lenT := len(tickets)
	var BT func(apS *airportSlice) bool
	BT = func(apS *airportSlice) bool {
		if len(res) == lenT+1 {
			return true
		} else if apS == nil || (*apS).used == true {
			return false
		}
		empty := true
		for i := 0; i < len((*apS).flag); i++ {
			if (*apS).flag[i] == true {
				continue
			} else {
				res = append(res, (*apS).name[i])
				(*apS).flag[i] = true
				ifOK := BT(airportMap[(*apS).name[i]])
				if ifOK == true {
					return true
				} else {
					res = res[:len(res)-1]
					(*apS).flag[i] = false
					empty = false
				}
			}
		}
		(*apS).used = empty
		return false
	}
	BT(airportMap["JFK"])
	return res
}

/*
0,0 0,0 6,4 0,0
0,0 4,4 0,0 6,6
0,0 0,0 0,0 0,0
0,0 0,0 0,0 0,0
*/

func solveNQueens(n int) [][]string {
	//row := make([]bool, n)
	colum := make([]bool, n)
	right := make([]bool, n*2)
	left := make([]bool, n*2)
	var res [][]string
	var tmpS []int
	var BT func(idx int)
	BT = func(idx int) {
		if idx == n {
			ttS := make([]string, 0)
			for i := 0; i < n; i++ {
				tmpStr := ""
				for j := 0; j < n; j++ {
					if tmpS[i] == j {
						tmpStr += "Q"
					} else {
						tmpStr += "."
					}
				}
				ttS = append(ttS, tmpStr)
			}
			res = append(res, ttS)
			return
		}
		for i := 0; i < n; i++ {
			if colum[i] == false && right[i+n-idx] == false && left[i+idx] == false {
				{
					tmpS = append(tmpS, i)
					colum[i] = true
					right[i+n-idx] = true
					left[i+idx] = true
				}
				BT(idx + 1)
				{
					tmpS = tmpS[:len(tmpS)-1]
					colum[i] = false
					right[i+n-idx] = false
					left[i+idx] = false
				}
			}
		}
	}
	BT(0)
	return res
}

//	type nine struct {
//		nType int //row0 colum1 block2
//		num   int
//		empty map[int]bool
//	}
//
//	func newNine() nine {
//		return nine{
//			nType: 0,
//			num:   0,
//			empty: map[int]bool{
//				1: true,
//				2: true,
//				3: true,
//				4: true,
//				5: true,
//				6: true,
//				7: true,
//				8: true,
//				9: true,
//			},
//		}
//	}
//
//	func solveSudoku(board [][]byte) {
//		nineS := make([]nine, 0)
//		for i := 0; i < 9; i++ {
//			tmpNine := newNine()
//			tmpNine.num = i
//			for j := 0; j < 9; j++ {
//				if board[i][j] != '.' {
//					delete(tmpNine.empty, int(board[i][j]))
//				}
//			}
//			nineS = append(nineS, tmpNine)
//		}
//		for i := 0; i < 9; i++ {
//			tmpNine := newNine()
//			tmpNine.nType = 1
//			tmpNine.num = i
//			for j := 0; j < 9; j++ {
//				if board[j][i] != '.' {
//					delete(tmpNine.empty, int(board[i][j]))
//				}
//			}
//			nineS = append(nineS, tmpNine)
//		}
//		fmt.Println(nineS)
//	}
func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	row, col := findNextEmptyCell(board)
	if row == -1 && col == -1 {
		return true // 数独已经填满，解题成功
	}

	candidates := getCandidates(board, row, col) // 获取候选数列表

	for _, num := range candidates {
		board[row][col] = num
		if solve(board) {
			return true
		}
		board[row][col] = '.' // 回溯
	}

	return false
}

func findNextEmptyCell(board [][]byte) (int, int) {
	minCandidates := 10
	nextRow, nextCol := -1, -1

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				candidates := getCandidates(board, row, col)
				if len(candidates) < minCandidates {
					minCandidates = len(candidates)
					nextRow, nextCol = row, col
				}
			}
		}
	}

	return nextRow, nextCol
}

func getCandidates(board [][]byte, row, col int) []byte {
	candidates := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

	// 检查所在行的数字，将已经出现的数字从候选数列表中删除
	for j := 0; j < 9; j++ {
		if board[row][j] != '.' {
			candidates = removeCandidate(candidates, board[row][j])
		}
	}

	// 检查所在列的数字，将已经出现的数字从候选数列表中删除
	for i := 0; i < 9; i++ {
		if board[i][col] != '.' {
			candidates = removeCandidate(candidates, board[i][col])
		}
	}

	// 检查所在九宫格的数字，将已经出现的数字从候选数列表中删除
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] != '.' {
				candidates = removeCandidate(candidates, board[startRow+i][startCol+j])
			}
		}
	}

	return candidates
}

func removeCandidate(candidates []byte, num byte) []byte {
	for i := 0; i < len(candidates); i++ {
		if candidates[i] == num {
			candidates[i] = candidates[len(candidates)-1]
			return candidates[:len(candidates)-1]
		}
	}
	return candidates
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(board)
}
