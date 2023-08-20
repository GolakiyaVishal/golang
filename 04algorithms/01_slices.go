package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{11, 12, 13}

	c1 := append(slice2, slice1...)
	sort.Ints(c1)

	c1 = append(c1, 100)
	c1 = c1[:]
	fmt.Println(c1)
	fmt.Println(c1[0])
}
