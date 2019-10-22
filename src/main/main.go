package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {

	var _, err = io.WriteString(w, "Hello, heroku visitors!")

	if err != nil {
		print(err)
	}
}

func main() {

	http.HandleFunc("/", hello)

	var err = http.ListenAndServe(":80", nil)

	if err != nil {
		print(err)
	}
}
