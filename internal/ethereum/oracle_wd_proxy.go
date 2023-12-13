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

// ProxyMetaData contains all meta data concerning the Proxy contract.
var ProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reqSource\",\"type\":\"bytes\"}],\"name\":\"dataRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"sourceData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reqSource\",\"type\":\"bytes\"}],\"name\":\"dataResponse\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"reqSource\",\"type\":\"bytes\"}],\"name\":\"requireOracleData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sourceData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reqSource\",\"type\":\"bytes\"}],\"name\":\"respondOracleData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyMetaData.ABI instead.
var ProxyABI = ProxyMetaData.ABI

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// RequireOracleData is a paid mutator transaction binding the contract method 0xf84a3d2d.
//
// Solidity: function requireOracleData(bytes reqSource) returns()
func (_Proxy *ProxyTransactor) RequireOracleData(opts *bind.TransactOpts, reqSource []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "requireOracleData", reqSource)
}

// RequireOracleData is a paid mutator transaction binding the contract method 0xf84a3d2d.
//
// Solidity: function requireOracleData(bytes reqSource) returns()
func (_Proxy *ProxySession) RequireOracleData(reqSource []byte) (*types.Transaction, error) {
	return _Proxy.Contract.RequireOracleData(&_Proxy.TransactOpts, reqSource)
}

// RequireOracleData is a paid mutator transaction binding the contract method 0xf84a3d2d.
//
// Solidity: function requireOracleData(bytes reqSource) returns()
func (_Proxy *ProxyTransactorSession) RequireOracleData(reqSource []byte) (*types.Transaction, error) {
	return _Proxy.Contract.RequireOracleData(&_Proxy.TransactOpts, reqSource)
}

// RespondOracleData is a paid mutator transaction binding the contract method 0x6b22545b.
//
// Solidity: function respondOracleData(bytes sourceData, bytes reqSource) returns()
func (_Proxy *ProxyTransactor) RespondOracleData(opts *bind.TransactOpts, sourceData []byte, reqSource []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "respondOracleData", sourceData, reqSource)
}

// RespondOracleData is a paid mutator transaction binding the contract method 0x6b22545b.
//
// Solidity: function respondOracleData(bytes sourceData, bytes reqSource) returns()
func (_Proxy *ProxySession) RespondOracleData(sourceData []byte, reqSource []byte) (*types.Transaction, error) {
	return _Proxy.Contract.RespondOracleData(&_Proxy.TransactOpts, sourceData, reqSource)
}

// RespondOracleData is a paid mutator transaction binding the contract method 0x6b22545b.
//
// Solidity: function respondOracleData(bytes sourceData, bytes reqSource) returns()
func (_Proxy *ProxyTransactorSession) RespondOracleData(sourceData []byte, reqSource []byte) (*types.Transaction, error) {
	return _Proxy.Contract.RespondOracleData(&_Proxy.TransactOpts, sourceData, reqSource)
}

// ProxyDataRequestIterator is returned from FilterDataRequest and is used to iterate over the raw logs and unpacked data for DataRequest events raised by the Proxy contract.
type ProxyDataRequestIterator struct {
	Event *ProxyDataRequest // Event containing the contract specifics and raw log

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
func (it *ProxyDataRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyDataRequest)
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
		it.Event = new(ProxyDataRequest)
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
func (it *ProxyDataRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyDataRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyDataRequest represents a DataRequest event raised by the Proxy contract.
type ProxyDataRequest struct {
	ReqSource []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDataRequest is a free log retrieval operation binding the contract event 0x2e8181db9d458de9cc4a18cd7927f32e2e0499a5d24eacfa49247fa1bfaa057f.
//
// Solidity: event dataRequest(bytes reqSource)
func (_Proxy *ProxyFilterer) FilterDataRequest(opts *bind.FilterOpts) (*ProxyDataRequestIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "dataRequest")
	if err != nil {
		return nil, err
	}
	return &ProxyDataRequestIterator{contract: _Proxy.contract, event: "dataRequest", logs: logs, sub: sub}, nil
}

// WatchDataRequest is a free log subscription operation binding the contract event 0x2e8181db9d458de9cc4a18cd7927f32e2e0499a5d24eacfa49247fa1bfaa057f.
//
// Solidity: event dataRequest(bytes reqSource)
func (_Proxy *ProxyFilterer) WatchDataRequest(opts *bind.WatchOpts, sink chan<- *ProxyDataRequest) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "dataRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyDataRequest)
				if err := _Proxy.contract.UnpackLog(event, "dataRequest", log); err != nil {
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

// ParseDataRequest is a log parse operation binding the contract event 0x2e8181db9d458de9cc4a18cd7927f32e2e0499a5d24eacfa49247fa1bfaa057f.
//
// Solidity: event dataRequest(bytes reqSource)
func (_Proxy *ProxyFilterer) ParseDataRequest(log types.Log) (*ProxyDataRequest, error) {
	event := new(ProxyDataRequest)
	if err := _Proxy.contract.UnpackLog(event, "dataRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyDataResponseIterator is returned from FilterDataResponse and is used to iterate over the raw logs and unpacked data for DataResponse events raised by the Proxy contract.
type ProxyDataResponseIterator struct {
	Event *ProxyDataResponse // Event containing the contract specifics and raw log

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
func (it *ProxyDataResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyDataResponse)
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
		it.Event = new(ProxyDataResponse)
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
func (it *ProxyDataResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyDataResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyDataResponse represents a DataResponse event raised by the Proxy contract.
type ProxyDataResponse struct {
	SourceData []byte
	ReqSource  []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDataResponse is a free log retrieval operation binding the contract event 0xd42fe4a1d37c99453a3e84823384d75d5b9ea59c92433284de2371340b7df7a8.
//
// Solidity: event dataResponse(bytes sourceData, bytes reqSource)
func (_Proxy *ProxyFilterer) FilterDataResponse(opts *bind.FilterOpts) (*ProxyDataResponseIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "dataResponse")
	if err != nil {
		return nil, err
	}
	return &ProxyDataResponseIterator{contract: _Proxy.contract, event: "dataResponse", logs: logs, sub: sub}, nil
}

// WatchDataResponse is a free log subscription operation binding the contract event 0xd42fe4a1d37c99453a3e84823384d75d5b9ea59c92433284de2371340b7df7a8.
//
// Solidity: event dataResponse(bytes sourceData, bytes reqSource)
func (_Proxy *ProxyFilterer) WatchDataResponse(opts *bind.WatchOpts, sink chan<- *ProxyDataResponse) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "dataResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyDataResponse)
				if err := _Proxy.contract.UnpackLog(event, "dataResponse", log); err != nil {
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

// ParseDataResponse is a log parse operation binding the contract event 0xd42fe4a1d37c99453a3e84823384d75d5b9ea59c92433284de2371340b7df7a8.
//
// Solidity: event dataResponse(bytes sourceData, bytes reqSource)
func (_Proxy *ProxyFilterer) ParseDataResponse(log types.Log) (*ProxyDataResponse, error) {
	event := new(ProxyDataResponse)
	if err := _Proxy.contract.UnpackLog(event, "dataResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
