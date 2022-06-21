package main

import (
	"log"
	"net/http"
)

func runServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO call handler function
		//go handler(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
