package main

import "fmt"

func main() {
	fmt.Println("Hello!")

	// test DB stuff
	connectToDB()
	users, err := loadAllUsers()
	if err == nil {
		fmt.Println("USERS:")
		for _, user := range users {
			fmt.Printf("Username: %s\n", user.username)
			fmt.Printf("Password: %s\n", user.password)
			fmt.Printf("Email: %s\n", user.email)
			fmt.Println("===")
		}
	}
}
