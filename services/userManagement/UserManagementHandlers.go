package userManagement

import (
	"fmt"
	"net/http"
)

func handleSignup(req Request) {
	users := getUsersByEmail(req.Email)
	if len(users) != 0 {
		req.ResponseChannel <- Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("user with email '%s' already exists", req.Email),
			Token:   "",
		}
		return
	}
	users = getUsersByUsername(req.Username)
	if len(users) != 0 {
		req.ResponseChannel <- Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("user with username '%s' already exists", req.Username),
			Token:   "",
		}
		return
	}
	success := insertUser(user{
		username: req.Username,
		email:    req.Email,
		password: req.Password,
	})
	if !success {
		req.ResponseChannel <- Response{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("user with email '%s' could not be saved", req.Email),
			Token:   "",
		}
		return
	}
	req.ResponseChannel <- Response{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("user with email '%s' successfully created", req.Email),
		Token:   "",
	}
}

func handleLogin(req Request) {
	users := getUsersByEmail(req.Email)
	if len(users) != 1 {
		req.ResponseChannel <- Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("user with email '%s' does not exist", req.Email),
			Token:   "",
		}
		return
	}
	users = getUsersByUsername(req.Username)
	if len(users) != 1 {
		req.ResponseChannel <- Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("user with username '%s' does not exist", req.Username),
			Token:   "",
		}
		return
	}
	success := users[0].password == req.Password

	if !success {
		req.ResponseChannel <- Response{
			Status:  http.StatusUnauthorized,
			Message: fmt.Sprintf("Password for user '%s' does not match", req.Username),
			Token:   "",
		}
		return
	}
	// TODO GB-7: generate auth token and save to redis by using TokenManagementService
	req.ResponseChannel <- Response{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Login for user '%s' successful", req.Username),
		Token:   "tbd",
	}
}
