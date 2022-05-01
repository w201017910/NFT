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
			_, err := contract.Trans_Sell(opts, big.NewInt(t), big.NewInt(p))
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"err": "上架失败",
				})
			}
		}
	}
	//contract.Trans_Sell()
	fmt.Println("tokenId: ", tokenId)
	fmt.Println("price: ", price)
	fmt.Println("owner: ", owner)
}
