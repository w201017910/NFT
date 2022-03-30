package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type IMG struct {
	tokenId     int
	owner       string
	creator     string
	cid         string
	type_       string
	description string
	isSell      bool
	balance     int
	thumbsUp    int
	browseCount int
}

func CreateImg(tokenId int, owner string, creator string, cid string, _type string, description string) {
	_, err := db.Exec("insert img  values (?,?,?,?,?,?,false,0,0,0)", tokenId, owner, creator, cid, _type, description)
	if err != nil {
		fmt.Println(err)

	}
}
func QueryImg(tokenID int) *IMG {
	rows, err := db.Query("select * from img where tokenId=?", tokenID)
	if err != nil {
		fmt.Println(err)
	}
	if rows.Next() {
		img := new(IMG)
		rows.Scan(&img.tokenId, &img.owner, &img.creator, &img.cid, &img.type_, &img.description, &img.isSell, &img.balance, &img.thumbsUp, &img.browseCount)
		return img
	}
	return nil
}
func QueryAllImg() []IMG {
	rows, err := db.Query("select * from img")
	if err != nil {
		fmt.Println(err)
	}
	var images []IMG
scan:
	if rows.Next() {
		img := new(IMG)
		rows.Scan(&img.tokenId, &img.owner, &img.creator, &img.cid, &img.type_, &img.description, &img.isSell, &img.balance, &img.thumbsUp, &img.browseCount)
		images = append(images, *img)
		goto scan
	}
	return images
}
func ChangeOwner(tokenId int, owner string) {
	_, err := db.Exec("UPDATE `img` SET `owner` = ? WHERE `tokenId` = ?", owner, tokenId)
	if err != nil {
		fmt.Println(err)
	}
}
func ChangeSell(tokenId int, isSell bool, balance int) {
	_, err := db.Exec("UPDATE `img` SET isSell=?,`balance` = ? WHERE `tokenId` = ?", isSell, balance, tokenId)
	if err != nil {
		fmt.Println(err)
	}
}
func ThumbsUpAdd(tokenId int) {
	_, err := db.Exec("UPDATE `img` SET thumbsUp=thumbsUp+1 WHERE `tokenId` = ?", tokenId)
	if err != nil {
		fmt.Println(err)
	}
}
func BrowseCountAdd(tokenId int) {
	_, err := db.Exec("UPDATE `img` SET browseCount=browseCount+1 WHERE `tokenId` = ?", tokenId)
	if err != nil {
		fmt.Println(err)
	}
}
