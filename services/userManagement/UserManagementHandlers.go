package userManagement

import "fmt"

func handleSignup(dto Request) {
	users := getUsersByEmail(dto.Email)
	if len(users) != 0 {
		dto.ResponseChannel <- Response{
			Id:      dto.Id,
			Status:  "ERROR",
			Message: fmt.Sprintf("user with email '%s' already exists", dto.Email),
			Token:   "",
		}
	}
	// TODO insert new user into db

	// TODO return response

}

func handleLogin(dto Request) {
	//TODO
	// check if user with email exists
	// check if password matches
	// return response with auth token
}
