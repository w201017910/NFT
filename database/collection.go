package database

import "fmt"

type Collection struct {
	TokenID int
	UserID  string
}

func CollectionCount(tokenID int) int {
	rows, err := Db.Query("SELECT COUNT(tokenID) FROM collection where tokenID=?", tokenID)
	if err != nil {
		fmt.Println(err)
	}
	defer CloseConnection(rows)
	var s int
	if rows.Next() {

		rows.Scan(&s)

	}
	return s
}
func InsertCollection(tokenID int, userID string) {
	_, err := Db.Exec("insert collection values (?,?)", tokenID, userID)
	if err != nil {
		fmt.Println(err)
	}
}
func QueryCollection(userID string) []Collection {
	rows, err := Db.Query("select * from collection where userID=?", userID)
	if err != nil {
		fmt.Println(err)
	}
	defer CloseConnection(rows)
	var collections []Collection
scan:
	if rows.Next() {
		collection := new(Collection)
		rows.Scan(&collection.TokenID, &collection.UserID)
		collections = append(collections, *collection)
		goto scan
	}
	return collections
}
