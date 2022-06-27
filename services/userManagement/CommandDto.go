package userManagement

type userManagementCommand string

const (
	signup userManagementCommand = "SIGNUP"
	login                        = "LOGIN"
)

type userManagementDto struct {
	command  userManagementCommand
	id       int
	email    string
	password string
}
