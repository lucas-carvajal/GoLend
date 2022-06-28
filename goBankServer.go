package main

import (
	"gobank.com/services/userManagement"
	"log"
	"net/http"
)

func runServer(userMgmChan chan userManagement.Request) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go handleRequest(w, r, userMgmChan)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
