package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print(getMinAverageIndex([]int{1, 5, 3, 4, 5, 1, 1}))
}

func getMinAverageIndex(A []int) int {
	index := 0
	minAverage := math.MaxInt64

	for i := 0; i < len(A)-2; i++ {
		if (A[i]+A[i+1])/2 < minAverage {
			index = i
			minAverage = (A[i] + A[i+1]) / 2
		}

		if (A[i]+A[i+1]+A[i+2])/3 < minAverage {
			index = i
			minAverage = (A[i] + A[i+1] + A[i+2]) / 3
		}
	}

	return index
}
