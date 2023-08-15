package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Posts []struct {
	UserId int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	params := url.Values{}
	params.Add("userId", "1")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?" + params.Encode())

	if err != nil {
		log.Printf("Requetst failed: %s", err)
		return
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	// bodyString := string(body)
	// log.Print(bodyString)

	// Unmarshal result
	posts := Posts{}
	err = json.Unmarshal(body, &posts)

	if err != nil {
		log.Printf("Error in unmarshal: %s", err)
		return
	}

	for i := 0; i < len(posts); i++ {
		fmt.Printf("%d: %s", i, posts[i].Title)
		fmt.Println()
	}

}
