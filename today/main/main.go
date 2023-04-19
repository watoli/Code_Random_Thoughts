package main

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for {
		if end < start {
			return -1
		} else if nums[(end-start)/2+start] == target {
			return (end-start)/2 + start
		}
		if nums[(end-start)/2+start] < target {
			start = (end-start)/2 + start + 1
		} else {
			end = (end-start)/2 + start - 1
		}
	}
}

func main() {
	fmt.Println("Hello World")
}
