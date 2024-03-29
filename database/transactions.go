package database

import (
	"fmt"
	"time"
)

type Transaction struct {
	TokenId    int
	FormID     string
	ToID       string
	TokenPrice string
	Times      int
}

func AddTransactions(tokenId int, formID string, toID string, tokenPrice int) {
	_, err := db.Exec("insert transactions  values (?,?,?,?,?)", tokenId, formID, toID, tokenPrice, time.Now().Unix())
	if err != nil {
		fmt.Println(err)
	}
}
func QueryTransaction(tokenID int) []Transaction {
	rows, err := db.Query("select * from transactions where tokenID=?", tokenID)
	if err != nil {
		fmt.Println(err)
	}
	var transactions []Transaction
scan:
	if rows.Next() {
		transaction := new(Transaction)
		rows.Scan(&transaction.TokenId, &transaction.FormID, &transaction.ToID, &transaction.TokenPrice, &transaction.Times)
		transactions = append(transactions, *transaction)
		goto scan
	}

	return transactions
}
func QueryAllTransaction() []Transaction {
	rows, err := db.Query("select * from transactions")
	if err != nil {
		fmt.Println(err)
	}
	var transactions []Transaction
scan:
	if rows.Next() {
		transaction := new(Transaction)
		rows.Scan(&transaction.TokenId, &transaction.FormID, &transaction.ToID, &transaction.TokenPrice, &transaction.Times)
		transactions = append(transactions, *transaction)
		goto scan
	}
	return transactions
}
