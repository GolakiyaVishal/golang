package main

import "fmt"

func main() {
	fmt.Print(findValues([]int32{2, 4, 7, 11, 15}, 22))
}

func findValues(nums []int32, target int32) (int32, int32) {
	left := 0
	right := len(nums) - 1

	for left <= right {
		sum := nums[left] + nums[right]
		fmt.Println(sum, left, right)

		if sum == target {
			return nums[left], nums[right]
		}

		if target > sum {
			left = left + 1
		}

		if target < sum {
			right = right - 1
		}
	}

	return -1, -1
}
