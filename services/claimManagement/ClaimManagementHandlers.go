package claimManagement

// The 'Loans' board contains the general information about a loan
//
// Loans - Lifecycle
// --- --- --- --- ---
// Loan claim created -> PENDING
// Loan claim accepted -> OPEN
// Loan claim denied -> DENIED
// ---
// Settlement for Loan claim requested -> SETTLEMENT_REQUESTED
// Settlement for Loan claim accepted -> SETTLED
// Settlement for Loan claim denied -> OPEN

func handleCreate(req Request) {
	// TODO create new request
	// check if already exists, if yes and state:DENIED -> reopen
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
