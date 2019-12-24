package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type RequestBody struct {
	Message Message `json:"message"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	logality(json.NewDecoder(r.Body).Decode(&requestBody), "hello().decode")

	marshalled, err := json.Marshal(requestBody.Message)
	logality(err, "marshalling requestBody.Message")
	log.Printf("user's message info: `%s`", marshalled)

	POST(requestBody.Message.Chat.ID, requestBody.Message.Text)
}

const STONE = "🗿"
const SCISSORS = "✂"
const PAPER = "🗒"

var MOVES = []string{STONE, SCISSORS, PAPER}

func POST(id int64, text string) {
	PostUrl := "https://api.telegram.org/bot" + os.Getenv("bot_token") + "/sendMessage"

	replyMarkup := ReplyKeyboardMarkup{Keyboard: [][]KeyboardButton{
		{
			{MOVES[0], false, false},
			{MOVES[1], false, false},
			{MOVES[2], false, false},
		},
	},
		ResizeKeyboard:  true,
		OneTimeKeyboard: false,
	}

	data := url.Values{}
	data.Set("chat_id", strconv.Itoa(int(id)))
	var UserGreeting = "Good day to you, kind sir! How may I be of service today?"

	UserGreeting = makeMove(text, UserGreeting)

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

func makeMove(text string, UserGreeting string) string {
	if contains(MOVES, text) {
		move := generateMove()
		UserGreeting = move + "\n"
		const UserWins = "Yo, you do win! 🏆"
		const BotWins = "I win! 😎"
		if move == text {
			UserGreeting += "It's a draw, mate! 🤷🙃‍"
		} else if move == STONE {
			if text == PAPER {
				UserGreeting += UserWins
			} else {
				UserGreeting += BotWins
			}
		} else if move == PAPER {
			if text == STONE {
				UserGreeting += BotWins
			} else {
				UserGreeting += UserWins
			}
		} else if move == SCISSORS {
			if text == PAPER {
				UserGreeting += BotWins
			} else {
				UserGreeting += UserWins
			}
		}
	}
	return UserGreeting
}

func generateMove() string {
	rand.Seed(time.Now().UnixNano())

	return MOVES[rand.Intn(len(MOVES))]
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
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
