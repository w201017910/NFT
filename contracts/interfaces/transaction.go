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

// TranaactionMetaData contains all meta data concerning the Tranaaction contract.
var TranaactionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_transaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"offer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"queryBlance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"reback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TranaactionABI is the input ABI used to generate the binding from.
// Deprecated: Use TranaactionMetaData.ABI instead.
var TranaactionABI = TranaactionMetaData.ABI

// Tranaaction is an auto generated Go binding around an Ethereum contract.
type Tranaaction struct {
	TranaactionCaller     // Read-only binding to the contract
	TranaactionTransactor // Write-only binding to the contract
	TranaactionFilterer   // Log filterer for contract events
}

// TranaactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type TranaactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TranaactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TranaactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TranaactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TranaactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TranaactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TranaactionSession struct {
	Contract     *Tranaaction      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TranaactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TranaactionCallerSession struct {
	Contract *TranaactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TranaactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TranaactionTransactorSession struct {
	Contract     *TranaactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TranaactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type TranaactionRaw struct {
	Contract *Tranaaction // Generic contract binding to access the raw methods on
}

// TranaactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TranaactionCallerRaw struct {
	Contract *TranaactionCaller // Generic read-only contract binding to access the raw methods on
}

// TranaactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TranaactionTransactorRaw struct {
	Contract *TranaactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTranaaction creates a new instance of Tranaaction, bound to a specific deployed contract.
func NewTranaaction(address common.Address, backend bind.ContractBackend) (*Tranaaction, error) {
	contract, err := bindTranaaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tranaaction{TranaactionCaller: TranaactionCaller{contract: contract}, TranaactionTransactor: TranaactionTransactor{contract: contract}, TranaactionFilterer: TranaactionFilterer{contract: contract}}, nil
}

// NewTranaactionCaller creates a new read-only instance of Tranaaction, bound to a specific deployed contract.
func NewTranaactionCaller(address common.Address, caller bind.ContractCaller) (*TranaactionCaller, error) {
	contract, err := bindTranaaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TranaactionCaller{contract: contract}, nil
}

// NewTranaactionTransactor creates a new write-only instance of Tranaaction, bound to a specific deployed contract.
func NewTranaactionTransactor(address common.Address, transactor bind.ContractTransactor) (*TranaactionTransactor, error) {
	contract, err := bindTranaaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TranaactionTransactor{contract: contract}, nil
}

// NewTranaactionFilterer creates a new log filterer instance of Tranaaction, bound to a specific deployed contract.
func NewTranaactionFilterer(address common.Address, filterer bind.ContractFilterer) (*TranaactionFilterer, error) {
	contract, err := bindTranaaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TranaactionFilterer{contract: contract}, nil
}

// bindTranaaction binds a generic wrapper to an already deployed contract.
func bindTranaaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TranaactionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tranaaction *TranaactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tranaaction.Contract.TranaactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tranaaction *TranaactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tranaaction.Contract.TranaactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tranaaction *TranaactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tranaaction.Contract.TranaactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tranaaction *TranaactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tranaaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tranaaction *TranaactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tranaaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tranaaction *TranaactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tranaaction.Contract.contract.Transact(opts, method, params...)
}

// Token is a free data retrieval call binding the contract method 0xb485a999.
//
// Solidity: function _token(uint256 ) view returns(address)
func (_Tranaaction *TranaactionCaller) Token(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Tranaaction.contract.Call(opts, &out, "_token", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xb485a999.
//
// Solidity: function _token(uint256 ) view returns(address)
func (_Tranaaction *TranaactionSession) Token(arg0 *big.Int) (common.Address, error) {
	return _Tranaaction.Contract.Token(&_Tranaaction.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xb485a999.
//
// Solidity: function _token(uint256 ) view returns(address)
func (_Tranaaction *TranaactionCallerSession) Token(arg0 *big.Int) (common.Address, error) {
	return _Tranaaction.Contract.Token(&_Tranaaction.CallOpts, arg0)
}

// Transaction is a free data retrieval call binding the contract method 0xc81db296.
//
// Solidity: function _transaction(address , uint256 ) view returns(uint256)
func (_Tranaaction *TranaactionCaller) Transaction(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tranaaction.contract.Call(opts, &out, "_transaction", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Transaction is a free data retrieval call binding the contract method 0xc81db296.
//
// Solidity: function _transaction(address , uint256 ) view returns(uint256)
func (_Tranaaction *TranaactionSession) Transaction(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Tranaaction.Contract.Transaction(&_Tranaaction.CallOpts, arg0, arg1)
}

// Transaction is a free data retrieval call binding the contract method 0xc81db296.
//
// Solidity: function _transaction(address , uint256 ) view returns(uint256)
func (_Tranaaction *TranaactionCallerSession) Transaction(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Tranaaction.Contract.Transaction(&_Tranaaction.CallOpts, arg0, arg1)
}

// QueryBlance is a free data retrieval call binding the contract method 0x84db5c6f.
//
// Solidity: function queryBlance(uint256 tokenId) view returns(uint256)
func (_Tranaaction *TranaactionCaller) QueryBlance(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tranaaction.contract.Call(opts, &out, "queryBlance", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryBlance is a free data retrieval call binding the contract method 0x84db5c6f.
//
// Solidity: function queryBlance(uint256 tokenId) view returns(uint256)
func (_Tranaaction *TranaactionSession) QueryBlance(tokenId *big.Int) (*big.Int, error) {
	return _Tranaaction.Contract.QueryBlance(&_Tranaaction.CallOpts, tokenId)
}

// QueryBlance is a free data retrieval call binding the contract method 0x84db5c6f.
//
// Solidity: function queryBlance(uint256 tokenId) view returns(uint256)
func (_Tranaaction *TranaactionCallerSession) QueryBlance(tokenId *big.Int) (*big.Int, error) {
	return _Tranaaction.Contract.QueryBlance(&_Tranaaction.CallOpts, tokenId)
}

// Offer is a paid mutator transaction binding the contract method 0x761610fc.
//
// Solidity: function offer(uint256 tokenId) returns()
func (_Tranaaction *TranaactionTransactor) Offer(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Tranaaction.contract.Transact(opts, "offer", tokenId)
}

// Offer is a paid mutator transaction binding the contract method 0x761610fc.
//
// Solidity: function offer(uint256 tokenId) returns()
func (_Tranaaction *TranaactionSession) Offer(tokenId *big.Int) (*types.Transaction, error) {
	return _Tranaaction.Contract.Offer(&_Tranaaction.TransactOpts, tokenId)
}

// Offer is a paid mutator transaction binding the contract method 0x761610fc.
//
// Solidity: function offer(uint256 tokenId) returns()
func (_Tranaaction *TranaactionTransactorSession) Offer(tokenId *big.Int) (*types.Transaction, error) {
	return _Tranaaction.Contract.Offer(&_Tranaaction.TransactOpts, tokenId)
}

// Reback is a paid mutator transaction binding the contract method 0xbf6c2e79.
//
// Solidity: function reback(uint256 tokenId) returns()
func (_Tranaaction *TranaactionTransactor) Reback(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Tranaaction.contract.Transact(opts, "reback", tokenId)
}

// Reback is a paid mutator transaction binding the contract method 0xbf6c2e79.
//
// Solidity: function reback(uint256 tokenId) returns()
func (_Tranaaction *TranaactionSession) Reback(tokenId *big.Int) (*types.Transaction, error) {
	return _Tranaaction.Contract.Reback(&_Tranaaction.TransactOpts, tokenId)
}

// Reback is a paid mutator transaction binding the contract method 0xbf6c2e79.
//
// Solidity: function reback(uint256 tokenId) returns()
func (_Tranaaction *TranaactionTransactorSession) Reback(tokenId *big.Int) (*types.Transaction, error) {
	return _Tranaaction.Contract.Reback(&_Tranaaction.TransactOpts, tokenId)
}

// Sell is a paid mutator transaction binding the contract method 0xd79875eb.
//
// Solidity: function sell(uint256 tokenId, uint256 balance) returns()
func (_Tranaaction *TranaactionTransactor) Sell(opts *bind.TransactOpts, tokenId *big.Int, balance *big.Int) (*types.Transaction, error) {
	return _Tranaaction.contract.Transact(opts, "sell", tokenId, balance)
}

// Sell is a paid mutator transaction binding the contract method 0xd79875eb.
//
// Solidity: function sell(uint256 tokenId, uint256 balance) returns()
func (_Tranaaction *TranaactionSession) Sell(tokenId *big.Int, balance *big.Int) (*types.Transaction, error) {
	return _Tranaaction.Contract.Sell(&_Tranaaction.TransactOpts, tokenId, balance)
}

// Sell is a paid mutator transaction binding the contract method 0xd79875eb.
//
// Solidity: function sell(uint256 tokenId, uint256 balance) returns()
func (_Tranaaction *TranaactionTransactorSession) Sell(tokenId *big.Int, balance *big.Int) (*types.Transaction, error) {
	return _Tranaaction.Contract.Sell(&_Tranaaction.TransactOpts, tokenId, balance)
}
