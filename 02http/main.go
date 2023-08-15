package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("http call https://jsonplaceholder.typicode.com/posts")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		log.Printf("Requiest error: %s", err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	bodyString := string(body)
	log.Print(bodyString)

}
