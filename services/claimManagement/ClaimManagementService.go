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

		// TODO check authentication

		switch req.Action {
		case CREATE:
			handleCreate(req)
		case ACCEPTCLAIM:
			handleAcceptClaim(req)
		case DENYCLAIM:
			handleDenyClaim(req)
		case REQUESTSETTLEMENT:
			handleRequestSettlement(req)
		case ACCEPTSETTLEMENT:
			handleAcceptSettlement(req)
		case DENYSETTLEMENT:
			handleDenySettlement(req)
		case DEADLINE:
			handleSetDeadline(req)
		case INTEREST:
			handleSetInterest(req)
		}
	}
	fmt.Println("[ClaimManagementService Channel] Closed")
}
