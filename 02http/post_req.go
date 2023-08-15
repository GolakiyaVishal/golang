package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Post struct {
	UserId string `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	params := url.Values{}
	params.Add("title", "New Post")
	params.Add("body", "New mbody text")
	params.Add("userId", "1")

	resp, err := http.PostForm("https://jsonplaceholder.typicode.com/posts", params)

	if err != nil {
		log.Printf("Request failed: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if isErr(err) {
		return
	}

	log.Printf("body::: %s", string(body))
	post := Post{}
	err = json.Unmarshal(body, &post)

	if isErr(err) {
		return
	}
	fmt.Printf("data: %s", post.Title)

}

func isErr(err error) bool {
	if err != nil {
		log.Printf("Error:: %s", err)
		return true
	}
	return false
}
