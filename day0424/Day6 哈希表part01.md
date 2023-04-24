# Day6 哈希表part01

## 242. 有效的字母异位词
1. 哈希表
```go
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := sMap[s[i]]; !ok {
			sMap[s[i]] = 1
		} else {
			sMap[s[i]]++
		}
	}
	for i := 0; i < len(t); i++ {
		if _, ok := sMap[t[i]]; !ok {
			return false
		} else {
			sMap[t[i]]--
		}
	}
	for _, v := range sMap {
		if v != 0 {
			return false
		}
	}
	return true

}

```

## 349. 两个数组的交集 
1. 没多想，哈希表,看答案可以用数组做哈希，挺有意思
```go
func intersection(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]bool)
	res := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		if _, ok := numMap[nums1[i]]; !ok {
			numMap[nums1[i]] = true
		}
	}
	for i := 0; i < len(nums2); i++ {
		if v, ok := numMap[nums2[i]]; ok {
			if v == true {
				res = append(res, nums2[i])
				numMap[nums2[i]] = false
			}
		}
	}
	return res
}
```

## 202. 快乐数
1. 直接哈希
```go
func isHappy(n int) bool {
	if n == 0 {
		return false
	} else if n == 1 {
		return true
	}
	happyMap := make(map[int]bool)
	for n != 1 {
		n = getSum(n)
		if _, ok := happyMap[n]; ok {
			return false
		}
		happyMap[n] = true
	}
	return true
}
```
2. 我应该想到，有环的东西就可以有快慢指针的
```go
func isHappy(n int) bool {
	if n == 1 || getSum(n) == 1 || getSum(getSum(n)) == 1 {
		return true
	}
	fast := getSum(getSum(n))
	slow := getSum(n)
	for fast != 1 && fast != slow {
		fast = getSum(getSum(fast))
		slow = getSum(slow)
	}
	return fast == 1
}
```
## 1. 两数之和
1. 直接哈希
```go
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if pos, ok := numMap[target-nums[i]]; ok {
			return []int{pos, i}
		}
		if _, ok := numMap[nums[i]]; !ok {
			numMap[nums[i]] = i
		}
	}
	return []int{}
}
```