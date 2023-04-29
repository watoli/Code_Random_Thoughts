# Day11 栈与队列part02

# 20. 有效的括号
1. 栈
```go
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
```

# 1047. 删除字符串中的所有相邻重复项
1. 栈实现
```go
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
```

# 150. 逆波兰表达式求值
1. 栈实现
```go
func evalRPN(tokens []string) int {
	operatorMap := map[string]func(x, y int) int{
		"+": func(x, y int) int {
			return x + y
		},
		"-": func(x, y int) int {
			return  y-x
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
```