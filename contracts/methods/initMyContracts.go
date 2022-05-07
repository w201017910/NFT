package methods

import "github.com/ethereum/go-ethereum/ethclient"

var Erc20Address = "0x5a3bd595BE596FCee19b11FfFf9af5B1fDC62913"
var Erc721Address = "0xf1daeeecA21937B5Ba4dd691a68A352B08299599"
var TransactionAddress = "0x0EAb21e5e1DD285d23660f1C86fCE2ecDA21308D"
var UniswapAddress = "0x2Bd2E76BdAa9dD336CB8a1c62a78152c0d8061B0"

func InitAll(client *ethclient.Client) {
	Init_ERC20(client, Erc20Address)
	Init_ERC721(client, Erc721Address)
	Init_Transaction(client, TransactionAddress)
	Init_Uniswap(client, UniswapAddress)
}
