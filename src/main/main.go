package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const Hello = "Hello, YOLO!"

func hello(w http.ResponseWriter, _ *http.Request) {

	if _, err := io.WriteString(w, Hello); err != nil {

		log.Printf("main.go:hello(): %+v", err)
	}
}

func main() {

	http.HandleFunc("/", hello)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {

		log.Fatalf("main.go:main(): %+v", err)
	}
}
