package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	username = "abc"
	password = "123"
)

type LoginRespose struct {
	LoginId int    `json:"loginId"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func main() {
	call("http://localhost:8080/example", "POST")
}

func call(url, method string) any {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	read, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Resposns body read: %s", err)
		return err
	}

	loginRespons := LoginRespose{}
	err = json.Unmarshal(read, &loginRespons)
	if err != nil {
		fmt.Printf("Resposne error: %s", err)
	}

	fmt.Println(loginRespons.Message)

	return loginRespons
}
