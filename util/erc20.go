package util

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	contract "nft/contracts/methods"
)

var erc20Address = "0xCA473e5E9272324864851852AEc94a5CCD5723Ad"
var opts *bind.TransactOpts

func QueryBalance(address common.Address, privateKey_ string) (balance *big.Int) {
	privateKey, err := crypto.HexToECDSA(privateKey_)
	if err != nil {
		log.Fatal(err)
	}
	opts = bind.NewKeyedTransactor(privateKey)
	balances, err := contract.ERC20_BalanceOf(nil, address)
	if err != nil {
		log.Fatal(err)
	}
	return balances
}
