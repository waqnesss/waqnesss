package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", Hello)

	log.Printf("server is starting at :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("error of starting server", err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, my name is Waqness.\nI'm Golang developer\n I've been doing this for 4 months now, working on cerating websites and applications."))
	w.Write([]byte("\nI'ts my first website written in Go"))
}
