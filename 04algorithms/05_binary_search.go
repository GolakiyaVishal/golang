package main

import "fmt"

func main() {
	index := findIndex([]int32{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10)
	fmt.Print(index)
}

func findIndex(nums []int32, n int32) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (high + low) / 2
		fmt.Print(mid)
		if n == nums[mid] {
			return mid
		}

		if n < nums[mid] {
			high = mid - 1
		}

		if n > nums[mid] {
			low = mid + 1
		}
	}
	return -1
}
