package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	username = "abc"
	password = "123"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/example", handler)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	u, p, ok := r.BasicAuth()

	if !ok {
		fmt.Println("Error parsing basic auth.")
		loginFailed(w)
		return
	}

	if u != username {
		fmt.Printf("Username is wrong %s", u)
		loginFailed(w)
		return
	}

	if p != password {
		fmt.Printf("Password is wrong %s", p)
		loginFailed(w)
		return
	}

	fmt.Printf("Username: %s", u)
	fmt.Println()
	fmt.Printf("Password: %s", p)

	jResp := make(map[string]any)
	jResp["message"] = "Login Success"
	jResp["loginId"] = 1
	jResp["token"] = "asdfasdfasdfaerqwere"
	jData, err := json.Marshal(jResp)
	if err != nil {
		w.WriteHeader(402)
		fmt.Printf("Json error: %s", err)
		return
	}
	w.WriteHeader(200)
	w.Write(jData)

	return

}

func loginFailed(w http.ResponseWriter) {
	w.WriteHeader(401)
	eResp := make(map[string]any)
	eResp["message"] = "Login failed"
	eData, _ := json.Marshal(eResp)
	w.Write(eData)
}
