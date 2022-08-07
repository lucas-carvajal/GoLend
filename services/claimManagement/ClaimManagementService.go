package claimManagement

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
	fmt.Println("[ClaimManagementService Channel] Creating Channel")
	channel = make(chan Request)
	db = inputDb
	tokenChannel = tokenManagementServiceChannel

	go run()

	return channel
}

func run() {
	for req := range channel {
		switch req.Action {
		case CREATE:
		case ACCEPT:
		case DENY:
		case REQUESTSETTLEMENT:
		case ACCEPTSETTLEMENT:
		case DENYSETTLEMENT:
		case DEADLINE:
		case INTEREST:
		}
	}
	fmt.Println("[ClaimManagementService Channel] Closed")
}
