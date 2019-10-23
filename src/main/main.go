package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

const Hello = "Hello, YOLO!"

type Chat struct {
	ID string `json:"id"`
}
type MessagePtr struct {
	Chat Chat `json:"chat"`
}
type Body struct {
	Message MessagePtr `json:"message"`
}

func hello(w http.ResponseWriter, r *http.Request) {

	var val Body

	if err := json.NewDecoder(r.Body).Decode(&val); err != nil {

		log.Printf("main.go:hello.writeHello(): %+v", err)
	}

	log.Printf("%s", val.Message.Chat.ID)

	if _, err := io.WriteString(w, Hello); err != nil {

		log.Printf("main.go:hello.writeHello(): %+v", err)
	}
}

func main() {

	http.HandleFunc("/", hello)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {

		log.Fatalf("main.go:main(): %+v", err)
	}
}
