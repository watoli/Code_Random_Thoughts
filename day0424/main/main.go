package main

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

func getSum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

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
