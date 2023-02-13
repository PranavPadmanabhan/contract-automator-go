package web3

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

type ContractCaller struct {
	contract *bind.BoundContract
}

type ContractTransactor struct {
	contract *bind.BoundContract
}

type ContractFilterer struct {
	contract *bind.BoundContract
}

type ContractRaw struct {
	Contract *Contract
}

type ContractCallerRaw struct {
	Contract *ContractCaller
}

type ContractTransactorRaw struct {
	Contract *ContractTransactor
}

type ContractSession struct {
	Contract     *Contract
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ContractCallerSession struct {
	Contract *ContractCaller
	CallOpts bind.CallOpts
}

type ContractTransactorSession struct {
	Contract     *ContractTransactor
	TransactOpts bind.TransactOpts
}

func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer, contractABI string) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func NewContract(address common.Address, backend bind.ContractBackend, abi string) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend, abi)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// read-only contract
func NewContractCaller(address common.Address, caller bind.ContractCaller, abi string) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil, abi)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// write only contract
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor, abi string) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil, abi)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

func (Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

func (Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

func (Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

func (Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return Contract.Contract.contract.Call(opts, result, method, params...)
}

func (Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return Contract.Contract.contract.Transfer(opts)
}
func (Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return Contract.Contract.contract.Transact(opts, method, params...)
}

func (Contract *ContractTransactor) Automate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return Contract.contract.Transact(opts, "automate")
}

func (Contract *ContractSession) Automate() (*types.Transaction, error) {
	return Contract.Contract.Automate(&Contract.TransactOpts)
}

func (Contract *ContractTransactorSession) Automate() (*types.Transaction, error) {
	return Contract.Contract.Automate(&Contract.TransactOpts)
}
func (Contract *ContractCaller) CheckAutomationStatus(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := Contract.contract.Call(opts, &out, "checkAutomationStatus")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (Contract *ContractSession) CheckAutomationStatus() (bool, error) {
	return Contract.Contract.CheckAutomationStatus(&Contract.CallOpts)
}

func (Contract *ContractCallerSession) CheckAutomationStatus() (bool, error) {
	return Contract.Contract.CheckAutomationStatus(&Contract.CallOpts)
}
