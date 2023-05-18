# Day30回溯算法part06

## 332.重新安排行程--回溯

这个题如果能想到如何标记机场是否已经去过，以及通过字典序排序候选机场，就比较容易使用模板做出来了。   
注意递归有无解情况时一定要使用返回值进行判断。

```go
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
```


## 51. N皇后--回溯+判断优化

通过设置行列与斜线是否为空，将检查行列或斜线是否有皇后的过程由O(n)优化为O(1)，优化的标记皇后代码如下：
```go
colum := make([]bool, n)
right := make([]bool, n*2)
left := make([]bool, n*2)
```
代码思想如下：

1. 创建一个布尔类型的数组`colum`，长度为N，用于表示每一列是否有皇后。初始时，所有元素均为`false`。

2. 创建两个布尔类型的数组`right`和`left`，长度分别为2N，用于表示主对角线和副对角线是否有皇后。初始时，所有元素均为`false`。

3. 创建一个二维字符串数组`res`，用于存储所有符合条件的皇后摆放方案。

4. 创建一个一维整数数组`tmpS`，用于存储当前正在处理的皇后位置。数组的索引代表行号，值代表列号。

5. 定义一个递归函数`BT`，参数`idx`表示当前正在处理的行号。

6. 在`BT`函数中，如果`idx`等于N，说明已经找到了一个完整的皇后摆放方案，将其转换为字符串表示，并将其添加到`res`中。

7. 在`BT`函数中，使用一个循环遍历当前行的所有列，尝试将皇后放置在每一列上。

8. 对于每个位置，通过检查`colum`数组、`right`数组和`left`数组，判断当前列、主对角线和副对角线是否已经有皇后。

9. 如果当前位置可以放置皇后，则执行以下操作：
   - 将当前列、主对角线和副对角线对应的数组元素标记为`true`，表示已经有皇后占据。
   - 将当前位置的列号添加到`tmpS`数组中，表示在当前行放置了一个皇后。
   - 递归调用`BT`函数处理下一行。
   - 回溯：将`tmpS`数组、`colum`数组、`right`数组和`left`数组恢复到之前的状态，以便尝试下一个位置。

10. 在主函数中，调用`BT`函数开始递归搜索，初始时处理第一行（行号为0）。

11. 返回存储所有符合条件的皇后摆放方案的`res`数组。

```go
func solveNQueens(n int) [][]string {
    // 创建辅助数组
    colum := make([]bool, n)          // 列是否有皇后的标记数组
    right := make([]bool, n*2)       // 主对角线是否有皇后的标记数组
    left := make([]bool, n*2)        // 副对角线是否有皇后的标记数组
    var res [][]string                // 存储结果的二维字符串数组
    var tmpS []int                    // 存储当前处理的皇后位置的一维整数数组

    // 定义回溯函数
    var BT func(idx int)
    BT = func(idx int) {
        // 如果处理完所有行，将当前皇后摆放方案转换为字符串形式，并添加到结果数组中
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
        
        // 遍历当前行的所有列，尝试将皇后放置在每一列上
        for i := 0; i < n; i++ {
            // 检查当前列、主对角线和副对角线是否已经有皇后
            if colum[i] == false && right[i+n-idx] == false && left[i+idx] == false {
                // 放置皇后
                tmpS = append(tmpS, i)
                colum[i] = true
                right[i+n-idx] = true
                left[i+idx] = true

                // 递归处理下一行
                BT(idx + 1)

                // 回溯：移除最后一个皇后，并将相关的数组元素标记恢复为false
                tmpS = tmpS[:len(tmpS)-1]
                colum[i] = false
                right[i+n-idx] = false
                left[i+idx] = false
            }
        }
    }

    // 调用回溯函数开始搜索，初始处理第一行（行号为0）
    BT(0)

    // 返回所有符合条件的皇后摆放方案
    return res
}
```

## 37. 解数独--回溯+约束搜索

题解上用的顺序遍历81个单元格进行回溯，但是当遇到数独中候选数字较多的情况时，回溯深度会非常大，因此设计了一个约束搜索方案尽可能降低回溯深度。

1. 首先观察数独盘面，找到已经给出的数字（已知数字）。

2. 创建一个候选数矩阵，记录每个空白格子可以填入的数字。初始时，所有空白格子的候选数都是1到9。

3. 从候选数矩阵中找到候选数最少的空白格子。也就是找到一个所在行、所在列和所在九宫格内候选数最少的格子。

4. 选择一个候选数填入该格子，并更新候选数矩阵。

5. 使用回溯的方法重复步骤3和步骤4，直到填满所有的空白格子。

这个算法还有改进的空间，也就是寻找存在最小候选数的格子时，目前是遍历了81个方格，
实际上可以通过堆排序或优先队列优化这个搜索，不过做完这个题脑袋已经迷迷糊糊了，留给后人去优化吧…

```go
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
```