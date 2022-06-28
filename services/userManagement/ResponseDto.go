package userManagement

type statusMessage string

const (
	success statusMessage = "SUCCESS"
	error                 = "ERROR"
)

type userManagementResponse struct {
	id      int
	status  statusMessage
	message string
	token   string
}
