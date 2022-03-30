package database

import "fmt"

type Collection struct {
	TokenID int
	UserID  string
}

func CollectionCount(tokenID int) int {
	rows, err := db.Query("SELECT COUNT(tokenID) FROM collection where tokenID=?", tokenID)
	if err != nil {
		fmt.Println(err)
	}
	var s int
	if rows.Next() {

		rows.Scan(&s)

	}
	return s
}
func InsertCollection(tokenID int, userID string) {
	_, err := db.Exec("insert collection values (?,?)", tokenID, userID)
	if err != nil {
		fmt.Println(err)
	}
}
func QueryCollection(userID string) []Collection {
	rows, err := db.Query("select * from collection where userID=?", userID)
	if err != nil {
		fmt.Println(err)
	}
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
