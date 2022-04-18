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

var uniswapIns *interfaces.UniswapExchange

func Init_Uniswap(client *ethclient.Client, contractAddr string, privateKey string) (swapIns *interfaces.UniswapExchange, opts *bind.TransactOpts) {
	uniswap, err := interfaces.NewUniswapExchange(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}
	privatekey, _ := crypto.HexToECDSA(privateKey)
	opts = bind.NewKeyedTransactor(privatekey)
	uniswapIns = uniswap
	return uniswap, opts
}
func Swap_InitilizeExchange(opt *bind.TransactOpts, tokenAmount *big.Int) (initRes *types.Transaction, err error) {
	initExchange, err := uniswapIns.InitializeExchange(opt, tokenAmount)
	if err != nil {
		return nil, err
	}
	return initExchange, nil
}
func Swap_EthToTokenSwap(opt *bind.TransactOpts, miniTokens *big.Int, _timeOut *big.Int) (result *types.Transaction, err error) {
	ETT, err := uniswapIns.EthToTokenSwap(opt, miniTokens, _timeOut)
	if err != nil {
		return nil, err
	}
	return ETT, nil
}
func Swap_EthToTokenPayment(opt *bind.TransactOpts, miniTokens *big.Int, timeOut *big.Int, recipent string) (result *types.Transaction, err error) {

	ETTP, err := uniswapIns.EthToTokenPayment(opt, miniTokens, timeOut, common.HexToAddress(recipent))
	if err != nil {
		return nil, err
	}
	return ETTP, nil
}
func Swap_TokenToEthSwap(opt *bind.TransactOpts, tokenAmount *big.Int, miniEth *big.Int, timeOut *big.Int) (result *types.Transaction, err error) {

	TTES, err := uniswapIns.TokenToEthSwap(opt, tokenAmount, miniEth, timeOut)
	if err != nil {
		return nil, err
	}
	return TTES, nil
}
func Swap_TokenToEthPayment(opt *bind.TransactOpts, tokenAmount *big.Int, miniEth *big.Int, timeOut *big.Int, recipent string) (result *types.Transaction, err error) {

	TTEP, err := uniswapIns.TokenToEthPayment(opt, tokenAmount, miniEth, timeOut, common.HexToAddress(recipent))
	if err != nil {
		return nil, err
	}
	return TTEP, nil
}
