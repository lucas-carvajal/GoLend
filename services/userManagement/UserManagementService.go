package userManagement

import (
	"database/sql"
	"fmt"
)

var channel chan string
var db *sql.DB

func Init(inputDb *sql.DB) chan string {
	channel = make(chan string)
	db = inputDb

	go run()

	return channel
}

func run() {
	for in := range channel {
		fmt.Println("[Channel] Message received: ", in)

	}
	fmt.Println("[Channel] Closed")
}
