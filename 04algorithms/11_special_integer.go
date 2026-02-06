package main

import "fmt"

// Given an array of k of integers, find the special integer x
// such that there are x integers in the array k that are larger than or equal to x.
// The special integer dosen't have to exists in array k.
//
// If no special integer are found, return -1.

func main() {
	i := solution([]int{0, 4, 1, 0, 4}, 5)
	fmt.Print(i)
}

func solution(k []int, x int) int {
	count := 0
	for _, e := range k {
		if e >= x {
			count++
		}
	}

	if count == 0 {
		return -1
	}

	return count
}
