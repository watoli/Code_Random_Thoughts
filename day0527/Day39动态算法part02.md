# Day39动态算法part02

##  62.不同路径
```go
func uniquePaths(m int, n int) int {
	//if m == 1 && n == 1 {
	//	return 1
	//}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		tmpS := make([]int, n)
		dp[i] = tmpS
		//dp[i] = append(dp[i], tmpS...)
	}
	dp[m-1][n-1] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				continue
			}
			if i == m-1 {
				dp[i][j] += dp[i][j+1]
			} else if j == n-1 {
				dp[i][j] += dp[i+1][j]
			} else {
				dp[i][j] += dp[i+1][j] + dp[i][j+1]
			}
		}
	}
	return dp[0][0]
}
```

## 63. 不同路径 II
```go
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = 1
	}
	tmpM := 0
	tmpN := 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			tmpM, tmpN = 0, 0
			fmt.Println("i:", i, " j:", j, "\n", dp)
			if obstacleGrid[i-1][j] != 1 {
				tmpM = dp[i-1][j]
			}
			if obstacleGrid[i][j-1] != 1 {
				tmpN = dp[i][j-1]
			}
			dp[i][j] = tmpM + tmpN
		}
	}
	return dp[m-1][n-1]
}
```