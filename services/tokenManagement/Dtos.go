package tokenManagement

type Request struct {
	Token           string
	ResponseChannel chan Response
}

type Response struct {
	OK       bool
	Email    string
	Username string
}
