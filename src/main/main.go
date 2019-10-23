package main

import (
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, _ *http.Request) {

	var _, err = io.WriteString(w, "Hello, YOLO!")

	if err != nil {
		print("hello", err.Error())
	}
}

func main() {

	http.HandleFunc("/", hello)

	var err = http.ListenAndServe(":"+getPort(), nil)

	if err != nil {
		print("main", err.Error())
	}
}

func getPort() string {

	var port = os.Getenv("PORT")

	print("(heroku:" + port + ")")

	if len(port) != 4 {
		port = "8000"
	}

	return port
}
