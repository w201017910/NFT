package util

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	contract "nft/contracts/methods"
)

func MintToken(opts *bind.TransactOpts, to common.Address, cid string, _type string, name string) (tokenId *big.Int, err error) {
	mintResult, mintErr := contract.ERC721_Mint(opts, to, cid, _type, name)
	if mintErr != nil {
		return nil, mintErr
	}
	tokenid, idErr := contract.ERC721_TotalToken(nil)
	if idErr != nil {
		return nil, idErr
	}
	//approveResult, err := contract.ERC721_Approve(opts, transContractAddr, tokenid)
	fmt.Println("MintToken", mintResult)
	return tokenid, nil
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
