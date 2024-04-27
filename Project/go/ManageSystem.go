// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package _go

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

// AssetManagementAsset is an auto generated low-level Go binding around an user-defined struct.
type AssetManagementAsset struct {
	Id              *big.Int
	Name            string
	Description     string
	OwnerAddress    common.Address
	IsApprovedOwner common.Address
	Status          *big.Int
}

// AssetManagementRecord is an auto generated low-level Go binding around an user-defined struct.
type AssetManagementRecord struct {
	Name1         string
	Description1  string
	OwnerAddress1 common.Address
	TypeOf        *big.Int
}

// AssetManagementRequest is an auto generated low-level Go binding around an user-defined struct.
type AssetManagementRequest struct {
	Id              *big.Int
	Name            string
	OwnerAddress    common.Address
	IsApprovedOwner common.Address
	Anwser          *big.Int
}

// AssetManagementreceiveAsset is an auto generated low-level Go binding around an user-defined struct.
type AssetManagementreceiveAsset struct {
	Id             *big.Int
	Name           string
	Description    string
	ReceiveAddress common.Address
	Result         *big.Int
}

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"AssetApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AssetCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AssetTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AssetTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"AssetGroup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"isApprovedOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"approveAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"approveAssetTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"isApprovedOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"createAsset\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"isApprovedOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structAssetManagement.Asset\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAssets\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"isApprovedOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structAssetManagement.Asset[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllRecords\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description1\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddress1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"typeOf\",\"type\":\"uint256\"}],\"internalType\":\"structAssetManagement.Record[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"isApprovedOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"anwser\",\"type\":\"uint256\"}],\"internalType\":\"structAssetManagement.Request[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceiveAsset\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"receiveAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"internalType\":\"structAssetManagement.receiveAsset[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getpersonAssets\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"isApprovedOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structAssetManagement.Asset[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getpersonRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"isApprovedOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"anwser\",\"type\":\"uint256\"}],\"internalType\":\"structAssetManagement.Request[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"personReceiveAsset\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"receiveAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"internalType\":\"structAssetManagement.receiveAsset\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"personReturnAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"receiveRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"receiveAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description1\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddress1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"typeOf\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"refuseAssetRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"}],\"name\":\"rejectAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"requestAssetTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"isApprovedOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"anwser\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalreceiveRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// AssetGroup is a free data retrieval call binding the contract method 0xbfdce43c.
//
// Solidity: function AssetGroup(uint256 , uint256 ) view returns(uint256 id, string name, string description, address ownerAddress, address isApprovedOwner, uint256 status)
func (_Main *MainCaller) AssetGroup(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Id              *big.Int
	Name            string
	Description     string
	OwnerAddress    common.Address
	IsApprovedOwner common.Address
	Status          *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "AssetGroup", arg0, arg1)

	outstruct := new(struct {
		Id              *big.Int
		Name            string
		Description     string
		OwnerAddress    common.Address
		IsApprovedOwner common.Address
		Status          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.OwnerAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.IsApprovedOwner = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AssetGroup is a free data retrieval call binding the contract method 0xbfdce43c.
//
// Solidity: function AssetGroup(uint256 , uint256 ) view returns(uint256 id, string name, string description, address ownerAddress, address isApprovedOwner, uint256 status)
func (_Main *MainSession) AssetGroup(arg0 *big.Int, arg1 *big.Int) (struct {
	Id              *big.Int
	Name            string
	Description     string
	OwnerAddress    common.Address
	IsApprovedOwner common.Address
	Status          *big.Int
}, error) {
	return _Main.Contract.AssetGroup(&_Main.CallOpts, arg0, arg1)
}

// AssetGroup is a free data retrieval call binding the contract method 0xbfdce43c.
//
// Solidity: function AssetGroup(uint256 , uint256 ) view returns(uint256 id, string name, string description, address ownerAddress, address isApprovedOwner, uint256 status)
func (_Main *MainCallerSession) AssetGroup(arg0 *big.Int, arg1 *big.Int) (struct {
	Id              *big.Int
	Name            string
	Description     string
	OwnerAddress    common.Address
	IsApprovedOwner common.Address
	Status          *big.Int
}, error) {
	return _Main.Contract.AssetGroup(&_Main.CallOpts, arg0, arg1)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Main *MainCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Main *MainSession) Admins(arg0 common.Address) (bool, error) {
	return _Main.Contract.Admins(&_Main.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Main *MainCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _Main.Contract.Admins(&_Main.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(uint256 id, string name, string description, address ownerAddress, address isApprovedOwner, uint256 status)
func (_Main *MainCaller) Assets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id              *big.Int
	Name            string
	Description     string
	OwnerAddress    common.Address
	IsApprovedOwner common.Address
	Status          *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "assets", arg0)

	outstruct := new(struct {
		Id              *big.Int
		Name            string
		Description     string
		OwnerAddress    common.Address
		IsApprovedOwner common.Address
		Status          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.OwnerAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.IsApprovedOwner = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(uint256 id, string name, string description, address ownerAddress, address isApprovedOwner, uint256 status)
func (_Main *MainSession) Assets(arg0 *big.Int) (struct {
	Id              *big.Int
	Name            string
	Description     string
	OwnerAddress    common.Address
	IsApprovedOwner common.Address
	Status          *big.Int
}, error) {
	return _Main.Contract.Assets(&_Main.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(uint256 id, string name, string description, address ownerAddress, address isApprovedOwner, uint256 status)
func (_Main *MainCallerSession) Assets(arg0 *big.Int) (struct {
	Id              *big.Int
	Name            string
	Description     string
	OwnerAddress    common.Address
	IsApprovedOwner common.Address
	Status          *big.Int
}, error) {
	return _Main.Contract.Assets(&_Main.CallOpts, arg0)
}

// GetAllAssets is a free data retrieval call binding the contract method 0x2acada4d.
//
// Solidity: function getAllAssets() view returns((uint256,string,string,address,address,uint256)[])
func (_Main *MainCaller) GetAllAssets(opts *bind.CallOpts) ([]AssetManagementAsset, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAllAssets")

	if err != nil {
		return *new([]AssetManagementAsset), err
	}

	out0 := *abi.ConvertType(out[0], new([]AssetManagementAsset)).(*[]AssetManagementAsset)

	return out0, err

}

// GetAllAssets is a free data retrieval call binding the contract method 0x2acada4d.
//
// Solidity: function getAllAssets() view returns((uint256,string,string,address,address,uint256)[])
func (_Main *MainSession) GetAllAssets() ([]AssetManagementAsset, error) {
	return _Main.Contract.GetAllAssets(&_Main.CallOpts)
}

// GetAllAssets is a free data retrieval call binding the contract method 0x2acada4d.
//
// Solidity: function getAllAssets() view returns((uint256,string,string,address,address,uint256)[])
func (_Main *MainCallerSession) GetAllAssets() ([]AssetManagementAsset, error) {
	return _Main.Contract.GetAllAssets(&_Main.CallOpts)
}

// GetAllRecords is a free data retrieval call binding the contract method 0xa7f9fe72.
//
// Solidity: function getAllRecords() view returns((string,string,address,uint256)[])
func (_Main *MainCaller) GetAllRecords(opts *bind.CallOpts) ([]AssetManagementRecord, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAllRecords")

	if err != nil {
		return *new([]AssetManagementRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]AssetManagementRecord)).(*[]AssetManagementRecord)

	return out0, err

}

// GetAllRecords is a free data retrieval call binding the contract method 0xa7f9fe72.
//
// Solidity: function getAllRecords() view returns((string,string,address,uint256)[])
func (_Main *MainSession) GetAllRecords() ([]AssetManagementRecord, error) {
	return _Main.Contract.GetAllRecords(&_Main.CallOpts)
}

// GetAllRecords is a free data retrieval call binding the contract method 0xa7f9fe72.
//
// Solidity: function getAllRecords() view returns((string,string,address,uint256)[])
func (_Main *MainCallerSession) GetAllRecords() ([]AssetManagementRecord, error) {
	return _Main.Contract.GetAllRecords(&_Main.CallOpts)
}

// GetAllRequests is a free data retrieval call binding the contract method 0xfb1a002a.
//
// Solidity: function getAllRequests() view returns((uint256,string,address,address,uint256)[])
func (_Main *MainCaller) GetAllRequests(opts *bind.CallOpts) ([]AssetManagementRequest, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getAllRequests")

	if err != nil {
		return *new([]AssetManagementRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]AssetManagementRequest)).(*[]AssetManagementRequest)

	return out0, err

}

// GetAllRequests is a free data retrieval call binding the contract method 0xfb1a002a.
//
// Solidity: function getAllRequests() view returns((uint256,string,address,address,uint256)[])
func (_Main *MainSession) GetAllRequests() ([]AssetManagementRequest, error) {
	return _Main.Contract.GetAllRequests(&_Main.CallOpts)
}

// GetAllRequests is a free data retrieval call binding the contract method 0xfb1a002a.
//
// Solidity: function getAllRequests() view returns((uint256,string,address,address,uint256)[])
func (_Main *MainCallerSession) GetAllRequests() ([]AssetManagementRequest, error) {
	return _Main.Contract.GetAllRequests(&_Main.CallOpts)
}

// GetReceiveAsset is a free data retrieval call binding the contract method 0x26adc6d6.
//
// Solidity: function getReceiveAsset() view returns((uint256,string,string,address,uint256)[])
func (_Main *MainCaller) GetReceiveAsset(opts *bind.CallOpts) ([]AssetManagementreceiveAsset, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getReceiveAsset")

	if err != nil {
		return *new([]AssetManagementreceiveAsset), err
	}

	out0 := *abi.ConvertType(out[0], new([]AssetManagementreceiveAsset)).(*[]AssetManagementreceiveAsset)

	return out0, err

}

// GetReceiveAsset is a free data retrieval call binding the contract method 0x26adc6d6.
//
// Solidity: function getReceiveAsset() view returns((uint256,string,string,address,uint256)[])
func (_Main *MainSession) GetReceiveAsset() ([]AssetManagementreceiveAsset, error) {
	return _Main.Contract.GetReceiveAsset(&_Main.CallOpts)
}

// GetReceiveAsset is a free data retrieval call binding the contract method 0x26adc6d6.
//
// Solidity: function getReceiveAsset() view returns((uint256,string,string,address,uint256)[])
func (_Main *MainCallerSession) GetReceiveAsset() ([]AssetManagementreceiveAsset, error) {
	return _Main.Contract.GetReceiveAsset(&_Main.CallOpts)
}

// GetpersonAssets is a free data retrieval call binding the contract method 0x8cd09b26.
//
// Solidity: function getpersonAssets(address _address) view returns((uint256,string,string,address,address,uint256)[])
func (_Main *MainCaller) GetpersonAssets(opts *bind.CallOpts, _address common.Address) ([]AssetManagementAsset, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getpersonAssets", _address)

	if err != nil {
		return *new([]AssetManagementAsset), err
	}

	out0 := *abi.ConvertType(out[0], new([]AssetManagementAsset)).(*[]AssetManagementAsset)

	return out0, err

}

// GetpersonAssets is a free data retrieval call binding the contract method 0x8cd09b26.
//
// Solidity: function getpersonAssets(address _address) view returns((uint256,string,string,address,address,uint256)[])
func (_Main *MainSession) GetpersonAssets(_address common.Address) ([]AssetManagementAsset, error) {
	return _Main.Contract.GetpersonAssets(&_Main.CallOpts, _address)
}

// GetpersonAssets is a free data retrieval call binding the contract method 0x8cd09b26.
//
// Solidity: function getpersonAssets(address _address) view returns((uint256,string,string,address,address,uint256)[])
func (_Main *MainCallerSession) GetpersonAssets(_address common.Address) ([]AssetManagementAsset, error) {
	return _Main.Contract.GetpersonAssets(&_Main.CallOpts, _address)
}

// GetpersonRequests is a free data retrieval call binding the contract method 0xacf82517.
//
// Solidity: function getpersonRequests(address _address) view returns((uint256,string,address,address,uint256)[])
func (_Main *MainCaller) GetpersonRequests(opts *bind.CallOpts, _address common.Address) ([]AssetManagementRequest, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getpersonRequests", _address)

	if err != nil {
		return *new([]AssetManagementRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]AssetManagementRequest)).(*[]AssetManagementRequest)

	return out0, err

}

// GetpersonRequests is a free data retrieval call binding the contract method 0xacf82517.
//
// Solidity: function getpersonRequests(address _address) view returns((uint256,string,address,address,uint256)[])
func (_Main *MainSession) GetpersonRequests(_address common.Address) ([]AssetManagementRequest, error) {
	return _Main.Contract.GetpersonRequests(&_Main.CallOpts, _address)
}

// GetpersonRequests is a free data retrieval call binding the contract method 0xacf82517.
//
// Solidity: function getpersonRequests(address _address) view returns((uint256,string,address,address,uint256)[])
func (_Main *MainCallerSession) GetpersonRequests(_address common.Address) ([]AssetManagementRequest, error) {
	return _Main.Contract.GetpersonRequests(&_Main.CallOpts, _address)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCallerSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// ReceiveRequest is a free data retrieval call binding the contract method 0xe8923a0a.
//
// Solidity: function receiveRequest(uint256 ) view returns(uint256 id, string name, string description, address receiveAddress, uint256 result)
func (_Main *MainCaller) ReceiveRequest(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id             *big.Int
	Name           string
	Description    string
	ReceiveAddress common.Address
	Result         *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "receiveRequest", arg0)

	outstruct := new(struct {
		Id             *big.Int
		Name           string
		Description    string
		ReceiveAddress common.Address
		Result         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ReceiveAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Result = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReceiveRequest is a free data retrieval call binding the contract method 0xe8923a0a.
//
// Solidity: function receiveRequest(uint256 ) view returns(uint256 id, string name, string description, address receiveAddress, uint256 result)
func (_Main *MainSession) ReceiveRequest(arg0 *big.Int) (struct {
	Id             *big.Int
	Name           string
	Description    string
	ReceiveAddress common.Address
	Result         *big.Int
}, error) {
	return _Main.Contract.ReceiveRequest(&_Main.CallOpts, arg0)
}

// ReceiveRequest is a free data retrieval call binding the contract method 0xe8923a0a.
//
// Solidity: function receiveRequest(uint256 ) view returns(uint256 id, string name, string description, address receiveAddress, uint256 result)
func (_Main *MainCallerSession) ReceiveRequest(arg0 *big.Int) (struct {
	Id             *big.Int
	Name           string
	Description    string
	ReceiveAddress common.Address
	Result         *big.Int
}, error) {
	return _Main.Contract.ReceiveRequest(&_Main.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x34461067.
//
// Solidity: function records(uint256 ) view returns(string name1, string description1, address ownerAddress1, uint256 typeOf)
func (_Main *MainCaller) Records(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name1         string
	Description1  string
	OwnerAddress1 common.Address
	TypeOf        *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "records", arg0)

	outstruct := new(struct {
		Name1         string
		Description1  string
		OwnerAddress1 common.Address
		TypeOf        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name1 = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Description1 = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.OwnerAddress1 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TypeOf = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Records is a free data retrieval call binding the contract method 0x34461067.
//
// Solidity: function records(uint256 ) view returns(string name1, string description1, address ownerAddress1, uint256 typeOf)
func (_Main *MainSession) Records(arg0 *big.Int) (struct {
	Name1         string
	Description1  string
	OwnerAddress1 common.Address
	TypeOf        *big.Int
}, error) {
	return _Main.Contract.Records(&_Main.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x34461067.
//
// Solidity: function records(uint256 ) view returns(string name1, string description1, address ownerAddress1, uint256 typeOf)
func (_Main *MainCallerSession) Records(arg0 *big.Int) (struct {
	Name1         string
	Description1  string
	OwnerAddress1 common.Address
	TypeOf        *big.Int
}, error) {
	return _Main.Contract.Records(&_Main.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 id, string name, address ownerAddress, address isApprovedOwner, uint256 anwser)
func (_Main *MainCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id              *big.Int
	Name            string
	OwnerAddress    common.Address
	IsApprovedOwner common.Address
	Anwser          *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		Id              *big.Int
		Name            string
		OwnerAddress    common.Address
		IsApprovedOwner common.Address
		Anwser          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.OwnerAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.IsApprovedOwner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Anwser = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 id, string name, address ownerAddress, address isApprovedOwner, uint256 anwser)
func (_Main *MainSession) Requests(arg0 *big.Int) (struct {
	Id              *big.Int
	Name            string
	OwnerAddress    common.Address
	IsApprovedOwner common.Address
	Anwser          *big.Int
}, error) {
	return _Main.Contract.Requests(&_Main.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 id, string name, address ownerAddress, address isApprovedOwner, uint256 anwser)
func (_Main *MainCallerSession) Requests(arg0 *big.Int) (struct {
	Id              *big.Int
	Name            string
	OwnerAddress    common.Address
	IsApprovedOwner common.Address
	Anwser          *big.Int
}, error) {
	return _Main.Contract.Requests(&_Main.CallOpts, arg0)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Main *MainCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Main *MainSession) TotalAssets() (*big.Int, error) {
	return _Main.Contract.TotalAssets(&_Main.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Main *MainCallerSession) TotalAssets() (*big.Int, error) {
	return _Main.Contract.TotalAssets(&_Main.CallOpts)
}

// TotalreceiveRequest is a free data retrieval call binding the contract method 0x29c489e2.
//
// Solidity: function totalreceiveRequest() view returns(uint256)
func (_Main *MainCaller) TotalreceiveRequest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "totalreceiveRequest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalreceiveRequest is a free data retrieval call binding the contract method 0x29c489e2.
//
// Solidity: function totalreceiveRequest() view returns(uint256)
func (_Main *MainSession) TotalreceiveRequest() (*big.Int, error) {
	return _Main.Contract.TotalreceiveRequest(&_Main.CallOpts)
}

// TotalreceiveRequest is a free data retrieval call binding the contract method 0x29c489e2.
//
// Solidity: function totalreceiveRequest() view returns(uint256)
func (_Main *MainCallerSession) TotalreceiveRequest() (*big.Int, error) {
	return _Main.Contract.TotalreceiveRequest(&_Main.CallOpts)
}

// ApproveAsset is a paid mutator transaction binding the contract method 0x497a8125.
//
// Solidity: function approveAsset(address initiator, uint256 _assetId) returns(bool)
func (_Main *MainTransactor) ApproveAsset(opts *bind.TransactOpts, initiator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "approveAsset", initiator, _assetId)
}

// ApproveAsset is a paid mutator transaction binding the contract method 0x497a8125.
//
// Solidity: function approveAsset(address initiator, uint256 _assetId) returns(bool)
func (_Main *MainSession) ApproveAsset(initiator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ApproveAsset(&_Main.TransactOpts, initiator, _assetId)
}

// ApproveAsset is a paid mutator transaction binding the contract method 0x497a8125.
//
// Solidity: function approveAsset(address initiator, uint256 _assetId) returns(bool)
func (_Main *MainTransactorSession) ApproveAsset(initiator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ApproveAsset(&_Main.TransactOpts, initiator, _assetId)
}

// ApproveAssetTransfer is a paid mutator transaction binding the contract method 0x8f66d92a.
//
// Solidity: function approveAssetTransfer(address initiator, uint256 _assetId) returns(bool)
func (_Main *MainTransactor) ApproveAssetTransfer(opts *bind.TransactOpts, initiator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "approveAssetTransfer", initiator, _assetId)
}

// ApproveAssetTransfer is a paid mutator transaction binding the contract method 0x8f66d92a.
//
// Solidity: function approveAssetTransfer(address initiator, uint256 _assetId) returns(bool)
func (_Main *MainSession) ApproveAssetTransfer(initiator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ApproveAssetTransfer(&_Main.TransactOpts, initiator, _assetId)
}

// ApproveAssetTransfer is a paid mutator transaction binding the contract method 0x8f66d92a.
//
// Solidity: function approveAssetTransfer(address initiator, uint256 _assetId) returns(bool)
func (_Main *MainTransactorSession) ApproveAssetTransfer(initiator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ApproveAssetTransfer(&_Main.TransactOpts, initiator, _assetId)
}

// CreateAsset is a paid mutator transaction binding the contract method 0x251f46e8.
//
// Solidity: function createAsset(address initiator, string _name, string _description) returns((uint256,string,string,address,address,uint256))
func (_Main *MainTransactor) CreateAsset(opts *bind.TransactOpts, initiator common.Address, _name string, _description string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "createAsset", initiator, _name, _description)
}

// CreateAsset is a paid mutator transaction binding the contract method 0x251f46e8.
//
// Solidity: function createAsset(address initiator, string _name, string _description) returns((uint256,string,string,address,address,uint256))
func (_Main *MainSession) CreateAsset(initiator common.Address, _name string, _description string) (*types.Transaction, error) {
	return _Main.Contract.CreateAsset(&_Main.TransactOpts, initiator, _name, _description)
}

// CreateAsset is a paid mutator transaction binding the contract method 0x251f46e8.
//
// Solidity: function createAsset(address initiator, string _name, string _description) returns((uint256,string,string,address,address,uint256))
func (_Main *MainTransactorSession) CreateAsset(initiator common.Address, _name string, _description string) (*types.Transaction, error) {
	return _Main.Contract.CreateAsset(&_Main.TransactOpts, initiator, _name, _description)
}

// PersonReceiveAsset is a paid mutator transaction binding the contract method 0x9c7b645e.
//
// Solidity: function personReceiveAsset(address initiator, uint256 _id) returns((uint256,string,string,address,uint256))
func (_Main *MainTransactor) PersonReceiveAsset(opts *bind.TransactOpts, initiator common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "personReceiveAsset", initiator, _id)
}

// PersonReceiveAsset is a paid mutator transaction binding the contract method 0x9c7b645e.
//
// Solidity: function personReceiveAsset(address initiator, uint256 _id) returns((uint256,string,string,address,uint256))
func (_Main *MainSession) PersonReceiveAsset(initiator common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PersonReceiveAsset(&_Main.TransactOpts, initiator, _id)
}

// PersonReceiveAsset is a paid mutator transaction binding the contract method 0x9c7b645e.
//
// Solidity: function personReceiveAsset(address initiator, uint256 _id) returns((uint256,string,string,address,uint256))
func (_Main *MainTransactorSession) PersonReceiveAsset(initiator common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PersonReceiveAsset(&_Main.TransactOpts, initiator, _id)
}

// PersonReturnAsset is a paid mutator transaction binding the contract method 0xd00e9223.
//
// Solidity: function personReturnAsset(address initiator, uint256 _id) returns(bool)
func (_Main *MainTransactor) PersonReturnAsset(opts *bind.TransactOpts, initiator common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "personReturnAsset", initiator, _id)
}

// PersonReturnAsset is a paid mutator transaction binding the contract method 0xd00e9223.
//
// Solidity: function personReturnAsset(address initiator, uint256 _id) returns(bool)
func (_Main *MainSession) PersonReturnAsset(initiator common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PersonReturnAsset(&_Main.TransactOpts, initiator, _id)
}

// PersonReturnAsset is a paid mutator transaction binding the contract method 0xd00e9223.
//
// Solidity: function personReturnAsset(address initiator, uint256 _id) returns(bool)
func (_Main *MainTransactorSession) PersonReturnAsset(initiator common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Main.Contract.PersonReturnAsset(&_Main.TransactOpts, initiator, _id)
}

// RefuseAssetRequest is a paid mutator transaction binding the contract method 0x20aa3814.
//
// Solidity: function refuseAssetRequest(address initiator, uint256 _assetId) returns(bool)
func (_Main *MainTransactor) RefuseAssetRequest(opts *bind.TransactOpts, initiator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "refuseAssetRequest", initiator, _assetId)
}

// RefuseAssetRequest is a paid mutator transaction binding the contract method 0x20aa3814.
//
// Solidity: function refuseAssetRequest(address initiator, uint256 _assetId) returns(bool)
func (_Main *MainSession) RefuseAssetRequest(initiator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RefuseAssetRequest(&_Main.TransactOpts, initiator, _assetId)
}

// RefuseAssetRequest is a paid mutator transaction binding the contract method 0x20aa3814.
//
// Solidity: function refuseAssetRequest(address initiator, uint256 _assetId) returns(bool)
func (_Main *MainTransactorSession) RefuseAssetRequest(initiator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RefuseAssetRequest(&_Main.TransactOpts, initiator, _assetId)
}

// RejectAsset is a paid mutator transaction binding the contract method 0xc34cad88.
//
// Solidity: function rejectAsset(address initiator, uint256 _assetId) returns(bool)
func (_Main *MainTransactor) RejectAsset(opts *bind.TransactOpts, initiator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "rejectAsset", initiator, _assetId)
}

// RejectAsset is a paid mutator transaction binding the contract method 0xc34cad88.
//
// Solidity: function rejectAsset(address initiator, uint256 _assetId) returns(bool)
func (_Main *MainSession) RejectAsset(initiator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RejectAsset(&_Main.TransactOpts, initiator, _assetId)
}

// RejectAsset is a paid mutator transaction binding the contract method 0xc34cad88.
//
// Solidity: function rejectAsset(address initiator, uint256 _assetId) returns(bool)
func (_Main *MainTransactorSession) RejectAsset(initiator common.Address, _assetId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RejectAsset(&_Main.TransactOpts, initiator, _assetId)
}

// RequestAssetTransfer is a paid mutator transaction binding the contract method 0x064d5d64.
//
// Solidity: function requestAssetTransfer(address initiator, uint256 _assetId, address _to) returns()
func (_Main *MainTransactor) RequestAssetTransfer(opts *bind.TransactOpts, initiator common.Address, _assetId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "requestAssetTransfer", initiator, _assetId, _to)
}

// RequestAssetTransfer is a paid mutator transaction binding the contract method 0x064d5d64.
//
// Solidity: function requestAssetTransfer(address initiator, uint256 _assetId, address _to) returns()
func (_Main *MainSession) RequestAssetTransfer(initiator common.Address, _assetId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Main.Contract.RequestAssetTransfer(&_Main.TransactOpts, initiator, _assetId, _to)
}

// RequestAssetTransfer is a paid mutator transaction binding the contract method 0x064d5d64.
//
// Solidity: function requestAssetTransfer(address initiator, uint256 _assetId, address _to) returns()
func (_Main *MainTransactorSession) RequestAssetTransfer(initiator common.Address, _assetId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Main.Contract.RequestAssetTransfer(&_Main.TransactOpts, initiator, _assetId, _to)
}

// MainAssetApprovedIterator is returned from FilterAssetApproved and is used to iterate over the raw logs and unpacked data for AssetApproved events raised by the Main contract.
type MainAssetApprovedIterator struct {
	Event *MainAssetApproved // Event containing the contract specifics and raw log

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
func (it *MainAssetApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainAssetApproved)
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
		it.Event = new(MainAssetApproved)
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
func (it *MainAssetApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainAssetApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainAssetApproved represents a AssetApproved event raised by the Main contract.
type MainAssetApproved struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAssetApproved is a free log retrieval operation binding the contract event 0x177a61fa4c6c132c66fca3aaf2f302ad52857ac3709daeb6b37f41977089f3a0.
//
// Solidity: event AssetApproved(uint256 indexed id)
func (_Main *MainFilterer) FilterAssetApproved(opts *bind.FilterOpts, id []*big.Int) (*MainAssetApprovedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "AssetApproved", idRule)
	if err != nil {
		return nil, err
	}
	return &MainAssetApprovedIterator{contract: _Main.contract, event: "AssetApproved", logs: logs, sub: sub}, nil
}

// WatchAssetApproved is a free log subscription operation binding the contract event 0x177a61fa4c6c132c66fca3aaf2f302ad52857ac3709daeb6b37f41977089f3a0.
//
// Solidity: event AssetApproved(uint256 indexed id)
func (_Main *MainFilterer) WatchAssetApproved(opts *bind.WatchOpts, sink chan<- *MainAssetApproved, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "AssetApproved", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainAssetApproved)
				if err := _Main.contract.UnpackLog(event, "AssetApproved", log); err != nil {
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

// ParseAssetApproved is a log parse operation binding the contract event 0x177a61fa4c6c132c66fca3aaf2f302ad52857ac3709daeb6b37f41977089f3a0.
//
// Solidity: event AssetApproved(uint256 indexed id)
func (_Main *MainFilterer) ParseAssetApproved(log types.Log) (*MainAssetApproved, error) {
	event := new(MainAssetApproved)
	if err := _Main.contract.UnpackLog(event, "AssetApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainAssetCreatedIterator is returned from FilterAssetCreated and is used to iterate over the raw logs and unpacked data for AssetCreated events raised by the Main contract.
type MainAssetCreatedIterator struct {
	Event *MainAssetCreated // Event containing the contract specifics and raw log

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
func (it *MainAssetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainAssetCreated)
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
		it.Event = new(MainAssetCreated)
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
func (it *MainAssetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainAssetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainAssetCreated represents a AssetCreated event raised by the Main contract.
type MainAssetCreated struct {
	Id          *big.Int
	Name        string
	Description string
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetCreated is a free log retrieval operation binding the contract event 0xb0776d3d4850f9acef0d415f615cdae8a795e98c0010efb0911aabc69927ea88.
//
// Solidity: event AssetCreated(uint256 indexed id, string name, string description, address indexed owner)
func (_Main *MainFilterer) FilterAssetCreated(opts *bind.FilterOpts, id []*big.Int, owner []common.Address) (*MainAssetCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "AssetCreated", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MainAssetCreatedIterator{contract: _Main.contract, event: "AssetCreated", logs: logs, sub: sub}, nil
}

// WatchAssetCreated is a free log subscription operation binding the contract event 0xb0776d3d4850f9acef0d415f615cdae8a795e98c0010efb0911aabc69927ea88.
//
// Solidity: event AssetCreated(uint256 indexed id, string name, string description, address indexed owner)
func (_Main *MainFilterer) WatchAssetCreated(opts *bind.WatchOpts, sink chan<- *MainAssetCreated, id []*big.Int, owner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "AssetCreated", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainAssetCreated)
				if err := _Main.contract.UnpackLog(event, "AssetCreated", log); err != nil {
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

// ParseAssetCreated is a log parse operation binding the contract event 0xb0776d3d4850f9acef0d415f615cdae8a795e98c0010efb0911aabc69927ea88.
//
// Solidity: event AssetCreated(uint256 indexed id, string name, string description, address indexed owner)
func (_Main *MainFilterer) ParseAssetCreated(log types.Log) (*MainAssetCreated, error) {
	event := new(MainAssetCreated)
	if err := _Main.contract.UnpackLog(event, "AssetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainAssetTransferRequestedIterator is returned from FilterAssetTransferRequested and is used to iterate over the raw logs and unpacked data for AssetTransferRequested events raised by the Main contract.
type MainAssetTransferRequestedIterator struct {
	Event *MainAssetTransferRequested // Event containing the contract specifics and raw log

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
func (it *MainAssetTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainAssetTransferRequested)
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
		it.Event = new(MainAssetTransferRequested)
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
func (it *MainAssetTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainAssetTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainAssetTransferRequested represents a AssetTransferRequested event raised by the Main contract.
type MainAssetTransferRequested struct {
	Id   *big.Int
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAssetTransferRequested is a free log retrieval operation binding the contract event 0x6942139dce1e6484f911aebee58c9378ace86f8bb455d4d081a3d8fdc4f1d635.
//
// Solidity: event AssetTransferRequested(uint256 indexed id, address indexed from, address indexed to)
func (_Main *MainFilterer) FilterAssetTransferRequested(opts *bind.FilterOpts, id []*big.Int, from []common.Address, to []common.Address) (*MainAssetTransferRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "AssetTransferRequested", idRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MainAssetTransferRequestedIterator{contract: _Main.contract, event: "AssetTransferRequested", logs: logs, sub: sub}, nil
}

// WatchAssetTransferRequested is a free log subscription operation binding the contract event 0x6942139dce1e6484f911aebee58c9378ace86f8bb455d4d081a3d8fdc4f1d635.
//
// Solidity: event AssetTransferRequested(uint256 indexed id, address indexed from, address indexed to)
func (_Main *MainFilterer) WatchAssetTransferRequested(opts *bind.WatchOpts, sink chan<- *MainAssetTransferRequested, id []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "AssetTransferRequested", idRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainAssetTransferRequested)
				if err := _Main.contract.UnpackLog(event, "AssetTransferRequested", log); err != nil {
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

// ParseAssetTransferRequested is a log parse operation binding the contract event 0x6942139dce1e6484f911aebee58c9378ace86f8bb455d4d081a3d8fdc4f1d635.
//
// Solidity: event AssetTransferRequested(uint256 indexed id, address indexed from, address indexed to)
func (_Main *MainFilterer) ParseAssetTransferRequested(log types.Log) (*MainAssetTransferRequested, error) {
	event := new(MainAssetTransferRequested)
	if err := _Main.contract.UnpackLog(event, "AssetTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainAssetTransferredIterator is returned from FilterAssetTransferred and is used to iterate over the raw logs and unpacked data for AssetTransferred events raised by the Main contract.
type MainAssetTransferredIterator struct {
	Event *MainAssetTransferred // Event containing the contract specifics and raw log

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
func (it *MainAssetTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainAssetTransferred)
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
		it.Event = new(MainAssetTransferred)
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
func (it *MainAssetTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainAssetTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainAssetTransferred represents a AssetTransferred event raised by the Main contract.
type MainAssetTransferred struct {
	Id   *big.Int
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAssetTransferred is a free log retrieval operation binding the contract event 0xa993eb3a10693085bc7afc1de0202310fbd5992b9e51edd263b198f62f20cdae.
//
// Solidity: event AssetTransferred(uint256 indexed id, address indexed from, address indexed to)
func (_Main *MainFilterer) FilterAssetTransferred(opts *bind.FilterOpts, id []*big.Int, from []common.Address, to []common.Address) (*MainAssetTransferredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "AssetTransferred", idRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MainAssetTransferredIterator{contract: _Main.contract, event: "AssetTransferred", logs: logs, sub: sub}, nil
}

// WatchAssetTransferred is a free log subscription operation binding the contract event 0xa993eb3a10693085bc7afc1de0202310fbd5992b9e51edd263b198f62f20cdae.
//
// Solidity: event AssetTransferred(uint256 indexed id, address indexed from, address indexed to)
func (_Main *MainFilterer) WatchAssetTransferred(opts *bind.WatchOpts, sink chan<- *MainAssetTransferred, id []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "AssetTransferred", idRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainAssetTransferred)
				if err := _Main.contract.UnpackLog(event, "AssetTransferred", log); err != nil {
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

// ParseAssetTransferred is a log parse operation binding the contract event 0xa993eb3a10693085bc7afc1de0202310fbd5992b9e51edd263b198f62f20cdae.
//
// Solidity: event AssetTransferred(uint256 indexed id, address indexed from, address indexed to)
func (_Main *MainFilterer) ParseAssetTransferred(log types.Log) (*MainAssetTransferred, error) {
	event := new(MainAssetTransferred)
	if err := _Main.contract.UnpackLog(event, "AssetTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
