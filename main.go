package main

import "net/http"

func main() {

	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/", &homeHandler{})

	// Run the server
	http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// const WEBHOOK = "https://discord.com/api/webhooks/1334528477549821993/eS0_aRfbAl9JTqzKsBRQVCxWhcJD61BU97jrLXQ5Ig88zPf9Zdy_PiSrf033aoq_0Knf"

// type ProxyMessage struct {
// 	WEBHOOK_URL string
// 	Data        interface{}
// }
// type WebHookMessage struct {
// 	Content string `json:"content"`
// }

// type homeHandler struct{}

// func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("This is my home page"))
// 	var msg WebHookMessage

// 	msg.Content = " TEST "
// 	body, err := json.Marshal(msg)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	res, err := http.Post(WEBHOOK, "application/json", bytes.NewBuffer(body))
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(res.StatusCode)
// }

// func main() {

// 	// Create a new request multiplexer
// 	// Take incoming requests and dispatch them to the matching handlers
// 	mux := http.NewServeMux()

// 	// Register the routes and handlers
// 	mux.Handle("/", &homeHandler{})

// 	// Run the server
// 	http.ListenAndServe(":8080", mux)
// }
