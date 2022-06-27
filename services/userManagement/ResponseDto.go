package userManagement

type statusMessage string

const (
	success statusMessage = "SUCCESS"
	error                 = "ERROR"
)

type signupResponse struct {
	id      int
	status  statusMessage
	message string
}

type loginResponse struct {
	id      int
	status  statusMessage
	message string
	// TODO add token
}
