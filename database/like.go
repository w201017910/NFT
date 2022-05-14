package database

import "fmt"

func LikeIt(tokenId int, owner string) {
	_, err := Db.Exec("INSERT INTO likes VALUES (?,?,?)", tokenId, owner, "1")
	if err != nil {
		fmt.Println(err)
	}
}
func DisLikeIt(tokenId int, owner string) {
	_, err := Db.Exec("DELETE FROM likes WHERE tokenId = ? AND owner = ? AND status = ?", tokenId, owner, "1")
	if err != nil {
		fmt.Println(err)
	}
}
func IsLiked(tokenId int, owner string) string {
	rows, err := Db.Query("SELECT status FROM likes WHERE tokenId = ? AND owner = ?", tokenId, owner)
	if err != nil {
		fmt.Println(err)
	}
	defer CloseConnection(rows)
	for rows.Next() {
		var status string
		rows.Scan(&status)
		return status
	}
	return ""
}
