package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type IMG struct {
	TokenId     int
	Name        string
	Owner       string
	Creator     string
	Cid         string
	Type_       string
	Description string
	IsSell      bool
	Balance     int
	ThumbsUp    int
	BrowseCount int
}

func CreateImg(tokenId int, name string, owner string, creator string, cid string, _type string, description string) {
	_, err := Db.Exec("insert img  values (?,?,?,?,?,?,?,false,0,0,0)", tokenId, name, owner, creator, cid, _type, description)
	if err != nil {
		fmt.Println(err)

	}
}
func QueryTokenId(address string) []int {
	rows, err := Db.Query("select `tokenId` from img where owner =?", address)
	if err != nil {
		log.Fatal(err)
	}
	defer CloseConnection(rows)
	var tokenIds []int
scan:
	if rows.Next() {
		img := new(IMG)
		rows.Scan(&img.TokenId)
		tokenIds = append(tokenIds, img.TokenId)
		goto scan
	}
	return tokenIds

}
func QueryImg(tokenID int) *IMG {
	rows, err := Db.Query("select * from img where tokenId=?", tokenID)
	if err != nil {
		fmt.Println(err)
	}
	defer CloseConnection(rows)
	for rows.Next() {
		img := new(IMG)
		rows.Scan(&img.TokenId, &img.Name, &img.Owner, &img.Creator, &img.Cid, &img.Type_, &img.Description, &img.IsSell, &img.Balance, &img.ThumbsUp, &img.BrowseCount)
		return img
	}
	return nil
}
func QueryImgFuzzily(content string) []IMG {
	rows, err := Db.Query("select * from img where name like concat('%',?,'%') or description like concat('%',?,'%') ", content, content)
	if err != nil {
		fmt.Println(err)
	}
	defer CloseConnection(rows)
	var images []IMG
	for rows.Next() {
		img := new(IMG)
		rows.Scan(&img.TokenId, &img.Name, &img.Owner, &img.Creator, &img.Cid, &img.Type_, &img.Description, &img.IsSell, &img.Balance, &img.ThumbsUp, &img.BrowseCount)
		images = append(images, *img)
	}
	return images
}
func QueryImgByCid(cid string) *IMG {
	rows, err := Db.Query("select * from img where cid=?", cid)
	if err != nil {
		fmt.Println(err)
	}
	defer CloseConnection(rows)
	if rows.Next() {
		img := new(IMG)
		rows.Scan(&img.TokenId, &img.Name, &img.Owner, &img.Creator, &img.Cid, &img.Type_, &img.Description, &img.IsSell, &img.Balance, &img.ThumbsUp, &img.BrowseCount)
		return img
	}
	return nil
}

func QueryMyOwnImg(address string) []IMG {
	rows, err := Db.Query("select * from img where creator = ?", address)
	if err != nil {
		fmt.Println(err)
	}
	defer CloseConnection(rows)
	var images []IMG
scan:
	if rows.Next() {
		img := new(IMG)
		rows.Scan(&img.TokenId, &img.Name, &img.Owner, &img.Creator, &img.Cid, &img.Type_, &img.Description, &img.IsSell, &img.Balance, &img.ThumbsUp, &img.BrowseCount)
		images = append(images, *img)
		goto scan
	}
	return images
}
func QueryMyLikeImg(address string) []IMG {
	rows, err := Db.Query("select img.* from img  left join likes on img.tokenId = likes.tokenId where likes.owner = ?", address)
	if err != nil {
		fmt.Println(err)
	}
	defer CloseConnection(rows)
	var images []IMG
scan:
	if rows.Next() {
		img := new(IMG)
		rows.Scan(&img.TokenId, &img.Name, &img.Owner, &img.Creator, &img.Cid, &img.Type_, &img.Description, &img.IsSell, &img.Balance, &img.ThumbsUp, &img.BrowseCount)
		images = append(images, *img)
		goto scan
	}
	return images
}
func QueryMyOwnImgCount(address string) int {
	rows, _ := Db.Query("select count(tokenId) from img where creator = ?", address)

	defer CloseConnection(rows)
	for rows.Next() {
		var count int
		rows.Scan(&count)
		return count
	}
	return 0
}
func QueryAllImg() []IMG {
	rows, err := Db.Query("select * from img")
	if err != nil {
		fmt.Println(err)
	}
	defer CloseConnection(rows)
	var images []IMG
scan:
	if rows.Next() {
		img := new(IMG)
		rows.Scan(&img.TokenId, &img.Name, &img.Owner, &img.Creator, &img.Cid, &img.Type_, &img.Description, &img.IsSell, &img.Balance, &img.ThumbsUp, &img.BrowseCount)
		images = append(images, *img)
		goto scan
	}
	return images
}
func ChangeOwner(tokenId int, owner string) {
	_, err := Db.Exec("UPDATE `img` SET `owner` = ?,`isSell` = ?,`balance` = ? WHERE `tokenId` = ?", owner, false, 0, tokenId)
	if err != nil {
		fmt.Println(err)
	}
}
func ChangeSell(tokenId int, isSell bool, balance int) {
	_, err := Db.Exec("UPDATE `img` SET isSell=?,`balance` = ? WHERE `tokenId` = ?", isSell, balance, tokenId)
	if err != nil {
		fmt.Println(err)
	}
}
func DeleteImg(tokenId int, owner string) {
	_, err := Db.Exec("DELETE FROM img where tokenId = ? AND owner = ?", tokenId, owner)
	if err != nil {
		fmt.Println(err)
	}
}
func ThumbsUpAdd(tokenId int) {
	_, err := Db.Exec("UPDATE `img` SET thumbsUp=thumbsUp+1 WHERE `tokenId` = ?", tokenId)
	if err != nil {
		fmt.Println(err)
	}
}
func BrowseCountAdd(tokenId int) {
	_, err := Db.Exec("UPDATE `img` SET browseCount=browseCount+1 WHERE `tokenId` = ?", tokenId)
	if err != nil {
		fmt.Println(err)
	}
}
