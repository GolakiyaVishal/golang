package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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

	fmt.Printf("special count: %d", specialCount([]int{0, 4, 1, 0, 4}, 2))
	fmt.Printf("special count: %d", specialCount([]int{0, 1, 1, 0, 1}, 2))
	fmt.Println()
	fmt.Println("Score:")
	fmt.Println(getScore([]string{"5", "2", "C", "D", "+"}))
	fmt.Println()
	fmt.Println("check for valid parentheses")
	fmt.Println(isValidParentheses(""))
}

func specialCount(list []int, x int) int {
	j := 0

	for i := 0; i < len(list); i++ {
		if list[i] >= x {
			j++
		}
	}
	return j
}

func getScore(list []string) int {
	sum := 0
	l1 := make([]int, 0)

	for i := 0; i < len(list); i++ {
		switch list[i] {
		case "+":
			v := l1[len(l1)-1] + l1[len(l1)-2]
			l1 = append(l1, v)
			break

		case "C":
			l1 = l1[:len(l1)-1]
			break

		case "D":
			v := l1[len(l1)-1]
			l1 = append(l1, v*2)
			break

		default:
			v1, _ := strconv.Atoi(list[i])
			l1 = append(l1, v1)
		}
	}

	for _, i := range l1 {
		sum += i
	}

	return sum
}

func isValidParentheses(str string) bool {
	l1 := make([]string, 0)
	list := []rune(str)

	for _, j := range list {
		i := string(j)
		if strings.Contains("({[", i) {
			l1 = append(l1, i)
		} else {
			if i == ")" && l1[len(l1)-1] != "(" || i == "}" && l1[len(l1)-1] != "{" || i == "]" && l1[len(l1)-1] != "[" {
				return false
			} else {
				l1 = l1[:len(l1)-1]
			}
		}
	}
	return len(l1) == 0
}
