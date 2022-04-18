// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interfaces

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// UniswapExchangeMetaData contains all meta data concerning the UniswapExchange contract.
var UniswapExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sharesBurned\",\"type\":\"uint256\"}],\"name\":\"Divestment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ethIn\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokensOut\",\"type\":\"uint256\"}],\"name\":\"EthToTokenPurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sharesPurchased\",\"type\":\"uint256\"}],\"name\":\"Investment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokensIn\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ethOut\",\"type\":\"uint256\"}],\"name\":\"TokenToEthPurchase\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sharesBurned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minTokens\",\"type\":\"uint256\"}],\"name\":\"divestLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"ethToTokenPayment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"ethToTokenSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"getShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"initializeExchange\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"invariant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minShares\",\"type\":\"uint256\"}],\"name\":\"investLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"tokenToEthPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"tokenToEthSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minTokens\",\"type\":\"uint256\"}],\"name\":\"tokenToTokenIn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenPurchased\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokensSold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minTokensReceived\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"tokenToTokenPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenPurchased\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokensSold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minTokensReceived\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"tokenToTokenSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UniswapExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapExchangeMetaData.ABI instead.
var UniswapExchangeABI = UniswapExchangeMetaData.ABI

// UniswapExchange is an auto generated Go binding around an Ethereum contract.
type UniswapExchange struct {
	UniswapExchangeCaller     // Read-only binding to the contract
	UniswapExchangeTransactor // Write-only binding to the contract
	UniswapExchangeFilterer   // Log filterer for contract events
}

// UniswapExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapExchangeSession struct {
	Contract     *UniswapExchange  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapExchangeCallerSession struct {
	Contract *UniswapExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UniswapExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapExchangeTransactorSession struct {
	Contract     *UniswapExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UniswapExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapExchangeRaw struct {
	Contract *UniswapExchange // Generic contract binding to access the raw methods on
}

// UniswapExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapExchangeCallerRaw struct {
	Contract *UniswapExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapExchangeTransactorRaw struct {
	Contract *UniswapExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapExchange creates a new instance of UniswapExchange, bound to a specific deployed contract.
func NewUniswapExchange(address common.Address, backend bind.ContractBackend) (*UniswapExchange, error) {
	contract, err := bindUniswapExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapExchange{UniswapExchangeCaller: UniswapExchangeCaller{contract: contract}, UniswapExchangeTransactor: UniswapExchangeTransactor{contract: contract}, UniswapExchangeFilterer: UniswapExchangeFilterer{contract: contract}}, nil
}

// NewUniswapExchangeCaller creates a new read-only instance of UniswapExchange, bound to a specific deployed contract.
func NewUniswapExchangeCaller(address common.Address, caller bind.ContractCaller) (*UniswapExchangeCaller, error) {
	contract, err := bindUniswapExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapExchangeCaller{contract: contract}, nil
}

// NewUniswapExchangeTransactor creates a new write-only instance of UniswapExchange, bound to a specific deployed contract.
func NewUniswapExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapExchangeTransactor, error) {
	contract, err := bindUniswapExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapExchangeTransactor{contract: contract}, nil
}

// NewUniswapExchangeFilterer creates a new log filterer instance of UniswapExchange, bound to a specific deployed contract.
func NewUniswapExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapExchangeFilterer, error) {
	contract, err := bindUniswapExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapExchangeFilterer{contract: contract}, nil
}

// bindUniswapExchange binds a generic wrapper to an already deployed contract.
func bindUniswapExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapExchange *UniswapExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapExchange.Contract.UniswapExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapExchange *UniswapExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapExchange.Contract.UniswapExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapExchange *UniswapExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapExchange.Contract.UniswapExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapExchange *UniswapExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapExchange *UniswapExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapExchange *UniswapExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapExchange.Contract.contract.Transact(opts, method, params...)
}

// FEERATE is a free data retrieval call binding the contract method 0x2d11c58a.
//
// Solidity: function FEE_RATE() view returns(uint256)
func (_UniswapExchange *UniswapExchangeCaller) FEERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapExchange.contract.Call(opts, &out, "FEE_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEERATE is a free data retrieval call binding the contract method 0x2d11c58a.
//
// Solidity: function FEE_RATE() view returns(uint256)
func (_UniswapExchange *UniswapExchangeSession) FEERATE() (*big.Int, error) {
	return _UniswapExchange.Contract.FEERATE(&_UniswapExchange.CallOpts)
}

// FEERATE is a free data retrieval call binding the contract method 0x2d11c58a.
//
// Solidity: function FEE_RATE() view returns(uint256)
func (_UniswapExchange *UniswapExchangeCallerSession) FEERATE() (*big.Int, error) {
	return _UniswapExchange.Contract.FEERATE(&_UniswapExchange.CallOpts)
}

// EthPool is a free data retrieval call binding the contract method 0xf16673a4.
//
// Solidity: function ethPool() view returns(uint256)
func (_UniswapExchange *UniswapExchangeCaller) EthPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapExchange.contract.Call(opts, &out, "ethPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthPool is a free data retrieval call binding the contract method 0xf16673a4.
//
// Solidity: function ethPool() view returns(uint256)
func (_UniswapExchange *UniswapExchangeSession) EthPool() (*big.Int, error) {
	return _UniswapExchange.Contract.EthPool(&_UniswapExchange.CallOpts)
}

// EthPool is a free data retrieval call binding the contract method 0xf16673a4.
//
// Solidity: function ethPool() view returns(uint256)
func (_UniswapExchange *UniswapExchangeCallerSession) EthPool() (*big.Int, error) {
	return _UniswapExchange.Contract.EthPool(&_UniswapExchange.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_UniswapExchange *UniswapExchangeCaller) FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapExchange.contract.Call(opts, &out, "factoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_UniswapExchange *UniswapExchangeSession) FactoryAddress() (common.Address, error) {
	return _UniswapExchange.Contract.FactoryAddress(&_UniswapExchange.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_UniswapExchange *UniswapExchangeCallerSession) FactoryAddress() (common.Address, error) {
	return _UniswapExchange.Contract.FactoryAddress(&_UniswapExchange.CallOpts)
}

// GetShares is a free data retrieval call binding the contract method 0xf04da65b.
//
// Solidity: function getShares(address _provider) view returns(uint256 _shares)
func (_UniswapExchange *UniswapExchangeCaller) GetShares(opts *bind.CallOpts, _provider common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UniswapExchange.contract.Call(opts, &out, "getShares", _provider)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetShares is a free data retrieval call binding the contract method 0xf04da65b.
//
// Solidity: function getShares(address _provider) view returns(uint256 _shares)
func (_UniswapExchange *UniswapExchangeSession) GetShares(_provider common.Address) (*big.Int, error) {
	return _UniswapExchange.Contract.GetShares(&_UniswapExchange.CallOpts, _provider)
}

// GetShares is a free data retrieval call binding the contract method 0xf04da65b.
//
// Solidity: function getShares(address _provider) view returns(uint256 _shares)
func (_UniswapExchange *UniswapExchangeCallerSession) GetShares(_provider common.Address) (*big.Int, error) {
	return _UniswapExchange.Contract.GetShares(&_UniswapExchange.CallOpts, _provider)
}

// Invariant is a free data retrieval call binding the contract method 0xb03a9a05.
//
// Solidity: function invariant() view returns(uint256)
func (_UniswapExchange *UniswapExchangeCaller) Invariant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapExchange.contract.Call(opts, &out, "invariant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Invariant is a free data retrieval call binding the contract method 0xb03a9a05.
//
// Solidity: function invariant() view returns(uint256)
func (_UniswapExchange *UniswapExchangeSession) Invariant() (*big.Int, error) {
	return _UniswapExchange.Contract.Invariant(&_UniswapExchange.CallOpts)
}

// Invariant is a free data retrieval call binding the contract method 0xb03a9a05.
//
// Solidity: function invariant() view returns(uint256)
func (_UniswapExchange *UniswapExchangeCallerSession) Invariant() (*big.Int, error) {
	return _UniswapExchange.Contract.Invariant(&_UniswapExchange.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_UniswapExchange *UniswapExchangeCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapExchange.contract.Call(opts, &out, "tokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_UniswapExchange *UniswapExchangeSession) TokenAddress() (common.Address, error) {
	return _UniswapExchange.Contract.TokenAddress(&_UniswapExchange.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() view returns(address)
func (_UniswapExchange *UniswapExchangeCallerSession) TokenAddress() (common.Address, error) {
	return _UniswapExchange.Contract.TokenAddress(&_UniswapExchange.CallOpts)
}

// TokenPool is a free data retrieval call binding the contract method 0x104e9929.
//
// Solidity: function tokenPool() view returns(uint256)
func (_UniswapExchange *UniswapExchangeCaller) TokenPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapExchange.contract.Call(opts, &out, "tokenPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenPool is a free data retrieval call binding the contract method 0x104e9929.
//
// Solidity: function tokenPool() view returns(uint256)
func (_UniswapExchange *UniswapExchangeSession) TokenPool() (*big.Int, error) {
	return _UniswapExchange.Contract.TokenPool(&_UniswapExchange.CallOpts)
}

// TokenPool is a free data retrieval call binding the contract method 0x104e9929.
//
// Solidity: function tokenPool() view returns(uint256)
func (_UniswapExchange *UniswapExchangeCallerSession) TokenPool() (*big.Int, error) {
	return _UniswapExchange.Contract.TokenPool(&_UniswapExchange.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_UniswapExchange *UniswapExchangeCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapExchange.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_UniswapExchange *UniswapExchangeSession) TotalShares() (*big.Int, error) {
	return _UniswapExchange.Contract.TotalShares(&_UniswapExchange.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_UniswapExchange *UniswapExchangeCallerSession) TotalShares() (*big.Int, error) {
	return _UniswapExchange.Contract.TotalShares(&_UniswapExchange.CallOpts)
}

// DivestLiquidity is a paid mutator transaction binding the contract method 0x231b4ca4.
//
// Solidity: function divestLiquidity(uint256 _sharesBurned, uint256 _minEth, uint256 _minTokens) returns()
func (_UniswapExchange *UniswapExchangeTransactor) DivestLiquidity(opts *bind.TransactOpts, _sharesBurned *big.Int, _minEth *big.Int, _minTokens *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.contract.Transact(opts, "divestLiquidity", _sharesBurned, _minEth, _minTokens)
}

// DivestLiquidity is a paid mutator transaction binding the contract method 0x231b4ca4.
//
// Solidity: function divestLiquidity(uint256 _sharesBurned, uint256 _minEth, uint256 _minTokens) returns()
func (_UniswapExchange *UniswapExchangeSession) DivestLiquidity(_sharesBurned *big.Int, _minEth *big.Int, _minTokens *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.DivestLiquidity(&_UniswapExchange.TransactOpts, _sharesBurned, _minEth, _minTokens)
}

// DivestLiquidity is a paid mutator transaction binding the contract method 0x231b4ca4.
//
// Solidity: function divestLiquidity(uint256 _sharesBurned, uint256 _minEth, uint256 _minTokens) returns()
func (_UniswapExchange *UniswapExchangeTransactorSession) DivestLiquidity(_sharesBurned *big.Int, _minEth *big.Int, _minTokens *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.DivestLiquidity(&_UniswapExchange.TransactOpts, _sharesBurned, _minEth, _minTokens)
}

// EthToTokenPayment is a paid mutator transaction binding the contract method 0x657aa489.
//
// Solidity: function ethToTokenPayment(uint256 _minTokens, uint256 _timeout, address _recipient) payable returns()
func (_UniswapExchange *UniswapExchangeTransactor) EthToTokenPayment(opts *bind.TransactOpts, _minTokens *big.Int, _timeout *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _UniswapExchange.contract.Transact(opts, "ethToTokenPayment", _minTokens, _timeout, _recipient)
}

// EthToTokenPayment is a paid mutator transaction binding the contract method 0x657aa489.
//
// Solidity: function ethToTokenPayment(uint256 _minTokens, uint256 _timeout, address _recipient) payable returns()
func (_UniswapExchange *UniswapExchangeSession) EthToTokenPayment(_minTokens *big.Int, _timeout *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _UniswapExchange.Contract.EthToTokenPayment(&_UniswapExchange.TransactOpts, _minTokens, _timeout, _recipient)
}

// EthToTokenPayment is a paid mutator transaction binding the contract method 0x657aa489.
//
// Solidity: function ethToTokenPayment(uint256 _minTokens, uint256 _timeout, address _recipient) payable returns()
func (_UniswapExchange *UniswapExchangeTransactorSession) EthToTokenPayment(_minTokens *big.Int, _timeout *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _UniswapExchange.Contract.EthToTokenPayment(&_UniswapExchange.TransactOpts, _minTokens, _timeout, _recipient)
}

// EthToTokenSwap is a paid mutator transaction binding the contract method 0x93ec7c3d.
//
// Solidity: function ethToTokenSwap(uint256 _minTokens, uint256 _timeout) payable returns()
func (_UniswapExchange *UniswapExchangeTransactor) EthToTokenSwap(opts *bind.TransactOpts, _minTokens *big.Int, _timeout *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.contract.Transact(opts, "ethToTokenSwap", _minTokens, _timeout)
}

// EthToTokenSwap is a paid mutator transaction binding the contract method 0x93ec7c3d.
//
// Solidity: function ethToTokenSwap(uint256 _minTokens, uint256 _timeout) payable returns()
func (_UniswapExchange *UniswapExchangeSession) EthToTokenSwap(_minTokens *big.Int, _timeout *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.EthToTokenSwap(&_UniswapExchange.TransactOpts, _minTokens, _timeout)
}

// EthToTokenSwap is a paid mutator transaction binding the contract method 0x93ec7c3d.
//
// Solidity: function ethToTokenSwap(uint256 _minTokens, uint256 _timeout) payable returns()
func (_UniswapExchange *UniswapExchangeTransactorSession) EthToTokenSwap(_minTokens *big.Int, _timeout *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.EthToTokenSwap(&_UniswapExchange.TransactOpts, _minTokens, _timeout)
}

// InitializeExchange is a paid mutator transaction binding the contract method 0xf9935f8f.
//
// Solidity: function initializeExchange(uint256 _tokenAmount) payable returns()
func (_UniswapExchange *UniswapExchangeTransactor) InitializeExchange(opts *bind.TransactOpts, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.contract.Transact(opts, "initializeExchange", _tokenAmount)
}

// InitializeExchange is a paid mutator transaction binding the contract method 0xf9935f8f.
//
// Solidity: function initializeExchange(uint256 _tokenAmount) payable returns()
func (_UniswapExchange *UniswapExchangeSession) InitializeExchange(_tokenAmount *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.InitializeExchange(&_UniswapExchange.TransactOpts, _tokenAmount)
}

// InitializeExchange is a paid mutator transaction binding the contract method 0xf9935f8f.
//
// Solidity: function initializeExchange(uint256 _tokenAmount) payable returns()
func (_UniswapExchange *UniswapExchangeTransactorSession) InitializeExchange(_tokenAmount *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.InitializeExchange(&_UniswapExchange.TransactOpts, _tokenAmount)
}

// InvestLiquidity is a paid mutator transaction binding the contract method 0x12675e01.
//
// Solidity: function investLiquidity(uint256 _minShares) payable returns()
func (_UniswapExchange *UniswapExchangeTransactor) InvestLiquidity(opts *bind.TransactOpts, _minShares *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.contract.Transact(opts, "investLiquidity", _minShares)
}

// InvestLiquidity is a paid mutator transaction binding the contract method 0x12675e01.
//
// Solidity: function investLiquidity(uint256 _minShares) payable returns()
func (_UniswapExchange *UniswapExchangeSession) InvestLiquidity(_minShares *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.InvestLiquidity(&_UniswapExchange.TransactOpts, _minShares)
}

// InvestLiquidity is a paid mutator transaction binding the contract method 0x12675e01.
//
// Solidity: function investLiquidity(uint256 _minShares) payable returns()
func (_UniswapExchange *UniswapExchangeTransactorSession) InvestLiquidity(_minShares *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.InvestLiquidity(&_UniswapExchange.TransactOpts, _minShares)
}

// TokenToEthPayment is a paid mutator transaction binding the contract method 0x007975e7.
//
// Solidity: function tokenToEthPayment(uint256 _tokenAmount, uint256 _minEth, uint256 _timeout, address _recipient) returns()
func (_UniswapExchange *UniswapExchangeTransactor) TokenToEthPayment(opts *bind.TransactOpts, _tokenAmount *big.Int, _minEth *big.Int, _timeout *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _UniswapExchange.contract.Transact(opts, "tokenToEthPayment", _tokenAmount, _minEth, _timeout, _recipient)
}

// TokenToEthPayment is a paid mutator transaction binding the contract method 0x007975e7.
//
// Solidity: function tokenToEthPayment(uint256 _tokenAmount, uint256 _minEth, uint256 _timeout, address _recipient) returns()
func (_UniswapExchange *UniswapExchangeSession) TokenToEthPayment(_tokenAmount *big.Int, _minEth *big.Int, _timeout *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _UniswapExchange.Contract.TokenToEthPayment(&_UniswapExchange.TransactOpts, _tokenAmount, _minEth, _timeout, _recipient)
}

// TokenToEthPayment is a paid mutator transaction binding the contract method 0x007975e7.
//
// Solidity: function tokenToEthPayment(uint256 _tokenAmount, uint256 _minEth, uint256 _timeout, address _recipient) returns()
func (_UniswapExchange *UniswapExchangeTransactorSession) TokenToEthPayment(_tokenAmount *big.Int, _minEth *big.Int, _timeout *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _UniswapExchange.Contract.TokenToEthPayment(&_UniswapExchange.TransactOpts, _tokenAmount, _minEth, _timeout, _recipient)
}

// TokenToEthSwap is a paid mutator transaction binding the contract method 0xaa852f56.
//
// Solidity: function tokenToEthSwap(uint256 _tokenAmount, uint256 _minEth, uint256 _timeout) returns()
func (_UniswapExchange *UniswapExchangeTransactor) TokenToEthSwap(opts *bind.TransactOpts, _tokenAmount *big.Int, _minEth *big.Int, _timeout *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.contract.Transact(opts, "tokenToEthSwap", _tokenAmount, _minEth, _timeout)
}

// TokenToEthSwap is a paid mutator transaction binding the contract method 0xaa852f56.
//
// Solidity: function tokenToEthSwap(uint256 _tokenAmount, uint256 _minEth, uint256 _timeout) returns()
func (_UniswapExchange *UniswapExchangeSession) TokenToEthSwap(_tokenAmount *big.Int, _minEth *big.Int, _timeout *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.TokenToEthSwap(&_UniswapExchange.TransactOpts, _tokenAmount, _minEth, _timeout)
}

// TokenToEthSwap is a paid mutator transaction binding the contract method 0xaa852f56.
//
// Solidity: function tokenToEthSwap(uint256 _tokenAmount, uint256 _minEth, uint256 _timeout) returns()
func (_UniswapExchange *UniswapExchangeTransactorSession) TokenToEthSwap(_tokenAmount *big.Int, _minEth *big.Int, _timeout *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.TokenToEthSwap(&_UniswapExchange.TransactOpts, _tokenAmount, _minEth, _timeout)
}

// TokenToTokenIn is a paid mutator transaction binding the contract method 0x8e9712e1.
//
// Solidity: function tokenToTokenIn(address _recipient, uint256 _minTokens) payable returns(bool)
func (_UniswapExchange *UniswapExchangeTransactor) TokenToTokenIn(opts *bind.TransactOpts, _recipient common.Address, _minTokens *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.contract.Transact(opts, "tokenToTokenIn", _recipient, _minTokens)
}

// TokenToTokenIn is a paid mutator transaction binding the contract method 0x8e9712e1.
//
// Solidity: function tokenToTokenIn(address _recipient, uint256 _minTokens) payable returns(bool)
func (_UniswapExchange *UniswapExchangeSession) TokenToTokenIn(_recipient common.Address, _minTokens *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.TokenToTokenIn(&_UniswapExchange.TransactOpts, _recipient, _minTokens)
}

// TokenToTokenIn is a paid mutator transaction binding the contract method 0x8e9712e1.
//
// Solidity: function tokenToTokenIn(address _recipient, uint256 _minTokens) payable returns(bool)
func (_UniswapExchange *UniswapExchangeTransactorSession) TokenToTokenIn(_recipient common.Address, _minTokens *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.TokenToTokenIn(&_UniswapExchange.TransactOpts, _recipient, _minTokens)
}

// TokenToTokenPayment is a paid mutator transaction binding the contract method 0x4bf213a2.
//
// Solidity: function tokenToTokenPayment(address _tokenPurchased, address _recipient, uint256 _tokensSold, uint256 _minTokensReceived, uint256 _timeout) returns()
func (_UniswapExchange *UniswapExchangeTransactor) TokenToTokenPayment(opts *bind.TransactOpts, _tokenPurchased common.Address, _recipient common.Address, _tokensSold *big.Int, _minTokensReceived *big.Int, _timeout *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.contract.Transact(opts, "tokenToTokenPayment", _tokenPurchased, _recipient, _tokensSold, _minTokensReceived, _timeout)
}

// TokenToTokenPayment is a paid mutator transaction binding the contract method 0x4bf213a2.
//
// Solidity: function tokenToTokenPayment(address _tokenPurchased, address _recipient, uint256 _tokensSold, uint256 _minTokensReceived, uint256 _timeout) returns()
func (_UniswapExchange *UniswapExchangeSession) TokenToTokenPayment(_tokenPurchased common.Address, _recipient common.Address, _tokensSold *big.Int, _minTokensReceived *big.Int, _timeout *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.TokenToTokenPayment(&_UniswapExchange.TransactOpts, _tokenPurchased, _recipient, _tokensSold, _minTokensReceived, _timeout)
}

// TokenToTokenPayment is a paid mutator transaction binding the contract method 0x4bf213a2.
//
// Solidity: function tokenToTokenPayment(address _tokenPurchased, address _recipient, uint256 _tokensSold, uint256 _minTokensReceived, uint256 _timeout) returns()
func (_UniswapExchange *UniswapExchangeTransactorSession) TokenToTokenPayment(_tokenPurchased common.Address, _recipient common.Address, _tokensSold *big.Int, _minTokensReceived *big.Int, _timeout *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.TokenToTokenPayment(&_UniswapExchange.TransactOpts, _tokenPurchased, _recipient, _tokensSold, _minTokensReceived, _timeout)
}

// TokenToTokenSwap is a paid mutator transaction binding the contract method 0x2fdb4612.
//
// Solidity: function tokenToTokenSwap(address _tokenPurchased, uint256 _tokensSold, uint256 _minTokensReceived, uint256 _timeout) returns()
func (_UniswapExchange *UniswapExchangeTransactor) TokenToTokenSwap(opts *bind.TransactOpts, _tokenPurchased common.Address, _tokensSold *big.Int, _minTokensReceived *big.Int, _timeout *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.contract.Transact(opts, "tokenToTokenSwap", _tokenPurchased, _tokensSold, _minTokensReceived, _timeout)
}

// TokenToTokenSwap is a paid mutator transaction binding the contract method 0x2fdb4612.
//
// Solidity: function tokenToTokenSwap(address _tokenPurchased, uint256 _tokensSold, uint256 _minTokensReceived, uint256 _timeout) returns()
func (_UniswapExchange *UniswapExchangeSession) TokenToTokenSwap(_tokenPurchased common.Address, _tokensSold *big.Int, _minTokensReceived *big.Int, _timeout *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.TokenToTokenSwap(&_UniswapExchange.TransactOpts, _tokenPurchased, _tokensSold, _minTokensReceived, _timeout)
}

// TokenToTokenSwap is a paid mutator transaction binding the contract method 0x2fdb4612.
//
// Solidity: function tokenToTokenSwap(address _tokenPurchased, uint256 _tokensSold, uint256 _minTokensReceived, uint256 _timeout) returns()
func (_UniswapExchange *UniswapExchangeTransactorSession) TokenToTokenSwap(_tokenPurchased common.Address, _tokensSold *big.Int, _minTokensReceived *big.Int, _timeout *big.Int) (*types.Transaction, error) {
	return _UniswapExchange.Contract.TokenToTokenSwap(&_UniswapExchange.TransactOpts, _tokenPurchased, _tokensSold, _minTokensReceived, _timeout)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapExchange *UniswapExchangeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapExchange.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapExchange *UniswapExchangeSession) Receive() (*types.Transaction, error) {
	return _UniswapExchange.Contract.Receive(&_UniswapExchange.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapExchange *UniswapExchangeTransactorSession) Receive() (*types.Transaction, error) {
	return _UniswapExchange.Contract.Receive(&_UniswapExchange.TransactOpts)
}

// UniswapExchangeDivestmentIterator is returned from FilterDivestment and is used to iterate over the raw logs and unpacked data for Divestment events raised by the UniswapExchange contract.
type UniswapExchangeDivestmentIterator struct {
	Event *UniswapExchangeDivestment // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniswapExchangeDivestmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapExchangeDivestment)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniswapExchangeDivestment)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniswapExchangeDivestmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapExchangeDivestmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapExchangeDivestment represents a Divestment event raised by the UniswapExchange contract.
type UniswapExchangeDivestment struct {
	LiquidityProvider common.Address
	SharesBurned      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDivestment is a free log retrieval operation binding the contract event 0x04c88fcf03215155a0d0a2fdbd430ab9f32013f00499b2899c180dc6077180ca.
//
// Solidity: event Divestment(address indexed liquidityProvider, uint256 indexed sharesBurned)
func (_UniswapExchange *UniswapExchangeFilterer) FilterDivestment(opts *bind.FilterOpts, liquidityProvider []common.Address, sharesBurned []*big.Int) (*UniswapExchangeDivestmentIterator, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}
	var sharesBurnedRule []interface{}
	for _, sharesBurnedItem := range sharesBurned {
		sharesBurnedRule = append(sharesBurnedRule, sharesBurnedItem)
	}

	logs, sub, err := _UniswapExchange.contract.FilterLogs(opts, "Divestment", liquidityProviderRule, sharesBurnedRule)
	if err != nil {
		return nil, err
	}
	return &UniswapExchangeDivestmentIterator{contract: _UniswapExchange.contract, event: "Divestment", logs: logs, sub: sub}, nil
}

// WatchDivestment is a free log subscription operation binding the contract event 0x04c88fcf03215155a0d0a2fdbd430ab9f32013f00499b2899c180dc6077180ca.
//
// Solidity: event Divestment(address indexed liquidityProvider, uint256 indexed sharesBurned)
func (_UniswapExchange *UniswapExchangeFilterer) WatchDivestment(opts *bind.WatchOpts, sink chan<- *UniswapExchangeDivestment, liquidityProvider []common.Address, sharesBurned []*big.Int) (event.Subscription, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}
	var sharesBurnedRule []interface{}
	for _, sharesBurnedItem := range sharesBurned {
		sharesBurnedRule = append(sharesBurnedRule, sharesBurnedItem)
	}

	logs, sub, err := _UniswapExchange.contract.WatchLogs(opts, "Divestment", liquidityProviderRule, sharesBurnedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapExchangeDivestment)
				if err := _UniswapExchange.contract.UnpackLog(event, "Divestment", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDivestment is a log parse operation binding the contract event 0x04c88fcf03215155a0d0a2fdbd430ab9f32013f00499b2899c180dc6077180ca.
//
// Solidity: event Divestment(address indexed liquidityProvider, uint256 indexed sharesBurned)
func (_UniswapExchange *UniswapExchangeFilterer) ParseDivestment(log types.Log) (*UniswapExchangeDivestment, error) {
	event := new(UniswapExchangeDivestment)
	if err := _UniswapExchange.contract.UnpackLog(event, "Divestment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapExchangeEthToTokenPurchaseIterator is returned from FilterEthToTokenPurchase and is used to iterate over the raw logs and unpacked data for EthToTokenPurchase events raised by the UniswapExchange contract.
type UniswapExchangeEthToTokenPurchaseIterator struct {
	Event *UniswapExchangeEthToTokenPurchase // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniswapExchangeEthToTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapExchangeEthToTokenPurchase)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniswapExchangeEthToTokenPurchase)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniswapExchangeEthToTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapExchangeEthToTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapExchangeEthToTokenPurchase represents a EthToTokenPurchase event raised by the UniswapExchange contract.
type UniswapExchangeEthToTokenPurchase struct {
	Buyer     common.Address
	EthIn     *big.Int
	TokensOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthToTokenPurchase is a free log retrieval operation binding the contract event 0xc2d4406fadcd716ab4a11dd67875462a28108de6a8911747c9ca93c8c3c32c22.
//
// Solidity: event EthToTokenPurchase(address indexed buyer, uint256 indexed ethIn, uint256 indexed tokensOut)
func (_UniswapExchange *UniswapExchangeFilterer) FilterEthToTokenPurchase(opts *bind.FilterOpts, buyer []common.Address, ethIn []*big.Int, tokensOut []*big.Int) (*UniswapExchangeEthToTokenPurchaseIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var ethInRule []interface{}
	for _, ethInItem := range ethIn {
		ethInRule = append(ethInRule, ethInItem)
	}
	var tokensOutRule []interface{}
	for _, tokensOutItem := range tokensOut {
		tokensOutRule = append(tokensOutRule, tokensOutItem)
	}

	logs, sub, err := _UniswapExchange.contract.FilterLogs(opts, "EthToTokenPurchase", buyerRule, ethInRule, tokensOutRule)
	if err != nil {
		return nil, err
	}
	return &UniswapExchangeEthToTokenPurchaseIterator{contract: _UniswapExchange.contract, event: "EthToTokenPurchase", logs: logs, sub: sub}, nil
}

// WatchEthToTokenPurchase is a free log subscription operation binding the contract event 0xc2d4406fadcd716ab4a11dd67875462a28108de6a8911747c9ca93c8c3c32c22.
//
// Solidity: event EthToTokenPurchase(address indexed buyer, uint256 indexed ethIn, uint256 indexed tokensOut)
func (_UniswapExchange *UniswapExchangeFilterer) WatchEthToTokenPurchase(opts *bind.WatchOpts, sink chan<- *UniswapExchangeEthToTokenPurchase, buyer []common.Address, ethIn []*big.Int, tokensOut []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var ethInRule []interface{}
	for _, ethInItem := range ethIn {
		ethInRule = append(ethInRule, ethInItem)
	}
	var tokensOutRule []interface{}
	for _, tokensOutItem := range tokensOut {
		tokensOutRule = append(tokensOutRule, tokensOutItem)
	}

	logs, sub, err := _UniswapExchange.contract.WatchLogs(opts, "EthToTokenPurchase", buyerRule, ethInRule, tokensOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapExchangeEthToTokenPurchase)
				if err := _UniswapExchange.contract.UnpackLog(event, "EthToTokenPurchase", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEthToTokenPurchase is a log parse operation binding the contract event 0xc2d4406fadcd716ab4a11dd67875462a28108de6a8911747c9ca93c8c3c32c22.
//
// Solidity: event EthToTokenPurchase(address indexed buyer, uint256 indexed ethIn, uint256 indexed tokensOut)
func (_UniswapExchange *UniswapExchangeFilterer) ParseEthToTokenPurchase(log types.Log) (*UniswapExchangeEthToTokenPurchase, error) {
	event := new(UniswapExchangeEthToTokenPurchase)
	if err := _UniswapExchange.contract.UnpackLog(event, "EthToTokenPurchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapExchangeInvestmentIterator is returned from FilterInvestment and is used to iterate over the raw logs and unpacked data for Investment events raised by the UniswapExchange contract.
type UniswapExchangeInvestmentIterator struct {
	Event *UniswapExchangeInvestment // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniswapExchangeInvestmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapExchangeInvestment)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniswapExchangeInvestment)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniswapExchangeInvestmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapExchangeInvestmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapExchangeInvestment represents a Investment event raised by the UniswapExchange contract.
type UniswapExchangeInvestment struct {
	LiquidityProvider common.Address
	SharesPurchased   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInvestment is a free log retrieval operation binding the contract event 0xfe123a1efb6782dccd20e9790951adcfab3cd0e11b85a6f06d8b9222d299a0a3.
//
// Solidity: event Investment(address indexed liquidityProvider, uint256 indexed sharesPurchased)
func (_UniswapExchange *UniswapExchangeFilterer) FilterInvestment(opts *bind.FilterOpts, liquidityProvider []common.Address, sharesPurchased []*big.Int) (*UniswapExchangeInvestmentIterator, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}
	var sharesPurchasedRule []interface{}
	for _, sharesPurchasedItem := range sharesPurchased {
		sharesPurchasedRule = append(sharesPurchasedRule, sharesPurchasedItem)
	}

	logs, sub, err := _UniswapExchange.contract.FilterLogs(opts, "Investment", liquidityProviderRule, sharesPurchasedRule)
	if err != nil {
		return nil, err
	}
	return &UniswapExchangeInvestmentIterator{contract: _UniswapExchange.contract, event: "Investment", logs: logs, sub: sub}, nil
}

// WatchInvestment is a free log subscription operation binding the contract event 0xfe123a1efb6782dccd20e9790951adcfab3cd0e11b85a6f06d8b9222d299a0a3.
//
// Solidity: event Investment(address indexed liquidityProvider, uint256 indexed sharesPurchased)
func (_UniswapExchange *UniswapExchangeFilterer) WatchInvestment(opts *bind.WatchOpts, sink chan<- *UniswapExchangeInvestment, liquidityProvider []common.Address, sharesPurchased []*big.Int) (event.Subscription, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}
	var sharesPurchasedRule []interface{}
	for _, sharesPurchasedItem := range sharesPurchased {
		sharesPurchasedRule = append(sharesPurchasedRule, sharesPurchasedItem)
	}

	logs, sub, err := _UniswapExchange.contract.WatchLogs(opts, "Investment", liquidityProviderRule, sharesPurchasedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapExchangeInvestment)
				if err := _UniswapExchange.contract.UnpackLog(event, "Investment", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvestment is a log parse operation binding the contract event 0xfe123a1efb6782dccd20e9790951adcfab3cd0e11b85a6f06d8b9222d299a0a3.
//
// Solidity: event Investment(address indexed liquidityProvider, uint256 indexed sharesPurchased)
func (_UniswapExchange *UniswapExchangeFilterer) ParseInvestment(log types.Log) (*UniswapExchangeInvestment, error) {
	event := new(UniswapExchangeInvestment)
	if err := _UniswapExchange.contract.UnpackLog(event, "Investment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapExchangeTokenToEthPurchaseIterator is returned from FilterTokenToEthPurchase and is used to iterate over the raw logs and unpacked data for TokenToEthPurchase events raised by the UniswapExchange contract.
type UniswapExchangeTokenToEthPurchaseIterator struct {
	Event *UniswapExchangeTokenToEthPurchase // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniswapExchangeTokenToEthPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapExchangeTokenToEthPurchase)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniswapExchangeTokenToEthPurchase)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniswapExchangeTokenToEthPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapExchangeTokenToEthPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapExchangeTokenToEthPurchase represents a TokenToEthPurchase event raised by the UniswapExchange contract.
type UniswapExchangeTokenToEthPurchase struct {
	Buyer    common.Address
	TokensIn *big.Int
	EthOut   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenToEthPurchase is a free log retrieval operation binding the contract event 0x7c27598163cc48e7b019ece0015ee6e3ea794251b47ae935d916f0f1cf5ac16d.
//
// Solidity: event TokenToEthPurchase(address indexed buyer, uint256 indexed tokensIn, uint256 indexed ethOut)
func (_UniswapExchange *UniswapExchangeFilterer) FilterTokenToEthPurchase(opts *bind.FilterOpts, buyer []common.Address, tokensIn []*big.Int, ethOut []*big.Int) (*UniswapExchangeTokenToEthPurchaseIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokensInRule []interface{}
	for _, tokensInItem := range tokensIn {
		tokensInRule = append(tokensInRule, tokensInItem)
	}
	var ethOutRule []interface{}
	for _, ethOutItem := range ethOut {
		ethOutRule = append(ethOutRule, ethOutItem)
	}

	logs, sub, err := _UniswapExchange.contract.FilterLogs(opts, "TokenToEthPurchase", buyerRule, tokensInRule, ethOutRule)
	if err != nil {
		return nil, err
	}
	return &UniswapExchangeTokenToEthPurchaseIterator{contract: _UniswapExchange.contract, event: "TokenToEthPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenToEthPurchase is a free log subscription operation binding the contract event 0x7c27598163cc48e7b019ece0015ee6e3ea794251b47ae935d916f0f1cf5ac16d.
//
// Solidity: event TokenToEthPurchase(address indexed buyer, uint256 indexed tokensIn, uint256 indexed ethOut)
func (_UniswapExchange *UniswapExchangeFilterer) WatchTokenToEthPurchase(opts *bind.WatchOpts, sink chan<- *UniswapExchangeTokenToEthPurchase, buyer []common.Address, tokensIn []*big.Int, ethOut []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokensInRule []interface{}
	for _, tokensInItem := range tokensIn {
		tokensInRule = append(tokensInRule, tokensInItem)
	}
	var ethOutRule []interface{}
	for _, ethOutItem := range ethOut {
		ethOutRule = append(ethOutRule, ethOutItem)
	}

	logs, sub, err := _UniswapExchange.contract.WatchLogs(opts, "TokenToEthPurchase", buyerRule, tokensInRule, ethOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapExchangeTokenToEthPurchase)
				if err := _UniswapExchange.contract.UnpackLog(event, "TokenToEthPurchase", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenToEthPurchase is a log parse operation binding the contract event 0x7c27598163cc48e7b019ece0015ee6e3ea794251b47ae935d916f0f1cf5ac16d.
//
// Solidity: event TokenToEthPurchase(address indexed buyer, uint256 indexed tokensIn, uint256 indexed ethOut)
func (_UniswapExchange *UniswapExchangeFilterer) ParseTokenToEthPurchase(log types.Log) (*UniswapExchangeTokenToEthPurchase, error) {
	event := new(UniswapExchangeTokenToEthPurchase)
	if err := _UniswapExchange.contract.UnpackLog(event, "TokenToEthPurchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
