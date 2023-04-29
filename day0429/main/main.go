package main

import (
	"fmt"
	"strconv"
)

type T interface {
}

type Stack struct {
	data []T
	len  int
}

func NewStack() *Stack {
	return &Stack{
		data: []T{},
		len:  0,
	}
}

func (s *Stack) Push(c T) {
	s.data = append(s.data, c)
	s.len++
}

func (s *Stack) Pop() T {
	res := s.data[s.len-1]
	s.data = s.data[:s.len-1]
	s.len--
	return res
}

func (s *Stack) Peek() T {
	return s.data[s.len-1]
}

func (s *Stack) Size() int {
	return s.len
}
func (s *Stack) Empty() bool {
	if s.len == 0 {
		return true
	} else {
		return false
	}
}

func isValid(s string) bool {

	bracketMap := map[byte]int{
		'(': 0, ')': 1,
		'{': 0, '}': 1,
		'[': 0, ']': 1,
	}
	bracketMatch := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	bracketStack := NewStack()
	for i := 0; i < len(s); i++ {
		//fmt.Println(bracketStack.data)
		//fmt.Println(s[i])
		v, ok := bracketMap[s[i]]
		if ok {
			if v == 0 {
				bracketStack.Push(s[i])
			} else if v == 1 {
				if bracketStack.Empty() {
					return false
				}
				if m, _ := bracketMatch[bracketStack.Pop().(byte)]; m != s[i] {
					return false
				}
			}
		}
	}
	if bracketStack.Empty() {
		return true
	} else {
		return false
	}
}

func removeDuplicates(s string) string {
	resS := NewStack()
	for i := 0; i < len(s); i++ {
		if (!resS.Empty()) && resS.Peek() == s[i] {
			resS.Pop()
		} else {
			resS.Push(s[i])
		}
	}
	//fmt.Println(resS.data)
	lenS := resS.Size()
	res := make([]byte, lenS)
	for i := 0; i < lenS; i++ {
		res[lenS-1-i] = resS.Pop().(byte)
	}
	return fmt.Sprintf("%s", res)
}

func evalRPN(tokens []string) int {
	operatorMap := map[string]func(x, y int) int{
		"+": func(x, y int) int {
			return x + y
		},
		"-": func(x, y int) int {
			return y - x
		},
		"*": func(x, y int) int {
			return x * y
		},
		"/": func(x, y int) int {
			return y / x
		},
	}
	resS := NewStack()
	for i := 0; i < len(tokens); i++ {
		funcOper, ok := operatorMap[tokens[i]]
		if !ok {
			atoi, err := strconv.Atoi(tokens[i])
			if err != nil {
				return -1
			}
			resS.Push(atoi)
		} else {
			resS.Push(funcOper(resS.Pop().(int), resS.Pop().(int)))
		}
	}
	return resS.Pop().(int)

}
