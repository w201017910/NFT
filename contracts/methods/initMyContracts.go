package methods

import "github.com/ethereum/go-ethereum/ethclient"

var Erc20Address = "0x966F1E31f406042C3c50309Dd744d3FdCc6ffBb4"
var Erc721Address = "0x11CCB0D2794b54fDD4bcf88fF7ac2c48fa8957D6"
var TransactionAddress = "0x403b9a0382fd5342d6b2c8722636ccFbB7797200"
var UniswapAddress = "0x1F891dD1fdD7b6659bC70eE9A12F3a06E4f6aFC6"

func InitAll(client *ethclient.Client) {
	Init_ERC20(client, Erc20Address)
	Init_ERC721(client, Erc721Address)
	Init_Transaction(client, TransactionAddress)
	Init_Uniswap(client, UniswapAddress)
}
