package main

import "fmt"

func runningSums(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return nums
}

func main() {
	fmt.Println(runningSums([]int{1, 2, 3, 4}))
}
