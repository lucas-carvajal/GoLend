package userManagement

import (
	"fmt"
	"net/http"
)

func handleSignup(dto Request) {
	users := getUsersByEmail(dto.Email)
	if len(users) != 0 {
		dto.ResponseChannel <- Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("user with email '%s' already exists", dto.Email),
			Token:   "",
		}
		return
	}
	users = getUsersByUsername(dto.Username)
	if len(users) != 0 {
		dto.ResponseChannel <- Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("user with username '%s' already exists", dto.Username),
			Token:   "",
		}
		return
	}
	success := insertUser(user{
		username: dto.Username,
		email:    dto.Email,
		password: dto.Password,
	})
	if !success {
		dto.ResponseChannel <- Response{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("user with email '%s' could not be saved", dto.Email),
			Token:   "",
		}
		return
	}
	dto.ResponseChannel <- Response{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("user with email '%s' successfully created", dto.Email),
		Token:   "",
	}
}

func handleLogin(dto Request) {
	users := getUsersByEmail(dto.Email)
	if len(users) != 1 {
		dto.ResponseChannel <- Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("user with email '%s' does not exist", dto.Email),
			Token:   "",
		}
		return
	}
	users = getUsersByUsername(dto.Username)
	if len(users) != 1 {
		dto.ResponseChannel <- Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("user with username '%s' does not exist", dto.Username),
			Token:   "",
		}
		return
	}
	sucess := users[0].password == dto.Password

	if !sucess {
		dto.ResponseChannel <- Response{
			Status:  http.StatusUnauthorized,
			Message: fmt.Sprintf("Password for user '%s' does not match", dto.Username),
			Token:   "",
		}
		return
	}
	// TODO generate auth token and save to redis by using TokenManagementService
	dto.ResponseChannel <- Response{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Login for user '%s' successful", dto.Username),
		Token:   "tbd",
	}
}
