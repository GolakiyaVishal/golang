package main

import "fmt"

func main2() {
	// fmt.Println(strings.Replace("a ", " ", "", 0))

	// fmt.Println(lengthOfLastWord("a "))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	rm := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	total := 0
	chars := []rune(s)
	for i := 0; i < len(chars)-1; i++ {
		t := rm[string(chars[i])]
		if t < rm[string(chars[i+1])] {
			total -= t
		} else {
			total += t
		}
	}
	total += rm[string(chars[len(chars)-1])]
	return total
}
