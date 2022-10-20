package claimManagement

import "time"

type loan struct {
	id          int
	title       string
	description string
	fromUser    string
	toUser      string
	date        string
	amount      int
	interest    int
	deadline    string
	status      string
}

func insertNewLoan(loan loan) bool {
	currentTime := time.Now()
	today := currentTime.Format("2006-01-02")
	_, err := db.Exec(
		"INSERT INTO loans "+
			"(title, description, fromUser, toUser, date, amount, interest, deadline, status) "+
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		loan.title, loan.description, loan.fromUser, loan.toUser, today,
		loan.amount, loan.interest, loan.deadline, PENDING)
	if err != nil {
		return false
	}
	return true
}
