package main

/*

An array A consisting of N integers is given. Rotation of the array means that each element is shifted right by one index, and the last element of the array is moved to the first place. For example, the rotation of array A = [3, 8, 9, 7, 6] is [6, 3, 8, 9, 7] (elements are shifted right by one index and 6 is moved to the first place).

The goal is to rotate array A K times; that is, each element of A will be shifted to the right K times.

Write a function:

func Solution(A []int, K int) []int

that, given an array A consisting of N integers and an integer K, returns the array A rotated K times.

For example, given

    A = [3, 8, 9, 7, 6]
    K = 3
the function should return [9, 7, 6, 3, 8]. Three rotations were made:

    [3, 8, 9, 7, 6] -> [6, 3, 8, 9, 7]
    [6, 3, 8, 9, 7] -> [7, 6, 3, 8, 9]
    [7, 6, 3, 8, 9] -> [9, 7, 6, 3, 8]
For another example, given

    A = [0, 0, 0]
    K = 1
the function should return [0, 0, 0]

Given

    A = [1, 2, 3, 4]
    K = 4
the function should return [1, 2, 3, 4]

Assume that:

N and K are integers within the range [0..100];
each element of array A is an integer within the range [−1,000..1,000].

*/

import "fmt"

func main() {

	p := Solution([]int{3, 8, 9, 7, 6}, 3)
	fmt.Println(p)

	q := Solution([]int{0, 0, 0}, 1)
	fmt.Println(q)

	r := Solution([]int{1, 2, 3, 4}, 4)
	fmt.Println(r)
}

func Solution(A []int, K int) []int {
	// check K is zero or equal length to A
	if K == 0 || K == len(A) {
		return A
	}

	tk := K
	if K > len(A) {
		tk = K % len(A)
	}

	var ans = make([]int, len(A))
	lp := 0

	for index, element := range A {
		ni := index + tk

		if ni < len(A) {
			ans[ni] = element
			lp = index
		} else {
			ans[index-lp-1] = element
		}
	}

	return ans
}
