package userManagement

import "fmt"

var channel chan string

func Init() chan string {
	channel = make(chan string)

	go run()

	return channel
}

func run() {
	for in := range channel {
		fmt.Println("[Channel] Message received: ", in)
	}
	fmt.Println("[Channel] Closed")
}
