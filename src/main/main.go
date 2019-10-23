package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
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

const UserGreeting = "Good day to you, kind sir! How may I be of service today?"

func hello(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	logality(json.NewDecoder(r.Body).Decode(&requestBody), "hello().decode")

	log.Printf("user wrote: `%s`", requestBody.Message.Text)

	fatality(json.NewEncoder(w).Encode(nil), "hello().encode")

	POST(requestBody.Message.Chat.ID)
}

func POST(id int) {
	PostUrl := "https://api.telegram.org/bot" + os.Getenv("bot_token") + "/sendMessage"

	data := url.Values{}
	data.Set("chat_id", string(id))
	data.Set("text", UserGreeting)

	log.Printf("dataset:<%s, %s>", data.Get("chat_id"), data.Get("text"))

	newRequest, err := http.NewRequest("POST", PostUrl, strings.NewReader(data.Encode()))
	fatality(err, "POST().newReq")

	_, err = http.DefaultClient.Do(newRequest)
	fatality(err, "POST().doReq")
}

func main() {

	http.HandleFunc("/", hello)

	fatality(http.ListenAndServe(":"+os.Getenv("PORT"), nil), "main()")
}

func fatality(err error, place string) {

	if err != nil {

		log.Fatalf("main.go:%s: %+v", place, err)
	}
}

func logality(err error, place string) {

	if err != nil {

		log.Printf("main.go:%s: %+v", place, err)
	}
}
