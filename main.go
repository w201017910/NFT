package main

import (
	"fmt"
	"nft/database"
)

func main() {
	database.Init()
	fmt.Println(database.QueryCollection("1"))
}
