# Day8字符串 part01

## 454. 四数相加 II

1. 哈希表，和两数相加本质上没区别，干就完了！
```go
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) (res int) {
	mapAB := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if _, ok := mapAB[nums1[i]+nums2[j]]; ok {
				mapAB[nums1[i]+nums2[j]]++
			} else {
				mapAB[nums1[i]+nums2[j]] = 1
			}
		}
	}
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			if _, ok := mapAB[-nums3[i]-nums4[j]]; ok {
				res += mapAB[-nums3[i]-nums4[j]]
			}
		}
	}
	return
}
```

## 383. 赎金信
1. 哈希，也可以用数组去做，空间会更节省一些
```go
func canConstruct(ransomNote string, magazine string) bool {
	mapMag := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		if _, ok := mapMag[magazine[i]]; !ok {
			mapMag[magazine[i]] = 1
		} else {
			mapMag[magazine[i]]++
		}
	}
	for i := 0; i < len(ransomNote); i++ {
		if _, ok := mapMag[ransomNote[i]]; !ok {
			return false
		}
		if mapMag[ransomNote[i]] == 0 {
			return false
		}
		mapMag[ransomNote[i]]--
	}
	return true
}
```

## 15. 三数之和
1. 没想到可以用双指针，好困啊，顶不住了要
```go
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1
	res := make([][]int, 0)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left = i + 1
		right = len(nums) - 1
		for left < right {
			if nums[left]+nums[right] == -nums[i] {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left]+nums[right] > -nums[i] {
				right--

			} else {
				left++

			}
		}
	}
	return res
}
```

## 18. 四数之和
1. 双指针，没什么好说的，和三数之和一模一样
```go
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if target > 0 && nums[i] > target {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if target > 0 && nums[j] > target-nums[i] {
				continue
			}
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left = j + 1
			right = len(nums) - 1
			for left < right {
				if nums[left]+nums[right] == target-nums[i]-nums[j] {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					right--
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if nums[i]+nums[j]+nums[left]+nums[right] > target {
					right--

				} else {
					left++

				}
			}

		}
	}
	return res
}
```