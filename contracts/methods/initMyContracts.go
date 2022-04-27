package methods

import "github.com/ethereum/go-ethereum/ethclient"

var ERC721Address = "0x20a65434E71a3955cC694e919805E232381fd1a2"
var TransContractAddr = "0xDf045181f5461C5A48051F177D8F845f236bE965"
var ERC20Address = "0x966F1E31f406042C3c50309Dd744d3FdCc6ffBb4"
var UniSwap = "0x1F891dD1fdD7b6659bC70eE9A12F3a06E4f6aFC6"

func InitAll(client *ethclient.Client) {
	Init_ERC20(client, ERC721Address)
	Init_ERC721(client, TransContractAddr)
	Init_Transaction(client, ERC20Address)
	Init_Uniswap(client, UniSwap)
}
