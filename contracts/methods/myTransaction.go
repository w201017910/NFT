package methods

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"nft/contracts/interfaces"
)

var transactionIns *interfaces.Tranaaction

func Init_Transaction(client *ethclient.Client, contractAddr string) {
	transaction, err := interfaces.NewTranaaction(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}
	transactionIns = transaction
}
func Trans_Sell(opt *bind.TransactOpts, tokenId *big.Int, price *big.Int) (sellRes *types.Transaction, err error) {
	sell, err := transactionIns.Sell(opt, tokenId, price)
	if err != nil {
		return nil, err
	}
	return sell, nil
}
func Trans_Buy(opt *bind.TransactOpts, tokenId *big.Int) (buyRes *types.Transaction, err error) {
	buy, err := transactionIns.Offer(opt, tokenId)
	if err != nil {
		return nil, err
	}
	return buy, nil
}
func Trans_Undo(opt *bind.TransactOpts, tokenId *big.Int) (undoRes *types.Transaction, err error) {
	undo, err := transactionIns.Reback(opt, tokenId)
	if err != nil {
		return nil, err
	}
	return undo, nil
}
func Trans_TokenToAddress(opt *bind.CallOpts, tokenId *big.Int) (addrRes common.Address, err error) {
	address, err := transactionIns.Token(opt, tokenId)
	if err != nil {
		return common.Address{}, err
	}
	return address, nil
}
func Trans_Transaction(opt *bind.CallOpts, address string, tokenId *big.Int) (transRes *big.Int, err error) {
	address_ := common.HexToAddress(address)
	transaction, err := transactionIns.Transaction(opt, address_, tokenId)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
func Trans_Balance(opt *bind.CallOpts, tokenId *big.Int) (balanceRes *big.Int, err error) {
	balance, err := transactionIns.QueryBlance(opt, tokenId)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
