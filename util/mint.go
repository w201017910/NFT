package util

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	contract "nft/contracts/methods"
)

//
//var ERC721Address = "0xf1daeeecA21937B5Ba4dd691a68A352B08299599"
var transContractAddr = "0x0EAb21e5e1DD285d23660f1C86fCE2ecDA21308D"

//var clients, _ = ethclient.Dial("ws://127.0.0.1:7545")

func MintToken(privateKey_ string, to common.Address, cid string, _type string, name string) (tokenId *big.Int) {
	privateKey, err := crypto.HexToECDSA(privateKey_)
	if err != nil {
		log.Fatal(err)
	}
	opts = bind.NewKeyedTransactor(privateKey)
	mintResult, err := contract.ERC721_Mint(opts, to, cid, _type, name)
	if err != nil {
		log.Fatal(err)
	}
	tokenid, err := contract.ERC721_TotalToken(nil)
	approveResult, err := contract.ERC721_Approve(opts, transContractAddr, tokenid)
	fmt.Println("MintToken", mintResult)
	fmt.Println("approveResult", approveResult)
	return tokenid
}
func QueryTokenBalance(owner string) (tokenCount *big.Int) {
	result, err := contract.ERC721_BalanceOf(nil, owner)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
func QueryTokenInfo(tokenId *big.Int) {
	contract.TokenInfo(nil, tokenId)
}
