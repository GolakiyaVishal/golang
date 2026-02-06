package main

import "fmt"

// Given a string s containing just character '{','}','(',')','[' and ']'
// determind the input string is valid
// An input string is valid if
// 	1. Open bracket must be closed by the same type of brackets.
// 	2. Open bracket must be closed in the correct order.

func main() {
	fmt.Print(solutions("[]({{})"))
}

func solutions(s string) bool {
	st := []string{}

	for _, e := range []rune(s) {
		if ")" == string(e) || "}" == string(e) || "]" == string(e) {
			if len(st) == 0 {
				return false
			}

			te := st[len(st)-1]
			if (string(e) == ")" && te != "(") ||
				(string(e) == "}" && te != "{") ||
				(string(e) == "]" && te != "[") {
				return false
			}
			st = st[:len(st)-1]
		} else {
			st = append(st, string(e))
		}
	}

	return len(st) == 0
}
