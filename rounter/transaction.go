package rounter

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	contract "nft/contracts/methods"
	"nft/database"
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
			privateKey, _ := crypto.HexToECDSA(database.QueryUser(cookie.Value).Keystore[2:])
			opts := bind.NewKeyedTransactor(privateKey)
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
		fmt.Println(b.String(), a1.String())
		fmt.Println("Address: ", user.Address)
		privateKey, _ := crypto.HexToECDSA(user.Keystore[2:])
		opts := bind.NewKeyedTransactor(privateKey)
		if b.Cmp(a1) == -1 {
			fmt.Println(1)
			_, err := contract.ERC20_Approve(opts, contract.TransactionAddress, a1)
			if err != nil {
				fmt.Println(err)
				c.JSON(200, gin.H{"value": "授权失败"})
				return
			}
		}

		_, err := contract.Trans_Buy(opts, big.NewInt(tokenId))
		if err != nil {
			fmt.Println("trans_err: ", err)
			c.JSON(http.StatusOK, gin.H{
				"err": "购买失败",
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
