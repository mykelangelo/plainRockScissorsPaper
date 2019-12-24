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

type RequestBody struct {
	Message Message `json:"message"`
}

const UserGreeting = "Good day to you, kind sir! How may I be of service today?"

func hello(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	logality(json.NewDecoder(r.Body).Decode(&requestBody), "hello().decode")

	marshalled, _ := json.Marshal(requestBody.Message)
	log.Printf("user's message info: `%s`", marshalled)

	//fatality(json.NewEncoder(w).Encode(nil), "hello().encode")

	POST(requestBody.Message.Chat.ID)
}

func POST(id int64) {
	PostUrl := "https://api.telegram.org/bot" + os.Getenv("bot_token") + "/sendMessage"

	replyMarkup := ReplyKeyboardMarkup{Keyboard: [][]KeyboardButton{
		{
			{"7", false, false},
			{"8", false, false},
			{"9", false, false},
		}, {
			{"4", false, false},
			{"5", false, false},
			{"6", false, false},
		}, {
			{"1", false, false},
			{"2", false, false},
			{"3", false, false},
		}, {
			{"0", false, false},
		},
	},
		ResizeKeyboard:  true,
		OneTimeKeyboard: false,
	}

	data := url.Values{}
	data.Set("chat_id", strconv.Itoa(int(id)))
	data.Set("text", UserGreeting)
	marshalled, err2 := json.Marshal(replyMarkup)
	logality(err2, "marshalling replyMarkup")
	data.Set("reply_markup", string(marshalled))

	log.Printf("keyboard markup, it must be: %s", string(marshalled))
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
