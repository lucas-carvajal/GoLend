package main

import (
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println("+++=========================+++")

	// TODO handle different routes
}
