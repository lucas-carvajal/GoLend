package tokenManagement

import "fmt"

var channel chan Request

func Init() chan Request {
	fmt.Println("[UserManagementService Channel] Creating Channel")
	channel = make(chan Request)

	go run()

	return channel
}

func run() {
	for req := range channel {
		handleTokenRequest(req)
	}
	fmt.Println("[TokenManagementService Channel] Closed")
}

func handleTokenRequest(req Request) {
	// TODO get token values from redis, return as response to channel or send error

	req.ResponseChannel <- Response{
		OK:       true,
		Email:    "pomp@motorboys.com",
		Username: "pomp",
	}
}
