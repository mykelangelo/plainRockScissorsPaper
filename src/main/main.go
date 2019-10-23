package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {

	print("hello-0")

	var _, err = io.WriteString(w, "Hello, YOLO!")

	print("hello-1")

	if err != nil {
		print("hello-err")
		print("hello", err.Error())
	}

	print("hello-end")
}

func main() {

	print("main-0")

	http.HandleFunc("/", hello)

	print("main-1")

	var err = http.ListenAndServe(":8000", nil)

	print("main-2")

	if err != nil {
		print("main-err")
		print("main", err.Error())
	}

	print("main-end")
}
