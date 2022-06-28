package userManagement

type userManagementCommand string

const (
	signup userManagementCommand = "SIGNUP"
	login                        = "LOGIN"
)

type Request struct {
	Id              int
	Command         userManagementCommand
	Email           string
	Password        string
	Token           string
	ResponseChannel chan Response
}
