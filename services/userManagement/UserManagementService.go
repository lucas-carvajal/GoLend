package userManagement

import (
	"database/sql"
	"fmt"
)

var channel chan Request
var db *sql.DB

func Init(inputDb *sql.DB) chan Request {
	fmt.Println("[UserManagementService Channel] Creating Channel")
	channel = make(chan Request)
	db = inputDb

	go run()

	return channel
}

func run() {
	for in := range channel {
		switch in.Command {
		case "SIGNUP":
			handleSignup(in)
		case "LOGIN":
			handleLogin(in)
		}
	}
	fmt.Println("[UserManagementService Channel] Closed")
}
