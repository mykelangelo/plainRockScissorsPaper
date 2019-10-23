package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {

	var _, err = io.WriteString(w, "Yello, YOLO!")

	if err != nil {
		print("hello", err.Error())
	}
}

func main() {

	http.HandleFunc("/", hello)

	var err = http.ListenAndServe(":8080", nil)

	if err != nil {
		print("main", err.Error())
	}
}
