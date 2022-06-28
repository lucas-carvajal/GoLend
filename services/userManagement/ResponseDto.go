package userManagement

type statusMessage string

const (
	success statusMessage = "SUCCESS"
	error                 = "ERROR"
)

type Response struct {
	Id      int
	Status  statusMessage
	Message string
	Token   string
}
