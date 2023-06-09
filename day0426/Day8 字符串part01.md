# Day8字符串part01

## 344. 反转字符串
1. 直接反转
```go
func reverseString(s []byte) {
	lenS := len(s)
	for i := 0; i < lenS/2; i++ {
		s[i], s[lenS-1-i] = s[lenS-1-i], s[i]
	}
}
```

## 541. 反转字符串 II
2. 还是直接反转，其中遍历是i+=2*k能够极大缩短时间，注意技巧
```go
func reverseStr(s string, k int) string {
	resByte := []byte(s)

	for i := 0; i < len(s); i = i + 2*k {

		if i+k <= len(s) {
			reverseString(resByte[i : i+k])
		} else if i+k > len(s) {
			reverseString(resByte[i:])
		}
	}
	return string(resByte)
}
```

## 剑指 Offer 05. 替换空格
1. 添加元素，空间复杂度为O(n)
```go
func replaceSpace(s string) (res string) {
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res += "%20"
		} else {
			res += fmt.Sprintf("%c", s[i])
		}
	}
	return
}
```
2. 双指针法
```go
func replaceSpace(s string) string {
	nSpace := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			nSpace++
		}
	}
	res := make([]byte, nSpace*2)
	res = append([]byte(s), res...)
	fast := len(s) - 1
	slow := len(res) - 1
	for fast >= 0 {
		if s[fast] == ' ' {
			res[slow] = '0'
			res[slow-1] = '2'
			res[slow-2] = '%'
			slow -= 3
			fast--
		} else {
			res[slow] = s[fast]
			slow--
			fast--
		}
	}
	return fmt.Sprintf("%s", res)
}
```

## 151. 反转字符串中的单词
1. 和反转字符串思路类似,需要优化删除空格的算法
```go
func reverseWords(s string) string {
	sliceS := make([][]byte, 0)
	tmpS := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			tmpS = append(tmpS, s[i])
		} else {
			if len(tmpS) != 0 {
				sliceS = append(sliceS, tmpS)
			}
			tmpS = []byte{}
		}
	}
	if len(tmpS) != 0 {
		sliceS = append(sliceS, tmpS)
	}
	for i := 0; i < len(sliceS)/2; i++ {
		sliceS[i], sliceS[len(sliceS)-1-i] = sliceS[len(sliceS)-1-i], sliceS[i]
	}
	res := fmt.Sprintf("%s", sliceS[0])
	for i := 1; i < len(sliceS); i++ {
		if len(sliceS[i]) != 0 {
			res += " " + fmt.Sprintf("%s", sliceS[i])
		}
	}
	return res
}
```

## 剑指 Offer 58 - II. 左旋转字符串
1. 字符串拼接
```go
func reverseLeftWords(s string, n int) string {
	res := make([]byte, len(s))
	lenS := len(s)
	for i := 0; i < n; i++ {
		res[lenS-n+i] = s[i]
	}
	for i := n; i < lenS; i++ {
		res[i-n] = s[i]
	}
	return fmt.Sprintf("%s", res)
}
```
2. 左旋
```go
func reverseLeftWords(s string, n int) string {
	res := []byte(s)
	reverseString(res[:n])
	reverseString(res[n:])
	reverseString(res)
	return fmt.Sprintf("%s", res)
}
```