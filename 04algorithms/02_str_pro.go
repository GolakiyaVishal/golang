package main

import (
	"fmt"
	"strings"
)

func main() {
	terms := "vishal"
	syntex := "find vishal from this string"

	fmt.Println(strings.Contains(syntex, terms))
}
