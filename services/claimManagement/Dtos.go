package claimManagement

type action string

const (
	CREATE            action = "CREATE"
	ACCEPTCLAIM              = "ACCEPTCLAIM"
	DENYCLAIM                = "DENYCLAIM"
	REQUESTSETTLEMENT        = "REQUESTSETTLEMENT"
	ACCEPTSETTLEMENT         = "ACCEPTSETTLEMENT"
	DENYSETTLEMENT           = "DENYSETTLEMENT"
	DEADLINE                 = "DEADLINE"
	INTEREST                 = "INTEREST"
)

type Request struct {
	Action          action
	Token           string
	ResponseChannel chan Response
	Loan            Loan
}

type Response struct {
	Status  int
	Message string
}
