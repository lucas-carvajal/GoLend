package claimManagement

// TODO define lifecycle of claims

func handleCreate(req Request) {
	// TODO create new request
}

func handleAcceptClaim(req Request) {
	// TODO move state of claim to accept
}

func handleDenyClaim(req Request) {
	// TODO move state of claim to denied
}

func handleRequestSettlement(req Request) {
	// TODO request settlement of claim
}

func handleAcceptSettlement(req Request) {
	// TODO move claim to settled
}

func handleDenySettlement(req Request) {
	// TODO move claim back to open
}

func handleSetDeadline(req Request) {
	// TODO set deadline for request
}

func handleSetInterest(req Request) {
	// TODO set interest for request
}
