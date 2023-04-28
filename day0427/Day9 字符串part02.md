# Day9字符串part02
## 28. 找出字符串中第一个匹配项的下标
1. 滑动窗口
```go
func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	lenN := len(needle)
	lenH := len(haystack)
	slow := 0
	fast := lenN - 1
	for fast != lenH {
		if haystack[fast] == needle[lenN-1] {
			isNeedle := true
			for i := 0; i < lenN; i++ {
				if haystack[slow+i] != needle[i] {
					isNeedle = false
				}
			}
			if isNeedle == true {
				return slow
			}
		}
		fast++
		slow++
	}
	return -1
}
```
2. KMP
```go
func getNext(next []int, s string) {
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	} else if len(needle) == 0 {
		return 0
	}
	next := make([]int, len(needle))
	getNext(next, needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}
```

## 459. 重复的子字符串
1. 字符串匹配
```go
func repeatedSubstringPattern(s string) bool {
	if len(s) == 1 {
		return false
	}
	fast := 1
	lenS := len(s)
	i := 0
	for i < len(s) {
		fmt.Println(fast)
		isSub := true
		for j := 0; j < fast; j++ {
			if s[i+j] != s[j] {
				isSub = false
			}
		}
		if isSub == false {
			if i > lenS/2 {
				return false
			}
			for k := 1; fast+k < lenS; k++ {
				if fast+k > lenS/2 {
					return false
				}
				if s[fast+k] == s[0] && lenS%(fast+k) == 0 {
					fast += k
					i = 0
					break
				}
			}

		}
		i += fast
		if i == lenS && isSub == true {
			return true
		}
	}
	return false
}
```

2. KMP
```go
func repeatedSubstringPattern(s string) bool {
	if len(s) == 1 || len(s) == 0 {
		return false
	}
	j := 0
	lenS := len(s)
	next := make([]int, lenS)
	next[0] = j
	for i := 1; i < lenS; i++ {
		for j > 0 && s[j] != s[i] {
			j = next[j-1]
		}
		if s[j] == s[i] {
			j++
		}
		next[i] = j
	}
	if next[lenS-1] > 0 && lenS%(lenS-next[lenS-1]) == 0 {
		return true
	}
	return false

}

```