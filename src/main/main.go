package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Chat struct {
	ID int `json:"id"`
}
type Message struct {
	Chat                Chat   `json:"chat"`
	Text                string `json:"text"`
	ReplyKeyboardMarkup Markup `json:"reply_markup"`
}
type RequestBody struct {
	Message Message `json:"message"`
}
type Keyboard struct {
	Text           [][]string `json:"text"`
	RequestContact bool       `json:"request_contact"`
}
type Markup struct {
	Keyboard        Keyboard `json:"keyboard"`
	ResizeKeyboard  bool     `json:"resize_keyboard"`
	OneTimeKeyboard bool     `json:"one_time_keyboard"`
}

const UserGreeting = "Good day to you, kind sir! How may I be of service today?"

func hello(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	logality(json.NewDecoder(r.Body).Decode(&requestBody), "hello().decode")

	log.Printf("user wrote: `%s`", requestBody.Message.Text)

	fatality(json.NewEncoder(w).Encode(nil), "hello().encode")

	POST(requestBody.Message.Chat.ID, requestBody.Message.ReplyKeyboardMarkup)
}

func POST(id int, markup Markup) {
	PostUrl := "https://api.telegram.org/bot" + os.Getenv("bot_token") + "/sendMessage"

	keyboard := Keyboard{Text: [][]string{
		{"7", "8", "9"},
		{"4", "5", "6"},
		{"1", "2", "3"},
		{"0"},
	}}

	replyMarkup := markup
	replyMarkup.Keyboard = keyboard
	replyMarkup.ResizeKeyboard = true
	replyMarkup.OneTimeKeyboard = true

	data := url.Values{}
	data.Set("chat_id", strconv.Itoa(id))
	data.Set("text", UserGreeting)
	//data.Set("reply_markup.resize_keyboard", strconv.FormatBool(replyMarkup.ResizeKeyboard))
	//data.Set("reply_markup.one_time_keyboard", strconv.FormatBool(replyMarkup.OneTimeKeyboard))
	marshalled, err2 := json.Marshal(replyMarkup)
	logality(err2, "marshalling replyMarkup")
	//data.Set("reply_markup", string(marshalled))

	log.Printf("cejvo klaviaturna rozkladka chy sho: %s", string(marshalled))
	log.Printf("dataset:<%s, %s>", data.Get("chat_id"), data.Get("text"))

	newRequest, err := http.NewRequest("POST", PostUrl, strings.NewReader(data.Encode()))
	fatality(err, "POST().newReq")

	newRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	newRequest.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	resp, err := client.Do(newRequest)
	fatality(err, "POST().doReq")

	_, err = fmt.Printf("{status: %s}", resp.Status)
	fatality(err, "POST().printStatus")
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
