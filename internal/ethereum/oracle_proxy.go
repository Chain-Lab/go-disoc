// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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
	_ = abi.ConvertType
)

// EthereumMetaData contains all meta data concerning the Ethereum contract.
var EthereumMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encryptedKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tau\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rSignature\",\"type\":\"bytes\"}],\"name\":\"dataRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encryptedData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encryptedKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"dataResponse\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encryptedKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"tau\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rSignature\",\"type\":\"bytes\"}],\"name\":\"requireOracleData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encryptedData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"encryptedKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"respondOracleData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061030e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631d07fbfd1461003b5780633899458914610057575b600080fd5b61005560048036038101906100509190610174565b610073565b005b610071600480360381019061006c9190610174565b6100bc565b005b7f55f6a9c9e937fcf2a23e13d208d4008e5baf112e5703f1956a9cf8cb9e4ea46e8686868686866040516100ac96959493929190610286565b60405180910390a1505050505050565b7f723a51d62591fc948ca09f58283b246c6c40e4c8c897a7e462330e654d582d778686868686866040516100f596959493929190610286565b60405180910390a1505050505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f8401126101345761013361010f565b5b8235905067ffffffffffffffff81111561015157610150610114565b5b60208301915083600182028301111561016d5761016c610119565b5b9250929050565b6000806000806000806060878903121561019157610190610105565b5b600087013567ffffffffffffffff8111156101af576101ae61010a565b5b6101bb89828a0161011e565b9650965050602087013567ffffffffffffffff8111156101de576101dd61010a565b5b6101ea89828a0161011e565b9450945050604087013567ffffffffffffffff81111561020d5761020c61010a565b5b61021989828a0161011e565b92509250509295509295509295565b600082825260208201905092915050565b82818337600083830152505050565b6000601f19601f8301169050919050565b60006102658385610228565b9350610272838584610239565b61027b83610248565b840190509392505050565b600060608201905081810360008301526102a181888a610259565b905081810360208301526102b6818688610259565b905081810360408301526102cb818486610259565b905097965050505050505056fea2646970667358221220a50ed5eb59923450212ea82ecbc79f7ca892cd73e6f646508960e512e74b517864736f6c63430008110033",
}

// EthereumABI is the input ABI used to generate the binding from.
// Deprecated: Use EthereumMetaData.ABI instead.
var EthereumABI = EthereumMetaData.ABI

// EthereumBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthereumMetaData.Bin instead.
var EthereumBin = EthereumMetaData.Bin

// DeployEthereum deploys a new Ethereum contract, binding an instance of Ethereum to it.
func DeployEthereum(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ethereum, error) {
	parsed, err := EthereumMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthereumBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ethereum{EthereumCaller: EthereumCaller{contract: contract}, EthereumTransactor: EthereumTransactor{contract: contract}, EthereumFilterer: EthereumFilterer{contract: contract}}, nil
}

// Ethereum is an auto generated Go binding around an Ethereum contract.
type Ethereum struct {
	EthereumCaller     // Read-only binding to the contract
	EthereumTransactor // Write-only binding to the contract
	EthereumFilterer   // Log filterer for contract events
}

// EthereumCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthereumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthereumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthereumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthereumSession struct {
	Contract     *Ethereum         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthereumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthereumCallerSession struct {
	Contract *EthereumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EthereumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthereumTransactorSession struct {
	Contract     *EthereumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthereumRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthereumRaw struct {
	Contract *Ethereum // Generic contract binding to access the raw methods on
}

// EthereumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthereumCallerRaw struct {
	Contract *EthereumCaller // Generic read-only contract binding to access the raw methods on
}

// EthereumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthereumTransactorRaw struct {
	Contract *EthereumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthereum creates a new instance of Ethereum, bound to a specific deployed contract.
func NewEthereum(address common.Address, backend bind.ContractBackend) (*Ethereum, error) {
	contract, err := bindEthereum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethereum{EthereumCaller: EthereumCaller{contract: contract}, EthereumTransactor: EthereumTransactor{contract: contract}, EthereumFilterer: EthereumFilterer{contract: contract}}, nil
}

// NewEthereumCaller creates a new read-only instance of Ethereum, bound to a specific deployed contract.
func NewEthereumCaller(address common.Address, caller bind.ContractCaller) (*EthereumCaller, error) {
	contract, err := bindEthereum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumCaller{contract: contract}, nil
}

// NewEthereumTransactor creates a new write-only instance of Ethereum, bound to a specific deployed contract.
func NewEthereumTransactor(address common.Address, transactor bind.ContractTransactor) (*EthereumTransactor, error) {
	contract, err := bindEthereum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumTransactor{contract: contract}, nil
}

// NewEthereumFilterer creates a new log filterer instance of Ethereum, bound to a specific deployed contract.
func NewEthereumFilterer(address common.Address, filterer bind.ContractFilterer) (*EthereumFilterer, error) {
	contract, err := bindEthereum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthereumFilterer{contract: contract}, nil
}

// bindEthereum binds a generic wrapper to an already deployed contract.
func bindEthereum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthereumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethereum *EthereumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethereum.Contract.EthereumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethereum *EthereumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.Contract.EthereumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethereum *EthereumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethereum.Contract.EthereumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethereum *EthereumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethereum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethereum *EthereumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethereum *EthereumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethereum.Contract.contract.Transact(opts, method, params...)
}

// RequireOracleData is a paid mutator transaction binding the contract method 0x1d07fbfd.
//
// Solidity: function requireOracleData(bytes encryptedKey, bytes tau, bytes rSignature) returns()
func (_Ethereum *EthereumTransactor) RequireOracleData(opts *bind.TransactOpts, encryptedKey []byte, tau []byte, rSignature []byte) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "requireOracleData", encryptedKey, tau, rSignature)
}

// RequireOracleData is a paid mutator transaction binding the contract method 0x1d07fbfd.
//
// Solidity: function requireOracleData(bytes encryptedKey, bytes tau, bytes rSignature) returns()
func (_Ethereum *EthereumSession) RequireOracleData(encryptedKey []byte, tau []byte, rSignature []byte) (*types.Transaction, error) {
	return _Ethereum.Contract.RequireOracleData(&_Ethereum.TransactOpts, encryptedKey, tau, rSignature)
}

// RequireOracleData is a paid mutator transaction binding the contract method 0x1d07fbfd.
//
// Solidity: function requireOracleData(bytes encryptedKey, bytes tau, bytes rSignature) returns()
func (_Ethereum *EthereumTransactorSession) RequireOracleData(encryptedKey []byte, tau []byte, rSignature []byte) (*types.Transaction, error) {
	return _Ethereum.Contract.RequireOracleData(&_Ethereum.TransactOpts, encryptedKey, tau, rSignature)
}

// RespondOracleData is a paid mutator transaction binding the contract method 0x38994589.
//
// Solidity: function respondOracleData(bytes encryptedData, bytes encryptedKey, bytes signature) returns()
func (_Ethereum *EthereumTransactor) RespondOracleData(opts *bind.TransactOpts, encryptedData []byte, encryptedKey []byte, signature []byte) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "respondOracleData", encryptedData, encryptedKey, signature)
}

// RespondOracleData is a paid mutator transaction binding the contract method 0x38994589.
//
// Solidity: function respondOracleData(bytes encryptedData, bytes encryptedKey, bytes signature) returns()
func (_Ethereum *EthereumSession) RespondOracleData(encryptedData []byte, encryptedKey []byte, signature []byte) (*types.Transaction, error) {
	return _Ethereum.Contract.RespondOracleData(&_Ethereum.TransactOpts, encryptedData, encryptedKey, signature)
}

// RespondOracleData is a paid mutator transaction binding the contract method 0x38994589.
//
// Solidity: function respondOracleData(bytes encryptedData, bytes encryptedKey, bytes signature) returns()
func (_Ethereum *EthereumTransactorSession) RespondOracleData(encryptedData []byte, encryptedKey []byte, signature []byte) (*types.Transaction, error) {
	return _Ethereum.Contract.RespondOracleData(&_Ethereum.TransactOpts, encryptedData, encryptedKey, signature)
}

// EthereumDataRequestIterator is returned from FilterDataRequest and is used to iterate over the raw logs and unpacked data for DataRequest events raised by the Ethereum contract.
type EthereumDataRequestIterator struct {
	Event *EthereumDataRequest // Event containing the contract specifics and raw log

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
func (it *EthereumDataRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumDataRequest)
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
		it.Event = new(EthereumDataRequest)
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
func (it *EthereumDataRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumDataRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumDataRequest represents a DataRequest event raised by the Ethereum contract.
type EthereumDataRequest struct {
	EncryptedKey []byte
	Tau          []byte
	RSignature   []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDataRequest is a free log retrieval operation binding the contract event 0x55f6a9c9e937fcf2a23e13d208d4008e5baf112e5703f1956a9cf8cb9e4ea46e.
//
// Solidity: event dataRequest(bytes encryptedKey, bytes tau, bytes rSignature)
func (_Ethereum *EthereumFilterer) FilterDataRequest(opts *bind.FilterOpts) (*EthereumDataRequestIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "dataRequest")
	if err != nil {
		return nil, err
	}
	return &EthereumDataRequestIterator{contract: _Ethereum.contract, event: "dataRequest", logs: logs, sub: sub}, nil
}

// WatchDataRequest is a free log subscription operation binding the contract event 0x55f6a9c9e937fcf2a23e13d208d4008e5baf112e5703f1956a9cf8cb9e4ea46e.
//
// Solidity: event dataRequest(bytes encryptedKey, bytes tau, bytes rSignature)
func (_Ethereum *EthereumFilterer) WatchDataRequest(opts *bind.WatchOpts, sink chan<- *EthereumDataRequest) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "dataRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumDataRequest)
				if err := _Ethereum.contract.UnpackLog(event, "dataRequest", log); err != nil {
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

// ParseDataRequest is a log parse operation binding the contract event 0x55f6a9c9e937fcf2a23e13d208d4008e5baf112e5703f1956a9cf8cb9e4ea46e.
//
// Solidity: event dataRequest(bytes encryptedKey, bytes tau, bytes rSignature)
func (_Ethereum *EthereumFilterer) ParseDataRequest(log types.Log) (*EthereumDataRequest, error) {
	event := new(EthereumDataRequest)
	if err := _Ethereum.contract.UnpackLog(event, "dataRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumDataResponseIterator is returned from FilterDataResponse and is used to iterate over the raw logs and unpacked data for DataResponse events raised by the Ethereum contract.
type EthereumDataResponseIterator struct {
	Event *EthereumDataResponse // Event containing the contract specifics and raw log

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
func (it *EthereumDataResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumDataResponse)
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
		it.Event = new(EthereumDataResponse)
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
func (it *EthereumDataResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumDataResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumDataResponse represents a DataResponse event raised by the Ethereum contract.
type EthereumDataResponse struct {
	EncryptedData []byte
	EncryptedKey  []byte
	Signature     []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDataResponse is a free log retrieval operation binding the contract event 0x723a51d62591fc948ca09f58283b246c6c40e4c8c897a7e462330e654d582d77.
//
// Solidity: event dataResponse(bytes encryptedData, bytes encryptedKey, bytes signature)
func (_Ethereum *EthereumFilterer) FilterDataResponse(opts *bind.FilterOpts) (*EthereumDataResponseIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "dataResponse")
	if err != nil {
		return nil, err
	}
	return &EthereumDataResponseIterator{contract: _Ethereum.contract, event: "dataResponse", logs: logs, sub: sub}, nil
}

// WatchDataResponse is a free log subscription operation binding the contract event 0x723a51d62591fc948ca09f58283b246c6c40e4c8c897a7e462330e654d582d77.
//
// Solidity: event dataResponse(bytes encryptedData, bytes encryptedKey, bytes signature)
func (_Ethereum *EthereumFilterer) WatchDataResponse(opts *bind.WatchOpts, sink chan<- *EthereumDataResponse) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "dataResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumDataResponse)
				if err := _Ethereum.contract.UnpackLog(event, "dataResponse", log); err != nil {
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

// ParseDataResponse is a log parse operation binding the contract event 0x723a51d62591fc948ca09f58283b246c6c40e4c8c897a7e462330e654d582d77.
//
// Solidity: event dataResponse(bytes encryptedData, bytes encryptedKey, bytes signature)
func (_Ethereum *EthereumFilterer) ParseDataResponse(log types.Log) (*EthereumDataResponse, error) {
	event := new(EthereumDataResponse)
	if err := _Ethereum.contract.UnpackLog(event, "dataResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
