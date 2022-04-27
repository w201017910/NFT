package rounter

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"net/http"
	contract "nft/contracts/methods"
	"nft/database"
	"nft/util"
)

var client, _ = ethclient.Dial("ws://127.0.0.1:7545")

func Index2(c *gin.Context) {
	cookie, e := c.Request.Cookie("name")
	isLogin := false
	if e == nil {
		isLogin = true
	}
	if isLogin {
		c.HTML(http.StatusOK, "index-2.html", gin.H{
			"isLogin":  isLogin,
			"username": cookie.Value,
			"img":      database.QueryUser(cookie.Value).Picture,
			"balance":  database.QueryUser(cookie.Value),
		})
	} else {
		c.HTML(http.StatusOK, "index-2.html", gin.H{
			"isLogin": isLogin,
		})
	}
}
func NotFind(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", gin.H{})
}
func Collection(c *gin.Context) {
	c.HTML(http.StatusOK, "collection.html", gin.H{})
}
func Drag(c *gin.Context) {
	c.HTML(http.StatusOK, "drag.html", gin.H{})
}
func Homepage(c *gin.Context) {

	cookie, e := c.Request.Cookie("name")
	isLogin := false
	if e == nil {
		isLogin = true
	}
	if isLogin {
		//ether := big.NewInt(int64(10 * math.Pow(10, 18)))
		address := database.QueryUser(cookie.Value).Address
		privateKey := database.QueryUser(cookie.Value).Keystore
		tokenBalance := util.QueryBalance(common.HexToAddress(address), privateKey[2:])
		etherBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
		tokenId := database.QueryTokenId(address)
		if err != nil {
			log.Fatal(err)
		}
		var coverHref []string
		for _, v := range tokenId {
			tokenId_ := big.NewInt(int64(v))
			coverHref = append(coverHref, contract.TokenInfo(nil, tokenId_).Cid)
		}
		c.HTML(http.StatusOK, "homepage.html", gin.H{
			"isLogin":      isLogin,
			"username":     cookie.Value,
			"img":          database.QueryUser(cookie.Value).Picture,
			"address":      address,
			"privateKey":   privateKey,
			"balance":      tokenBalance,
			"etherBalance": etherBalance,
			"cover0":       coverHref[0],
			"cover":        coverHref[1:],
			"cid0":         coverHref[0][32:],
		})
	} else {
		c.HTML(http.StatusOK, "homepage.html", gin.H{
			"isLogin": isLogin,
		})
	}
}
func Index1(c *gin.Context) {
	cookie, e := c.Request.Cookie("name")
	isLogin := false
	if e == nil {
		isLogin = true
	}
	if isLogin {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"isLogin":  isLogin,
			"username": cookie.Value,
			"img":      database.QueryUser(cookie.Value).Picture,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"isLogin": isLogin,
		})
	}
}
func ItemDetails(c *gin.Context) {
	cid_head := "http://175.178.215.53:8080/ipfs/"
	href := c.Param("href")
	item_details := database.QueryImgByCid(cid_head + href)
	c.HTML(http.StatusOK, "item-details.html", gin.H{
		"images":      item_details.Cid,
		"tokenId":     item_details.TokenId,
		"description": item_details.Description,
		"price":       item_details.Balance,
		"is_on_sale":  item_details.IsSell,
		"type":        item_details.Type_,
	})
}
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
func Ranking(c *gin.Context) {
	cookie, e := c.Request.Cookie("name")
	isLogin := false
	if e == nil {
		isLogin = true
	}
	if isLogin {
		c.HTML(http.StatusOK, "ranking.html", gin.H{
			"isLogin":  isLogin,
			"username": cookie.Value,
			"img":      database.QueryUser(cookie.Value).Picture,
		})
	} else {
		c.HTML(http.StatusOK, "ranking.html", gin.H{
			"isLogin": isLogin,
		})
	}
}
func RegisterPage(c *gin.Context) {
	cookie, e := c.Request.Cookie("name")
	isLogin := false
	if e == nil {
		isLogin = true
	}
	if isLogin {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"isLogin":  isLogin,
			"username": cookie.Value,
			"img":      database.QueryUser(cookie.Value).Picture,
		})
	} else {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"isLogin": isLogin,
		})
	}
}
func Swap(c *gin.Context) {
	cookie, e := c.Request.Cookie("name")
	isLogin := false
	if e == nil {
		isLogin = true
	}
	if isLogin {
		c.HTML(http.StatusOK, "swap.html", gin.H{
			"isLogin":  isLogin,
			"username": cookie.Value,
			"img":      database.QueryUser(cookie.Value).Picture,
		})
	} else {
		c.HTML(http.StatusOK, "swap.html", gin.H{
			"isLogin": isLogin,
		})
	}
}
func Wallet(c *gin.Context) {
	c.HTML(http.StatusOK, "wallet.html", gin.H{})
}
