package main

import (
	"fmt"
)

func reverseString(s []byte) {
	lenS := len(s)
	for i := 0; i < lenS/2; i++ {
		s[i], s[lenS-1-i] = s[lenS-1-i], s[i]
	}
}

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

func reverseLeftWords(s string, n int) string {
	res := []byte(s)
	reverseString(res[:n])
	reverseString(res[n:])
	reverseString(res)
	return fmt.Sprintf("%s", res)
}
