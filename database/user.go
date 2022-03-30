package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() {
	conSte := "root:123456@tcp(127.0.0.1:3306)/nft"
	var err error
	db, err = sql.Open("mysql", conSte)
	if err != nil {
		fmt.Print(err)
	}

}
func InsertUser(name string, password string, email string) bool {
	_, err := db.Exec("insert user (uname,password,email) values (?,?,?)", name, password, email)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func ChangeImg(name string, img string) {
	_, err := db.Exec("UPDATE `user` SET `picture` = ? WHERE `uname` = ?", img, name)
	if err != nil {
		fmt.Println(err)
	}
}
func ChangeAddress(name string, address string, keystore string) {
	_, err := db.Exec("UPDATE `user` SET `address`=?,keystore = ? WHERE `uname` = ?", address, keystore, name)
	if err != nil {
		fmt.Println(err)
	}
}
func SignIn(name string, password string) *Person {
	rows, err := db.Query("select * from user where uname=? and password=?", name, password)
	if err != nil {
		fmt.Println(err)
	}

	if rows.Next() {
		person := new(Person)
		rows.Scan(&person.name, &person.password, &person.address, person.picture, person.email, person.keystore)
		return person
	}
	return nil
}
func QueryUser(name string) *Person {

	rows, err := db.Query("select * from user where uname=?", name)
	if err != nil {
		fmt.Println(err)
	}

	if rows.Next() {
		person := new(Person)
		rows.Scan(&person.name, &person.password, &person.address, &person.picture, &person.email, &person.keystore)
		return person
	}
	return nil
}

type Person struct {
	name     string
	password string
	address  string
	picture  string
	email    string
	keystore string
}
