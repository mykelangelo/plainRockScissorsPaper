package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {

	var _, err = io.WriteString(w, "Yello, YOLO!")

	if err != nil {
		print(err)
	}
}

func main() {

	http.HandleFunc("/", hello)

	var err = http.ListenAndServe(":71", nil)

	if err != nil {
		print(err)
	}
}
