package claimManagement

type loanStatus string

const (
	PENDING              loanStatus = "PENDING"
	OPEN                            = "OPEN"
	DENIED                          = "DENIED"
	SETTLEMENT_REQUESTED            = "SETTLEMENT_REQUESTED"
	SETTLED                         = "SETTLED"
)

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
