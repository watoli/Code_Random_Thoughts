package main

import "fmt"

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
