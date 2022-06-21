package main

import (
	"log"
	"net/http"
)

func runServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go handleRequest(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
