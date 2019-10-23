package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Chat struct {
	ID int `json:"id"`
}
type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}
type RequestBody struct {
	Message Message `json:"message"`
}

type ResponseBody struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

const UserGreeting = "Good day to you, kind sir! How may I be of service today?"

func hello(w http.ResponseWriter, r *http.Request) {

	var requestBody RequestBody

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {

		log.Printf("main.go:hello().read: %+v", err)
	}

	log.Printf("user wrote: `%s`", requestBody.Message.Text)

	if err := json.NewEncoder(w).Encode(&ResponseBody{
		ChatId: requestBody.Message.Chat.ID,
		Text:   UserGreeting,
	}); err != nil {

		log.Printf("main.go:hello().write: %+v", err)
	}
}

func main() {

	http.HandleFunc("/", hello)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {

		log.Fatalf("main.go:main(): %+v", err)
	}
}
