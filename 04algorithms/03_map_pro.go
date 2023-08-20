package main

import "fmt"

func main() {
	map1 := make(map[string]any)
	map1["user1"] = 10
	map1["user2"] = 11

	fmt.Println(map1)

	map1["user1"] = 13
	fmt.Println(map1)

	delete(map1, "user1")
	fmt.Println(map1)
}
