package methods

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"nft/contracts/interfaces"
)

var erc721Ins *interfaces.ERC721

func Init_ERC721(client *ethclient.Client, contractAddr string, privateKey string) (ERC721INS *interfaces.ERC721, opts *bind.TransactOpts) {
	erc721, err := interfaces.NewERC721(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}
	privatekey, _ := crypto.HexToECDSA(privateKey)
	opt := bind.NewKeyedTransactor(privatekey)
	erc721Ins = erc721
	return erc721, opt
}
func ERC721_Approve(opt *bind.TransactOpts, to string, tokenId *big.Int) (approveRes *types.Transaction, err error) {
	toAddr := common.HexToAddress(to)
	approve, err := erc721Ins.Approve(opt, toAddr, tokenId)
	if err != nil {
		return nil, err
	}
	return approve, nil
}
func ERC721_Mint(opt *bind.TransactOpts, to common.Address, cid string, _type string, name string) (mintRes *types.Transaction, err error) {
	mint, err := erc721Ins.Mint1(opt, to, cid, _type, name)
	if err != nil {
		return nil, err
	}
	return mint, nil
}
func ERC721_SafeTransfer(opt *bind.TransactOpts, from string, to string, tokenId int64) (safeTransRes *types.Transaction, err error) {
	fromAddr := common.HexToAddress(from)
	toAddr := common.HexToAddress(to)
	tokenid := big.NewInt(tokenId)
	safeTrans, err := erc721Ins.SafeTransferFrom(opt, fromAddr, toAddr, tokenid)
	if err != nil {
		return nil, err
	}
	return safeTrans, nil
}
func ERC721_TotalToken(opt *bind.CallOpts) (totalTokens *big.Int, err error) {
	totalToken, err := erc721Ins.TotalToken(opt)
	if err != nil {
		return nil, err
	}
	return totalToken, nil
}
func ERC721_BalanceOf(opt *bind.CallOpts, owner string) (balanceRes *big.Int, err error) {
	ownerAddr := common.HexToAddress(owner)
	balance, err := erc721Ins.BalanceOf(opt, ownerAddr)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
func ERC721_GrtApproved(opt *bind.CallOpts, tokenId int64) (getApprovedRes common.Address, err error) {
	tokenid := big.NewInt(tokenId)
	getApproved, err := erc721Ins.GetApproved(opt, tokenid)
	if err != nil {
		return common.Address{}, err
	}
	return getApproved, nil
}
func ERC721_OwnerOf(opt *bind.CallOpts, tokenId int64) (ownerRes common.Address, err error) {
	tokenid := big.NewInt(tokenId)
	ownerOf, err := erc721Ins.OwnerOf(opt, tokenid)
	if err != nil {
		return common.Address{}, err
	}
	return ownerOf, nil
}
func ERC721_Name(opt *bind.CallOpts) (nameRes string, err error) {
	name, err := erc721Ins.Name(opt)
	if err != nil {
		return "", err
	}
	return name, nil
}
