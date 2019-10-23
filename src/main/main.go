package main

import (
	"encoding/json"
	"io"
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
	ChatId   int    `json:"chat_id"`
	Text     string `json:"text"`
	contents string
	offset   int
}

func (b *ResponseBody) Read(p []byte) (int, error) {
	if b.offset >= len(b.contents) {
		return 0, io.EOF
	}
	n := copy(p, b.contents[b.offset:])
	b.offset += n
	return n, nil
}

const UserGreeting = "Good day to you, kind sir! How may I be of service today?"

func hello(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	logality(json.NewDecoder(r.Body).Decode(&requestBody), "hello().decode")

	log.Printf("user wrote: `%s`", requestBody.Message.Text)

	fatality(json.NewEncoder(w).Encode(&ResponseBody{
		ChatId: requestBody.Message.Chat.ID,
		Text:   UserGreeting,
	}), "hello().encode")

	POST(requestBody.Message.Chat.ID)
}

func POST(id int) {
	responseBody := ResponseBody{
		ChatId: id,
		Text:   UserGreeting,
	}

	PostUrl := "https://api.telegram.org/bot" + os.Getenv("bot_token") + "/sendMessage"

	newRequest, err := http.NewRequest("POST", PostUrl, &responseBody)
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
