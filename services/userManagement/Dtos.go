package userManagement

type userManagementCommand string

const (
	signup userManagementCommand = "SIGNUP"
	login                        = "LOGIN"
)

type Request struct {
	Command         userManagementCommand
	Username        string
	Email           string
	Password        string
	Token           string
	ResponseChannel chan Response
}

type Response struct {
	Status  int
	Message string
	Token   string
}
