package rounter

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"math/big"
	"net/http"
	contract "nft/contracts/methods"
	"nft/database"
	"nft/util"
	"strconv"
	"time"
)

var client, _ = ethclient.Dial("ws://127.0.0.1:7545")

func Index2(c *gin.Context) {
	cookie, e := c.Request.Cookie("name")
	isLogin := false
	var username string
	var img string
	if e == nil {
		isLogin = true
		username = cookie.Value
		img = database.QueryUser(cookie.Value).Picture
	}

	allTrans := database.QueryAllImg()
	var tokenId []int
	for _, v := range allTrans {
		if len(tokenId) >= 10 {
			break
		}

		tokenId = append(tokenId, v.TokenId)

	}
	var token []*database.IMG
	var ownerName []string
	var pic []string
	for i, v := range tokenId {
		token = append(token, database.QueryImg(v))
		ownerName = append(ownerName, database.QueryUserByAddress(token[i].Owner).Name)
		pic = append(pic, database.QueryUserByAddress(token[i].Owner).Picture)
	}
	p := database.QueryTypeImg("影视动漫")
	var users []database.Person
	for _, v := range p {
		users = append(users, *database.QueryUserByAddress(v.Owner))
	}
	p1 := database.QueryTypeImg("人文艺术")
	var user1 []database.Person
	for _, v := range p1 {
		user1 = append(users, *database.QueryUserByAddress(v.Owner))
	}
	p2 := database.QueryTypeImg("自然百态")
	var user2 []database.Person
	for _, v := range p2 {
		user2 = append(users, *database.QueryUserByAddress(v.Owner))
	}
	p3 := database.QueryTypeImg("万物生灵")
	var user3 []database.Person
	for _, v := range p3 {
		user3 = append(users, *database.QueryUserByAddress(v.Owner))
	}
	p4 := database.QueryTypeImg("其他")
	var user4 []database.Person
	for _, v := range p4 {
		user4 = append(users, *database.QueryUserByAddress(v.Owner))
	}
	c.HTML(http.StatusOK, "index-2.html", gin.H{
		"img":      img,
		"username": username,
		"isLogin":  isLogin,
		"token":    token,
		"name":     ownerName,
		"pic":      pic,
		"tokens":   p,
		"users":    users,
		"token1":   p1,
		"user1":    user1,
		"token2":   p2,
		"user2":    user2,
		"token3":   p3,
		"user3":    user3,
		"token4":   p4,
		"user4":    user4,
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
	var myLikeImg []database.IMG
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
			myLikeImg = database.QueryMyLikeImg(address)
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
				"myLikeImg":   myLikeImg,
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
			tokenBalance := util.QueryBalance(common.HexToAddress(address))
			etherBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
			fmt.Println(tokenBalance, etherBalance)
			tokenId = database.QueryTokenId(address)
			ownImg = database.QueryMyOwnImg(address)
			ownImgCount = database.QueryMyOwnImgCount(address)
			myLikeImg = database.QueryMyLikeImg(address)
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range tokenId {
				token = append(token, database.QueryImg(v))
				count = count + 1
				tokenId_ := big.NewInt(int64(v))
				coverHref = append(coverHref, contract.TokenInfo(nil, tokenId_).Cid)
			}

			if len(coverHref) == 0 {

				cover = nil
			} else {
				cover = coverHref
			}

			c.HTML(http.StatusOK, "homepage1.html", gin.H{
				"isLogin":      isLogin,
				"username":     cookie.Value,
				"img":          database.QueryUser(cookie.Value).Picture,
				"address":      address,
				"balance":      tokenBalance,
				"etherBalance": etherBalance,
				"cover":        cover,

				"token":       token,
				"count":       count,
				"ownImgCount": ownImgCount,
				"ownImg":      ownImg,
				"myLikeImg":   myLikeImg,
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
	var username string
	var img string
	if e == nil {
		isLogin = true
		username = cookie.Value
		img = database.QueryUser(cookie.Value).Picture
	}
	img1 := database.QueryImg(7)
	img2 := database.QueryImg(32)
	img3 := database.QueryImg(33)
	img4 := database.QueryAllImg()
	var img5 []database.IMG
	for i, v := range img4 {
		if i >= 8 {
			break
		}
		img5 = append(img5, v)
	}
	var user5 []database.Person
	for _, v := range img5 {
		user5 = append(user5, *database.QueryUserByAddress(v.Owner))
	}
	var count []int
	for _, v := range img5 {
		count = append(count, database.LikeCount(v.TokenId))
	}
	var min []int
	for _, v := range img5 {
		min = append(min, database.MinTransaction(v.TokenId))
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"isLogin":  isLogin,
		"username": username,
		"img":      img,
		"user1":    database.QueryUserByAddress(img1.Owner),
		"user2":    database.QueryUserByAddress(img2.Owner),
		"user3":    database.QueryUserByAddress(img3.Owner),
		"balance1": img1.Balance,
		"balance2": img2.Balance,
		"balance3": img3.Balance,
		"img5":     img5,
		"user5":    user5,
		"count":    count,
		"min":      min,
	})
}
func ItemDetails(c *gin.Context) {
	cookie, e := c.Request.Cookie("name")
	isLogin := false
	var img string
	var username string
	if e == nil {
		isLogin = true
		img = database.QueryUser(cookie.Value).Picture
		username = cookie.Value
	}
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
	isLiked := false
	if e == nil {
		userName := database.QueryUser(username)
		isLike := database.IsLiked(item_details.TokenId, userName.Address)
		if isLike == "" {
			isLiked = true
		}

		if database.QueryUserByAddress(item_details.Owner).Name == userName.Name {
			isOwner = true
		}
	}
	c.HTML(http.StatusOK, "item-details.html", gin.H{
		"username":        username,
		"img":             img,
		"isLogin":         isLogin,
		"isOwner":         isOwner,
		"isLiked":         isLiked,
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
		"likeCount":       item_details.ThumbsUp,
	})
}
func Like(c *gin.Context) {
	cookie, e := c.Request.Cookie("name")
	tokenId, _ := strconv.Atoi(c.PostForm("id"))
	if e == nil {
		address := database.QueryUser(cookie.Value).Address
		isLiked := database.IsLiked(tokenId, address)
		if isLiked == "1" {
			database.LikeIt(tokenId, address)
		} else if isLiked == "" {
			database.DisLikeIt(tokenId, address)
		}
	} else {
		c.JSON(200, gin.H{
			"message": "请先登录！",
		})
	}
}
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
func Ranking(c *gin.Context) {
	cookie, e := c.Request.Cookie("name")
	isLogin := false
	var img string
	var username string
	if e == nil {
		isLogin = true
		img = database.QueryUser(cookie.Value).Picture
		username = cookie.Value
	}
	popularToken := database.QueryPopularTransactions()
	var count []int
	//nowTime := time.Now().Unix()
	//var week = 604800
	//var month = 2592000
	var token []database.IMG
	var ownerName []string
	var weeks []string
	var mouth []string
	var week1 []string
	var mouth1 []string
	for i, v := range popularToken {
		count = append(count, v.Count)
		token = append(token, *database.QueryImg(v.TokenId))
		ownerName = append(ownerName, database.QueryUserByAddress(token[i].Owner).Name)
		t := database.QueryTransaction(v.TokenId)
		f := math.Trunc(util.Week(t)*1e2) * 1e-2
		weeks = append(weeks, strconv.FormatFloat(f, 'f', -1, 64))

		if f >= 0 {
			week1 = append(week1, "text-success")
		} else {
			week1 = append(week1, "text-danger")
		}
		f = math.Trunc(util.Mouth(t)*1e2) * 1e-2
		mouth = append(mouth, strconv.FormatFloat(f, 'f', -1, 64))

		if f >= 0 {
			mouth1 = append(mouth1, "text-success")
		} else {
			mouth1 = append(mouth1, "text-danger")
		}

	}

	c.HTML(http.StatusOK, "ranking.html", gin.H{
		"isLogin":  isLogin,
		"img":      img,
		"username": username,
		"token":    token,
		"name":     ownerName,
		"weeks":    weeks,
		"week1":    week1,
		"mouths":   mouth,
		"mouth1":   mouth1,
		"count":    count,
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
	searchResult := database.QueryImgFuzzily(name)
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
func ExportPrivateKey(c *gin.Context) {
	auth := c.PostForm("auth")
	cookie, e := c.Request.Cookie("name")
	if e == nil {
		keyFile := database.QueryUser(cookie.Value).Keystore
		//passwd := database.QueryUser(cookie.Value).Password
		privateKey, err := util.ExportPrivateKey(keyFile, auth)
		if err != nil {
			c.JSON(200, gin.H{
				"err": "导出失败，请检查账号正常或密码正确！",
			})
		} else {
			c.JSON(200, gin.H{
				"key": privateKey,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "未登录！",
		})
	}
}
