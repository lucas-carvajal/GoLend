package userManagement

type userManagementCommand string

const (
	signup userManagementCommand = "SIGNUP"
	login                        = "LOGIN"
)

type userManagementRequest struct {
	id              int
	command         userManagementCommand
	email           string
	password        string
	token           string
	responseChannel chan userManagementResponse
}
