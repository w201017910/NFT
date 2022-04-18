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

var erc20Ins *interfaces.ERC20

func Init_ERC20(client *ethclient.Client, contractAddr string, privateKey string) (ERC20INS *interfaces.ERC20, opts *bind.TransactOpts) {
	erc20, err := interfaces.NewERC20(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}
	privatekey, _ := crypto.HexToECDSA(privateKey)
	opt := bind.NewKeyedTransactor(privatekey)
	erc20Ins = erc20
	return erc20, opt
}
func ERC20_Approve(opts *bind.TransactOpts, address string, amount *big.Int) (transResult *types.Transaction, err error) {
	spender := common.HexToAddress(address)
	approve, err := erc20Ins.Approve(opts, spender, amount)
	if err != nil {
		return nil, err
	}
	return approve, nil
}
func ERC20_DecreaseAllowance(opts *bind.TransactOpts, address string, amount int64) (deAllowRes *types.Transaction, err error) {
	spender := common.HexToAddress(address)
	amounts := big.NewInt(amount)
	deAllow, err := erc20Ins.DecreaseAllowance(opts, spender, amounts)
	if err != nil {
		return nil, err
	}
	return deAllow, nil
}
func ERC20_IncreaseAllowance(opts *bind.TransactOpts, address string, amount int64) (inAllowRes *types.Transaction, err error) {
	spender := common.HexToAddress(address)
	addedValue := big.NewInt(amount)
	inAllow, err := erc20Ins.IncreaseAllowance(opts, spender, addedValue)
	if err != nil {
		return nil, err
	}
	return inAllow, nil
}
func ERC20_Transfer(opts *bind.TransactOpts, address string, amount int64) (transRes *types.Transaction, err error) {
	recipient := common.HexToAddress(address)
	amounts := big.NewInt(amount)
	transfer, err := erc20Ins.Transfer(opts, recipient, amounts)
	if err != nil {
		return nil, err
	}
	return transfer, nil
}
func ERC20_TransferFrom(opts *bind.TransactOpts, sender string, receiver string, amount int64) (transFromRes *types.Transaction, err error) {
	senders := common.HexToAddress(sender)
	recipient := common.HexToAddress(receiver)
	amounts := big.NewInt(amount)
	transferFrom, err := erc20Ins.TransferFrom(opts, senders, recipient, amounts)
	if err != nil {
		return nil, err
	}
	return transferFrom, nil
}
func ERC20_Allowance(opts *bind.CallOpts, owner string, spender string) (allRes *big.Int, err error) {
	spenderAddr := common.HexToAddress(spender)
	ownersAddr := common.HexToAddress(owner)
	allowance, err := erc20Ins.Allowance(opts, ownersAddr, spenderAddr)
	if err != nil {
		return nil, err
	}
	return allowance, nil
}
func ERC20_BalanceOf(opts *bind.CallOpts, account string) (balanceRes *big.Int, err error) {
	accountAddr := common.HexToAddress(account)
	balanceOf, err := erc20Ins.BalanceOf(opts, accountAddr)
	if err != nil {
		return nil, err
	}
	return balanceOf, nil
}
func ERC20_Decimals(opts *bind.CallOpts) (decimalsRes uint8, err error) {
	decimals, err := erc20Ins.Decimals(opts)
	if err != nil {
		return 0, err
	}
	return decimals, nil
}
func ERC20_Name(opts *bind.CallOpts) (nameRes string, err error) {
	name, err := erc20Ins.Name(opts)
	if err != nil {
		return "", err
	}
	return name, nil
}
func ERC20_Symbol(opts *bind.CallOpts) (symbolRes string, err error) {
	symbol, err := erc20Ins.Symbol(opts)
	if err != nil {
		return "", err
	}
	return symbol, nil
}
func ERC20_TotalSupply(opts *bind.CallOpts) (supplyRes *big.Int, err error) {
	totalSupply, err := erc20Ins.TotalSupply(opts)
	if err != nil {
		return nil, err
	}
	return totalSupply, nil
}
