package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ProxyMessage struct {
	WEBHOOK_URL string  `json:"webhook_url,omitempty"`
	Data        webhook `json:"data,omitempty"`
}

type fields struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type thumbnail struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type author struct {
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type provider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type video struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type image struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}
type footer struct {
	Text         string `json:"text,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type embed struct {
	Title       string    `json:"title,omitempty"`
	Type        string    `json:"type,omitempty"`
	Description string    `json:"description,omitempty"`
	URL         string    `json:"url,omitempty"`
	Timestamp   string    `json:"timestamp,omitempty"`
	Color       int       `json:"color,omitempty"`
	Footer      footer    `json:"footer,omitempty"`
	Image       image     `json:"image,omitempty"`
	Thumbnail   thumbnail `json:"thumbnail,omitempty"`
	Video       video     `json:"video,omitempty"`
	Provider    provider  `json:"provider,omitempty"`
	Author      author    `json:"author,omitempty"`
	Fields      []fields  `json:"fields,omitempty"`
}

type webhook struct {
	Content    string  `json:"content,omitempty"`
	Username   string  `json:"username,omitempty"`
	Avatar_url string  `json:"avatar_url,omitempty"`
	Tts        bool    `json:"tts,omitempty"`
	Embeds     []embed `json:"embeds,omitempty"`
	// Allowed_mentions *allowed_mentions `json:"allowed_mentions"`
	// Components       *components       `json:"components"`
	// Files            *files            `json:"files"`
	// Payload_json     *payload_json     `json:"payload_json"`
	// Attachments      *attachments      `json:"attachments"`
	Flags       int    `json:"flags,omitempty"`
	Thread_name string `json:"thread_name,omitempty"`
	// Applied_tags *applied_tags `json:"applied_tags"`
	// Poll         *poll         `json:"poll"`
}

const (
	PORT = "8080"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &homeHandler{})
	http.ListenAndServe(":"+PORT, mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer r.Body.Close()

		var PM ProxyMessage
		err = json.Unmarshal(body, &PM)
		if err != nil {
			fmt.Println(err)
			return
		}
		webhookData, err := json.Marshal(PM.Data)
		if err != nil {
			fmt.Println(err)
			return
		}

		res, err := http.Post(PM.WEBHOOK_URL, "application/json", bytes.NewBuffer(webhookData))
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(res.StatusCode)

	}
}
