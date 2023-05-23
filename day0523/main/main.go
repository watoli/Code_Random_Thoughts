package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

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

/*
[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
[7,0] [7,1] [6,1] [5,0] [5,2] [6,1]
[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
*/
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

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("欢迎来到优化版射击反应速度游戏！")
	fmt.Println("当屏幕上出现目标时，请尽快按下回车键进行射击。")

	// 等待用户准备好
	fmt.Print("按下回车键开始游戏：")
	fmt.Scanln()

	playGame()
}

func playGame() {
	score := 0
	rounds := 0

	// 创建通道用于接收用户输入
	inputCh := make(chan string)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			inputCh <- strings.TrimSpace(input)
		}
	}()

	for {
		rounds++
		fmt.Printf("第 %d 轮开始！\n", rounds)

		// 随机等待时间
		waitTime := rand.Intn(3000) + 1000
		time.Sleep(time.Duration(waitTime) * time.Millisecond)

		fmt.Println("目标出现！")
		startTime := time.Now()

		select {
		case <-inputCh:
			endTime := time.Now()
			reactionTime := endTime.Sub(startTime)

			if reactionTime.Seconds() < 0.2 {
				fmt.Println("命中！")
				score++
			} else {
				fmt.Println("未命中！")
			}

			fmt.Printf("反应时间：%.2f 秒\n", reactionTime.Seconds())
			fmt.Printf("得分：%d\n", score)

		case <-time.After(2 * time.Second):
			fmt.Println("未命中！")
			fmt.Println("超时！")
		}

		// 再次询问用户是否继续游戏
		fmt.Print("是否继续游戏？(y/n): ")
		go func() {
			var choice string
			fmt.Scanln(&choice)
			inputCh <- strings.TrimSpace(choice)
		}()

		choice := <-inputCh

		if strings.ToLower(choice) == "y" {
			continue
		} else {
			fmt.Println("游戏结束。")
			break
		}
	}
}
