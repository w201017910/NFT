package rounter

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	contract "nft/contracts/methods"
	"nft/database"
	"strconv"
)

func Balance(c *gin.Context) {
	ethPool, err := contract.Swap_EthPool()
	if err != nil {
		fmt.Println(err)
	}
	tokenPool, err := contract.Swap_TokenPool()
	if err != nil {
		fmt.Println(err)
	}
	invariant, err := contract.Swap_Invariant()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{"eth": ethPool.String(), "token": tokenPool.String(), "invariant": invariant})
}
func EthToTokenSwap(c *gin.Context) {
	cookie, e := c.Request.Cookie("name")
	if e != nil {
		fmt.Println("用户未登录")
		c.JSON(200, gin.H{"value": "用户未登录"})
		return
	}
	u := database.QueryUser(cookie.Value)
	privateKey, err := crypto.HexToECDSA(u.Keystore[2:])
	if err != nil {
		log.Fatal(err)
	}
	a, e := strconv.Atoi(c.PostForm("value"))
	if e != nil {
		fmt.Println("不要输入非数字")
		c.JSON(200, gin.H{"value": "不要输入非数字"})
		return
	}
	opts := bind.NewKeyedTransactor(privateKey)
	opts.Value = big.NewInt(int64(a))
	_, err = contract.Swap_EthToTokenSwap(opts, big.NewInt(1), big.NewInt(99999999999999))
	if err != nil {
		fmt.Println("交易失败")
		c.JSON(200, gin.H{"value": "交易失败"})
		fmt.Println(err)
		return
	}
	fmt.Println("交易成功")
	c.JSON(200, gin.H{"value": "交易成功"})

}
func TokenToEthSwap(c *gin.Context) {
	cookie, e := c.Request.Cookie("name")
	if e != nil {
		fmt.Println("用户未登录")
		c.JSON(200, gin.H{"value": "用户未登录"})
		return
	}
	u := database.QueryUser(cookie.Value)
	privateKey, err := crypto.HexToECDSA(u.Keystore[2:])
	if err != nil {
		log.Fatal(err)
	}
	a, e := strconv.Atoi(c.PostForm("value"))
	if e != nil {
		fmt.Println("不要输入非数字")
		c.JSON(200, gin.H{"value": "不要输入非数字"})
		return
	}
	opts := bind.NewKeyedTransactor(privateKey)

	fmt.Println(a)

	opts.From = common.HexToAddress(u.Address)
	b, e := contract.ERC20_Allowance(nil, u.Address, contract.UniswapAddress)
	a1 := big.NewInt(int64(a))
	fmt.Println(b.String(), a1.String())
	if b.Cmp(a1) == -1 {
		fmt.Println(1)
		_, err = contract.ERC20_Approve(opts, contract.UniswapAddress, a1)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"value": "授权失败"})
			return
		}
	}

	_, err = contract.Swap_TokenToEthSwap(opts, a1, big.NewInt(1), big.NewInt(999999999999999))

	if err != nil {
		fmt.Println("交易失败")
		c.JSON(200, gin.H{"value": "交易失败"})
		fmt.Println(err)
		return
	}
	fmt.Println("交易成功")
	c.JSON(200, gin.H{"value": "交易成功"})

}
