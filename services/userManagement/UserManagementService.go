package userManagement

import (
	"database/sql"
	"fmt"
	"gobank.com/services/tokenManagement"
)

var channel chan Request
var db *sql.DB
var tokenChannel chan tokenManagement.Request

func Init(
	inputDb *sql.DB,
	tokenManagementServiceChannel chan tokenManagement.Request,
) chan Request {
	fmt.Println("[UserManagementService Channel] Creating Channel")
	channel = make(chan Request)
	db = inputDb
	tokenChannel = tokenManagementServiceChannel

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
