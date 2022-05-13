package rounter

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"net/http"
	contract "nft/contracts/methods"
	"nft/database"
	"nft/util"
	"time"
)

var client, _ = ethclient.Dial("ws://127.0.0.1:7545")

func Index2(c *gin.Context) {
	//cookie, e := c.Request.Cookie("name")
	//isLogin := false
	//if e == nil {
	//	isLogin = true
	//}
	//if isLogin {
	//	c.HTML(http.StatusOK, "index-2.html", gin.H{
	//		"isLogin":  isLogin,
	//		"username": cookie.Value,
	//		"img":      database.QueryUser(cookie.Value).Picture,
	//		"balance":  database.QueryUser(cookie.Value),
	//	})
	//} else {
	//	c.HTML(http.StatusOK, "index-2.html", gin.H{
	//		"isLogin": isLogin,
	//	})
	//}
	//search := database.QueryImgFuzzyly()

	allTrans := database.QueryAllImg()
	var tokenId []int
	if len(allTrans) > 5 {
		for _, v := range allTrans[len(allTrans)-5:] {
			tokenId = append(tokenId, v.TokenId)
		}
	} else {
		for _, v := range allTrans {
			tokenId = append(tokenId, v.TokenId)
		}
	}
	var token []*database.IMG
	var ownerName []string
	var pic []string
	for i, v := range tokenId {
		token = append(token, database.QueryImg(v))
		ownerName = append(ownerName, database.QueryUserByAddress(token[i].Owner).Name)
		pic = append(pic, database.QueryUserByAddress(token[i].Owner).Picture)
	}
	fmt.Println(pic)
	c.HTML(http.StatusOK, "index-2.html", gin.H{
		"token": token,
		"name":  ownerName,
		"pic":   pic,
	})
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
	var address string
	var tokenId []int
	var ownImg []database.IMG
	var ownImgCount int
	var coverHref []string
	var token []*database.IMG
	var count int
	var cover []string
	username := c.Query("username")
	if username != "" {
		address = database.QueryUser(username).Address
		if address != "" {
			tokenId = database.QueryTokenId(address)
			ownImg = database.QueryMyOwnImg(address)
			ownImgCount = database.QueryMyOwnImgCount(address)
			for _, v := range tokenId {
				token = append(token, database.QueryImg(v))
				count = count + 1
				tokenId_ := big.NewInt(int64(v))
				coverHref = append(coverHref, contract.TokenInfo(nil, tokenId_).Cid)
			}
			c.HTML(http.StatusOK, "homepage1.html", gin.H{
				"username":    username,
				"img":         database.QueryUser(username).Picture,
				"address":     address,
				"cover":       cover,
				"token":       token,
				"count":       count,
				"ownImgCount": ownImgCount,
				"ownImg":      ownImg,
			})
		}
	} else {
		cookie, e := c.Request.Cookie("name")
		isLogin := false
		if e == nil {
			isLogin = true
		}
		if isLogin {
			//ether := big.NewInt(int64(10 * math.Pow(10, 18)))
			address = database.QueryUser(cookie.Value).Address
			privateKey := database.QueryUser(cookie.Value).Keystore
			tokenBalance := util.QueryBalance(common.HexToAddress(address), privateKey[2:])
			etherBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
			tokenId = database.QueryTokenId(address)
			ownImg = database.QueryMyOwnImg(address)
			ownImgCount = database.QueryMyOwnImgCount(address)
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range tokenId {
				token = append(token, database.QueryImg(v))
				count = count + 1
				tokenId_ := big.NewInt(int64(v))
				coverHref = append(coverHref, contract.TokenInfo(nil, tokenId_).Cid)
			}
			judge := true
			if len(coverHref) == 0 {
				judge = false
				cover = nil
			} else {
				cover = coverHref
			}
			c.HTML(http.StatusOK, "homepage1.html", gin.H{
				"isLogin":      isLogin,
				"username":     cookie.Value,
				"img":          database.QueryUser(cookie.Value).Picture,
				"address":      address,
				"privateKey":   privateKey,
				"balance":      tokenBalance,
				"etherBalance": etherBalance,
				"cover":        cover,
				"judge":        judge,
				"token":        token,
				"count":        count,
				"ownImgCount":  ownImgCount,
				"ownImg":       ownImg,
			})
		} else {
			c.HTML(http.StatusOK, "homepage1.html", gin.H{
				"isLogin": isLogin,
			})
		}
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
	cookie, e := c.Request.Cookie("name")
	cid_head := "http://175.178.215.53:8080/ipfs/"
	href := c.Param("href")
	item_details := database.QueryImgByCid(cid_head + href)
	transactions := database.QueryTransaction(item_details.TokenId)
	var fromName []string
	var toName []string
	var format []string
	for i, v := range transactions {
		fromName = append(fromName, database.QueryUserByAddress(v.FormID).Name)
		toName = append(toName, database.QueryUserByAddress(v.ToID).Name)
		format = append(format, time.Unix(int64(transactions[i].Times), 0).Format("2006/01/02 15:04"))
	}
	isOwner := false
	if e == nil {
		userName := database.QueryUser(cookie.Value)
		if cookie.Value == userName.Name {
			isOwner = true
		}
	}
	c.HTML(http.StatusOK, "item-details.html", gin.H{
		"isOwner":         isOwner,
		"images":          item_details.Cid,
		"tokenId":         item_details.TokenId,
		"description":     item_details.Description,
		"price":           item_details.Balance,
		"is_on_sale":      item_details.IsSell,
		"type":            item_details.Type_,
		"cid":             item_details.Cid,
		"creator":         item_details.Creator,
		"owner":           item_details.Owner,
		"contractAddress": contract.Erc721Address,
		"time":            format,
		"transaction":     transactions,
		"from":            fromName,
		"to":              toName,
	})
}
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
func Ranking(c *gin.Context) {

	popularToken := database.QueryPopularTransactions()
	var count []int
	//nowTime := time.Now().Unix()
	//var week = 604800
	//var month = 2592000
	var token []database.IMG
	var ownerName []string
	var weeks []float64
	var mouth []float64
	var week1 []string
	var mouth1 []string
	for i, v := range popularToken {
		count = append(count, v.Count)
		token = append(token, *database.QueryImg(v.TokenId))
		ownerName = append(ownerName, database.QueryUserByAddress(token[i].Owner).Name)
		t := database.QueryTransaction(v.TokenId)
		weeks = append(weeks, util.Week(t))

		if weeks[i] >= 0 {
			week1 = append(week1, "text-success")
		} else {
			week1 = append(week1, "text-danger")
		}
		mouth = append(mouth, util.Mouth(t))

		if mouth[i] >= 0 {
			mouth1 = append(mouth1, "text-success")
		} else {
			mouth1 = append(mouth1, "text-danger")
		}

	}

	c.HTML(http.StatusOK, "ranking.html", gin.H{
		"token":  token,
		"name":   ownerName,
		"weeks":  weeks,
		"week1":  week1,
		"mouths": mouth,
		"mouth1": mouth1,
		"count":  count,
	})
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
func Search(c *gin.Context) {
	name := c.Param("name")
	searchResult := database.QueryImgFuzzyly(name)
	var ownerName []string
	var pic []string
	if searchResult != nil {
		for _, v := range searchResult {
			ownerName = append(ownerName, database.QueryUserByAddress(v.Owner).Name)
			pic = append(pic, database.QueryUserByAddress(v.Owner).Picture)
		}
		c.HTML(http.StatusOK, "search.html", gin.H{
			"result": searchResult,
			"pic":    pic,
			"name":   ownerName,
		})
	} else {
		c.JSON(200, gin.H{
			"err": "nil",
		})
	}
}
func Homepage1(c *gin.Context) {
	//c.HTML(http.StatusOK, "search.html", gin.H{})
}
