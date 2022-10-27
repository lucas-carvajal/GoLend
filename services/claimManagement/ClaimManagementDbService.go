package claimManagement

import (
	"fmt"
	"strconv"
	"time"
)

type Loan struct {
	Id          int
	Title       string
	Description string
	FromUser    string
	ToUser      string
	Date        string
	Amount      int
	Interest    int
	Deadline    string
	Status      string
}

func insertNewLoan(loan Loan) bool {
	currentTime := time.Now()
	today := currentTime.Format("2006-01-02")
	_, err := db.Exec(
		"INSERT INTO loans "+
			"(title, description, fromUser, toUser, date, amount, interest, deadline, status) "+
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		loan.Title, loan.Description, loan.FromUser, loan.ToUser, today,
		loan.Amount, loan.Interest, loan.Deadline, PENDING)
	fmt.Println("INSERTING NEW LOAN " + strconv.FormatBool(err != nil))
	if err != nil {
		return false
	}
	return true
}

func updatedLoan(id, column, value string) bool {
	_, err := db.Exec(
		"UPDATE loans SET ? = ? WHERE id = ?;",
		column, value, id)
	if err != nil {
		return false
	}
	return true
}
