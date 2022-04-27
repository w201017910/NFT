package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"html/template"
	"nft/contracts/methods"
	"nft/database"
	"nft/rounter"
)

func main() {
	database.Init()
	clients, _ := ethclient.Dial("ws://127.0.0.1:7545")
	methods.InitAll(clients)
	defer clients.Close()
	Start("localhost:8080")
}

func silice(href string) string {
	return href[32:]
}
func Start(addr string) (err error) {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"silice": silice,
	})
	r.Static("/frontend", "./frontend")
	r.LoadHTMLGlob("./template/*")
	r.GET("/index-2.html", rounter.Index2)
	r.GET("/404.html", rounter.NotFind)
	r.GET("/collection.html", rounter.Collection)
	r.GET("/drag.html", rounter.Drag)
	r.GET("/homepage.html", rounter.Homepage)
	r.GET("/item-details.html/:href", rounter.ItemDetails)
	r.GET("/login.html", rounter.LoginPage)
	r.GET("/ranking.html", rounter.Ranking)
	r.GET("/register.html", rounter.RegisterPage)
	r.GET("/swap.html", rounter.Swap)
	r.GET("/wallet.html", rounter.Wallet)
	r.GET("/", rounter.Index1)
	r.POST("/register", rounter.Register)
	r.POST("/login", rounter.Login)
	r.POST("/logout", rounter.LoginOut)
	r.POST("/upLoad", rounter.ReadFile)

	err = r.Run(addr)
	return err
}
