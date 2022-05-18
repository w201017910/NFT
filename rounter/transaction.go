package rounter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	contract "nft/contracts/methods"
	"nft/database"
	"nft/util"
	"strconv"
)

func Transaction(c *gin.Context) {
	tokenId := c.PostForm("id")
	price := c.PostForm("price")
	owner := c.PostForm("owner")
	cookie, e := c.Request.Cookie("name")
	if e == nil {
		address := database.QueryUser(cookie.Value).Address
		if owner != address {
			c.JSON(http.StatusOK, gin.H{
				"err": "暂无权限出价！",
			})
		} else {
			p, _ := strconv.ParseInt(price, 10, 64)
			t, _ := strconv.ParseInt(tokenId, 10, 64)
			//privateKey, _ := crypto.HexToECDSA(database.QueryUser(cookie.Value).Keystore[2:])
			//opts := bind.NewKeyedTransactor(privateKey)
			opts, _ := util.BindOptsByKeystore(database.QueryUser(cookie.Value).Keystore, database.QueryUser(cookie.Value).Password)
			_, _ = contract.ERC721_Approve(opts, contract.TransactionAddress, big.NewInt(t))
			_, err := contract.Trans_Sell(opts, big.NewInt(t), big.NewInt(p))
			if err != nil {
				fmt.Println("transErr", err)
				c.JSON(http.StatusOK, gin.H{
					"err": "上架失败",
				})
			} else {
				token, _ := strconv.Atoi(tokenId)
				prices, _ := strconv.Atoi(price)
				database.ChangeSell(token, true, prices)
				c.JSON(http.StatusOK, gin.H{
					"ok": "ok",
				})
			}
		}
	}
}
func Buy(c *gin.Context) {
	tokenId, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	price, _ := strconv.Atoi(c.PostForm("price"))
	owner := c.PostForm("owner")
	cookie, e := c.Request.Cookie("name")
	if e == nil {
		user := database.QueryUser(cookie.Value)
		b, _ := contract.ERC20_Allowance(nil, user.Address, contract.TransactionAddress)
		a1 := big.NewInt(int64(price))
		//privateKey, _ := crypto.HexToECDSA(user.Keystore[2:])
		//opts := bind.NewKeyedTransactor(privateKey)
		opts, _ := util.BindOptsByKeystore(database.QueryUser(cookie.Value).Keystore, database.QueryUser(cookie.Value).Password)
		fmt.Println("opts: ", opts)
		if b.Cmp(a1) == -1 {
			_, err := contract.ERC20_Approve(opts, contract.TransactionAddress, a1)
			if err != nil {
				fmt.Println(err)
				c.JSON(200, gin.H{"value": "授权失败"})
				return
			}
		}
		_, err := contract.Trans_Buy(opts, big.NewInt(tokenId))
		if err != nil {
			fmt.Println("Trans_Buy: ", err)
			c.JSON(http.StatusOK, gin.H{
				"err": "购买失败,请确保账户正常或余额充足！",
			})
		} else {
			database.AddTransactions(int(tokenId), owner, user.Address, price)
			database.ChangeOwner(int(tokenId), user.Address)
			c.JSON(http.StatusOK, gin.H{
				"ok": "ok",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"err": "请先登录账号",
		})
	}
}
func Undercarriage(c *gin.Context) {
	tokenId, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	owner := c.PostForm("owner")
	cookie, e := c.Request.Cookie("name")
	if e == nil {
		user := database.QueryUser(cookie.Value)
		if user.Address == owner {
			//privateKey, _ := crypto.HexToECDSA(user.Keystore[2:])
			//opts := bind.NewKeyedTransactor(privateKey)
			opts, _ := util.BindOptsByKeystore(database.QueryUser(cookie.Value).Keystore, database.QueryUser(cookie.Value).Password)
			_, err := contract.Trans_Undo(opts, big.NewInt(tokenId))
			if err != nil {
				fmt.Println(err)
				c.JSON(200, gin.H{
					"err": "下架失败",
				})
			} else {
				database.ChangeSell(int(tokenId), false, 0)
				c.JSON(200, gin.H{
					"ok": "下架成功！",
				})
			}
		} else {
			c.JSON(200, gin.H{
				"err": "暂无权限！",
			})
		}
	} else {
		c.JSON(200, gin.H{
			"err": "请先登录账号",
		})
	}
}
