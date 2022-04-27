package methods

import "github.com/ethereum/go-ethereum/ethclient"

var erc20Address = "0x5a3bd595BE596FCee19b11FfFf9af5B1fDC62913"
var erc721Address = "0xf1daeeecA21937B5Ba4dd691a68A352B08299599"
var transactionAddress = "0x0EAb21e5e1DD285d23660f1C86fCE2ecDA21308D"
var uniswapAddress = "0xEBFe19f0f2A419B91Ca514a36006C6682f98C31B"

func InitAll(client *ethclient.Client) {
	Init_ERC20(client, erc20Address)
	Init_ERC721(client, erc721Address)
	Init_Transaction(client, transactionAddress)
	Init_Uniswap(client, uniswapAddress)
}
