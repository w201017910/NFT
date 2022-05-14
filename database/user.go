package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Init() {
	conSte := "root:12345678@tcp(127.0.0.1:3306)/nft"
	var err error
	Db, err = sql.Open("mysql", conSte)
	if err != nil {
		fmt.Print(err)
	}
}
func InsertUser(name string, password string, address string, email string, picture string, keyStore string, mneonic string) bool {
	_, err := Db.Exec("insert user values (?,?,?,?,?,?,?)", name, password, address, picture, email, keyStore, mneonic)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func ChangeImg(name string, img string) {
	_, err := Db.Exec("UPDATE `user` SET `picture` = ? WHERE `uname` = ?", img, name)
	if err != nil {
		fmt.Println(err)
	}
}
func ChangeAddress(name string, address string, keystore string, mneonic string) {
	_, err := Db.Exec("UPDATE `user` SET `address`=?,keystore = ?,mneonic=? WHERE `uname` = ?", address, keystore, mneonic, name)
	if err != nil {
		fmt.Println(err)
	}
}
func SignIn(name string, password string) *Person {
	rows, err := Db.Query("select * from user where uname=? and password=?", name, password)
	if err != nil {
		fmt.Println(err)
	}
	defer CloseConnection(rows)
	if rows.Next() {
		person := new(Person)
		rows.Scan(&person.Name, &person.Password, &person.Address, &person.Picture, &person.Email, &person.Keystore, &person.Mneonic)
		return person
	}
	return nil
}
func QueryUser(name string) *Person {

	rows, err := Db.Query("select * from user where uname=?", name)
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}

	if rows.Next() {
		person := new(Person)
		rows.Scan(&person.Name, &person.Password, &person.Address, &person.Picture, &person.Email, &person.Keystore, &person.Mneonic)
		return person
	}
	return nil
}

func QueryUserByAddress(address string) *Person {

	rows, err := Db.Query("select * from user where address=?", address)
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}

	if rows.Next() {
		person := new(Person)
		rows.Scan(&person.Name, &person.Password, &person.Address, &person.Picture, &person.Email, &person.Keystore, &person.Mneonic)
		return person
	}
	return nil
}

type Person struct {
	Name     string
	Password string
	Address  string
	Picture  string
	Email    string
	Keystore string
	Mneonic  string
}
