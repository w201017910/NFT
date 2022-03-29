package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Init() {
	conSte := "root:123456@tcp(127.0.0.1:3306)/nft"
	var err error
	Db, err = sql.Open("mysql", conSte)
	if err != nil {
		fmt.Print(err)
	}

}
func Insert_user(name string, password string, email string) bool {
	_, err := Db.Exec("insert user (uname,password,email) values (?,?)", name, password, email)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func Change_img(name string, img string) {
	_, err := Db.Exec("UPDATE `user` SET `picture` = ? WHERE `uname` = ?", name, img)
	if err != nil {
		fmt.Println(err)
	}
}
func Change_address(name string, address string, keystore string) {
	_, err := Db.Exec("UPDATE `user` SET `address`,keystore = ? WHERE `uname` = ?", address, keystore, name)
	if err != nil {
		fmt.Println(err)
	}
}
func Sign_in(name string, password string) *Person {
	rows, err := Db.Query("select * from user where uname=?,password=?", name, password)
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
func Query(name string) *Person {

	rows, err := Db.Query("select * from user where uname=?", name)
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

type Person struct {
	name     string
	password string
	address  string
	picture  string
	email    string
	keystore string
}
