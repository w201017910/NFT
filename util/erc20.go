package util

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	contract "nft/contracts/methods"
)

var erc20Address = "0x5a3bd595BE596FCee19b11FfFf9af5B1fDC62913"
var opts *bind.TransactOpts

func QueryBalance(address common.Address) (balance *big.Int) {
	balances, err := contract.ERC20_BalanceOf(nil, address)
	if err != nil {
		log.Fatal(err)
	}
	return balances
}
