package main

import (
	"fmt"
	"gobank.com/services/userManagement"
	"gobank.com/util"
	"net/http"
	"strings"
)

func handleRequest(
	w http.ResponseWriter,
	r *http.Request,
	userMgmChan chan userManagement.Request,
) {
	urlParts := strings.Split(r.URL.String(), "/")
	urlParts = util.Filter(urlParts, func(s string) bool {
		return s != ""
	})
	switch {
	case len(urlParts) == 0:
		// default route "/"
		doSomething(w, r)
		return
	case urlParts[0] == "signup":
		// signup route "/signup"
		doSomething(w, r)
		userMgmChan <- userManagement.Request{
			Id:      1,
			Command: "SIGNUP",
		}
		return
	case urlParts[0] == "login":
		// login route "/login"
		doSomething(w, r)
		userMgmChan <- userManagement.Request{
			Id:      1,
			Command: "LOGIN",
		}
		return
	case urlParts[0] == "claim" && len(urlParts) == 1:
		// claim route "/claim" -> file a claim
		doSomething(w, r)
		return
	case urlParts[0] == "claim" && len(urlParts) > 1:
		switch {
		case len(urlParts) == 2:
			// claim id route "/claim/1424" -> retrieve details for id=1424
			doSomething(w, r)
			return
		case urlParts[2] == "approve" && len(urlParts) == 3:
			// approve claim route "/claim/1424/approve"
			doSomething(w, r)
			return
		case urlParts[2] == "deny" && len(urlParts) == 3:
			// deny claim route "/claim/1424/deny"
			doSomething(w, r)
			return
		case urlParts[2] == "settleRequest" && len(urlParts) == 3:
			// settle claim request route "/claim/1424/settleRequest"
			doSomething(w, r)
			return
		case urlParts[2] == "settle" && len(urlParts) == 3:
			// settle claim route "/claim/1424/settle"
			doSomething(w, r)
			return
		case urlParts[2] == "deadline" && len(urlParts) == 3:
			// set deadline for claim route "/claim/1424/deadline"
			doSomething(w, r)
			return
		case urlParts[2] == "interest" && len(urlParts) == 3:
			// set interest for claim route "/claim/1424/interest"
			doSomething(w, r)
			return
		default:
			// TODO return error message
			return
		}
	default:
		// TODO return error message
		return
	}
}

func doSomething(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Call to %s\n", r.URL.String())
}
