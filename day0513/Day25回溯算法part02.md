# Day25回溯算法part02

##  216.组合总和III
这个题的回溯解法很容易想，但也有很有意思的动态规划解法，思想类似N数之和。
1. 回溯+双剪枝
```go
func combinationSum3(k int, n int) [][]int {
    res := make([][]int, 0)
    // 首先判断是否存在这样的组合，如果不存在则直接返回空结果集
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
    // tmpS表示当前已选择的数字集合，tmpSum表示已选择数字的和
    tmpS := make([]int, 0)
    tmpSum := 0
    // BT表示回溯函数，m表示当前可选数字的起点，k表示需要选择的数字个数，n表示数字和
    var BT func(m, k, n int)
    BT = func(m, k, n int) {
        // 如果已经选择了k个数字，且数字和为n，则将当前结果加入到结果集中
        if len(tmpS) == k {
            if tmpSum == n {
                copyS := make([]int, k)
                copy(copyS, tmpS)
                res = append(res, copyS)
            }
            return
        }
        // 计算当前可选数字的上界，这里使用了一个数学公式，可以让时间复杂度更优
        lim := (2*n - 2*tmpSum - (k-len(tmpS))*(k-len(tmpS)) + (k - len(tmpS))) / (2 * (k - len(tmpS)))
        if lim > 9 {
            lim = 9
        }
        // 从m开始遍历可选数字，选一个数字加入到tmpS中，并递归调用BT
        for i := m; i <= lim; i++ {
            tmpS = append(tmpS, i)
            tmpSum += i
            BT(i+1, k, n)
            tmpS = tmpS[:len(tmpS)-1] // 回溯时需要将tmpS和tmpSum恢复到上一个状态
            tmpSum -= i
        }
    }
    BT(1, k, n)
    return res
}

```
2. 动态规划

这个动态规划解法是用来解决组合总和问题的，找出所有相加之和为 n 的 k 个数的组合。其中每个数字只能使用一次，且只能使用数字 1 到 9。

首先，我们定义一个二维布尔数组 dp，其中 dp[i][j] 表示是否存在 k 个数的和为 j，其中所有数字都在 1 到 i 之间。因为我们只需要知道是否存在一种组合，而不需要知道具体的组合，所以数组中不需要存储具体的数字。

接下来，我们可以初始化 dp 数组。当 k 为 0 且 n 为 0 时，可以设置 dp[0][0] 为 true。当 k 不为 0 但 n 为 0 时，可以设置 dp[k][0] 为 true，因为 0 也可以被表示为 0 个数字的和。当 k 为 0 但 n 不为 0 时，可以设置 dp[0][n] 为 false，因为无法使用数字来表示非零数。

然后，我们可以使用一个三重循环来填充 dp 数组。外层循环枚举数字 1 到 9，中层循环枚举 k 到 1，内层循环枚举 n 到 i。对于每个 dp[j][s]，我们可以使用 dp[j][s] = dp[j][s] || dp[j-1][s-i] 来更新它，其中 dp[j][s-i] 表示是否存在 j-1 个数字的和为 s-i，而 i 表示当前数字。

最后，我们可以根据 dp 数组来构造结果。如果 dp[k][n] 为 true，则表示存在 k 个数字的和为 n，我们可以使用 DFS 来查找所有可能的组合。我们从数字 1 开始，对于每个数字，我们可以选择不将它添加到组合中，或者将它添加到组合中。如果添加到组合中，则需要更新 sum 的值，直到我们找到 k 个数字的和为 n。如果当前的数字超过了 9，或者当前的 sum 超过了 n，或者当前的组合长度已经达到了 k，则可以停止搜索。

```go
func combinationSum3(k int, n int) [][]int {
    // 初始化一个二维bool数组dp，表示用k个数能否组成和为i
    dp := make([][]bool, k+1)
    for i := range dp {
        dp[i] = make([]bool, n+1)
    }
    // 当不选取任何数时，和为0，所以dp[0][0]=true
    dp[0][0] = true

    // 枚举数字1~9
    for i := 1; i <= 9; i++ {
        // 从k到1逆序枚举，保证每个数字只选一次
        for j := k; j >= 1; j-- {
            // 枚举和i到n，注意不能从0开始，否则会出现重复组合
            for s := n; s >= i; s-- {
                // 两种情况：不选i和选i
                // 如果不选i，dp[j][s]的值就不变，等于上一轮的结果
                // 如果选i，dp[j-1][s-i]表示用j-1个数组成s-i的情况，再加上i就能组成s
                dp[j][s] = dp[j][s] || dp[j-1][s-i]
            }
        }
    }

    // 如果能够用k个数组成和为n，就把所有可能的组合添加到结果res中
    res := make([][]int, 0)
    if dp[k][n] {
        tmp := make([]int, 0)
        // 定义一个dfs函数，用于回溯查找所有可能的组合
        var dfs func(int, int)
        dfs = func(cur, sum int) {
            // 如果找到一组长度为k且和为n的组合，就把该组合添加到结果res中
            if len(tmp) == k && sum == n {
                tmpCopy := make([]int, len(tmp))
                copy(tmpCopy, tmp)
                res = append(res, tmpCopy)
                return
            }
            // 如果当前数字大于9，或当前和大于n，或当前组合的长度已经达到k，就结束递归
            if cur > 9 || sum > n || len(tmp) >= k {
                return
            }
            // 两种情况：不选当前数字和选当前数字
            // 不选当前数字，直接递归到下一个数字
            dfs(cur+1, sum)
            // 选当前数字，把当前数字添加到组合中，递归到下一个数字
            tmp = append(tmp, cur)
            dfs(cur+1, sum+cur)
            tmp = tmp[:len(tmp)-1]
        }
        dfs(1, 0)
    }

    return res
}
```

##  17.电话号码的字母组合 
1. 有趣的ASCII解法
```go
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
```