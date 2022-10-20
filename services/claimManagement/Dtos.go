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
	responseChannel chan Response
	Id              int
	Title           string
	Description     string
	toUsername      string
	amount          uint64
	interest        uint64
}

type Response struct {
	Status  int
	Message string
}
