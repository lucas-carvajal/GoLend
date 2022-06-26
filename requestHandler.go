package main

import (
	"fmt"
	"gobank.com/util"
	"net/http"
	"strings"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	//TODO remove trailing "/"
	switch {
	case url == "/":
		doSomething(w, r)
		return
	case url == "/signup":
		doSomething(w, r)
		return
	case url == "/login":
		doSomething(w, r)
		return
	case url == "/claim":
		doSomething(w, r)
		return
	case strings.Contains(url, "/claim/"):
		switch {
		case strings.Contains(url, "/approve"):
			doSomething(w, r)
			return
		case strings.Contains(url, "/deny"):
			doSomething(w, r)
			return
		case strings.Contains(url, "/settleRequest"):
			doSomething(w, r)
			return
		case strings.Contains(url, "/settle"):
			doSomething(w, r)
			return
		case strings.Contains(url, "/deadline"):
			doSomething(w, r)
			return
		case strings.Contains(url, "/interest"):
			doSomething(w, r)
			return
		case !strings.Contains(util.RemoveFirstUrlPathElement(url), "/"):
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
	fmt.Printf("Call to %s", r.URL.String())
}
