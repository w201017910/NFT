package util

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	contract "nft/contracts/methods"
)

var ERC721Address = "0x71397280CFaf5B1B55CB2899187DeB307011a709"
var transContractAddr = "0x9f9bAD3eA2928C8cE16289F49143FF9C5e943224"
var ERC20Address = "0x71397280CFaf5B1B55CB2899187DeB307011a709"

func MintToken(client *ethclient.Client, privateKey_ string, to common.Address, cid string, _type string, name string) (tokenId *big.Int) {
	fmt.Println("1privateKey", privateKey_)
	_, opts := contract.Init_ERC721(client, ERC721Address, privateKey_)
	privateKey, err := crypto.HexToECDSA(privateKey_)
	if err != nil {
		fmt.Println("HexToECDSAERR:  ", err)
	}
	fmt.Println("2privateKey", privateKey)
	opts = bind.NewKeyedTransactor(privateKey)
	mintResult, err := contract.ERC721_Mint(opts, to, cid, _type, name)
	if err != nil {
		fmt.Println("MintTokenErr:", err)
	}
	tokenid, err := contract.ERC721_TotalToken(nil)
	approveResult, err := contract.ERC721_Approve(opts, transContractAddr, tokenid)
	fmt.Println("MintToken", mintResult)
	fmt.Println("approveResult", approveResult)
	return tokenid
}
