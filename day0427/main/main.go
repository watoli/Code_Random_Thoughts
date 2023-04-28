package main

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
