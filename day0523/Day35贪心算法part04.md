# Day35贪心算法part04

## 860.柠檬水找零 
```go
func lemonadeChange(bills []int) bool {
	cashMap := map[int]int{5: 0, 10: 0, 20: 0}
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			cashMap[5] += 1
		} else if bills[i] == 10 {
			if cashMap[5] == 0 {
				return false
			} else {
				cashMap[10] += 1
				cashMap[5] -= 1
			}
		} else {
			if cashMap[10] > 0 && cashMap[5] > 0 {
				cashMap[10] -= 1
				cashMap[5] -= 1
			} else if cashMap[5] > 2 {
				cashMap[5] -= 3
			} else {
				return false
			}
		}
	}
	return true
}
```
## 406.根据身高重建队列
```go
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		} else if people[i][0] < people[j][0] {
			return false
		} else {
			if people[i][1] <= people[j][1] {
				return true
			} else {
				return false
			}
		}
	})
	queue := make([][]int, 0)
	for i := 0; i < len(people); i++ {
		if people[i][1] == len(queue) {
			queue = append(queue, people[i])
		} else {
			queue = append(queue[:people[i][1]], append([][]int{people[i]}, queue[people[i][1]:]...)...)
		}
	}
	return queue

}
```

## 452. 用最少数量的箭引爆气球
```go
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] < points[j][0] {
			return true
		} else if points[i][0] > points[j][0] {
			return false
		} else {
			if points[i][1] < points[j][1] {
				return true
			}
			return false
		}
	})
	//fmt.Println(points)
	idx := 0
	res := 1
	start := points[idx][0]
	limit := points[idx][1]
	for i := idx + 1; i < len(points); i++ {
		if points[i][1] < limit {
			limit = points[i][1]
		}
		start = points[i][0]
		if start > limit {
			res++
			start = points[i][0]
			limit = points[i][1]
		}
	}
	return res
}
```