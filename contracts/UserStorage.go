// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package UserStorage

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

// UserStorageMetaData contains all meta data concerning the UserStorage contract.
var UserStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"SetUserSucc\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"VertifySignFail\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_dirHash\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"setUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_dirHash\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611423806100606000396000f3fe6080604052600436106100595760003560e01c80632dd34f0f146100625780635e01eb5a1461009f5780636f77926b146100ca5780637dc6b859146101075780638da5cb5b14610130578063a6f9dae11461015b57610060565b3661006057005b005b34801561006e57600080fd5b5061008960048036038101906100849190610ad4565b610184565b6040516100969190610b84565b60405180910390f35b3480156100ab57600080fd5b506100b46101d2565b6040516100c19190610bae565b60405180910390f35b3480156100d657600080fd5b506100f160048036038101906100ec9190610bc9565b6101da565b6040516100fe9190610c86565b60405180910390f35b34801561011357600080fd5b5061012e60048036038101906101299190610ad4565b6102ae565b005b34801561013c57600080fd5b5061014561049a565b6040516101529190610bae565b60405180910390f35b34801561016757600080fd5b50610182600480360381019061017d9190610bc9565b6104be565b005b60008086868660405160200161019c93929190610d2f565b60405160208183030381529060405280519060200120905060006101c288838787610600565b9050809250505095945050505050565b600033905090565b6060600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001805461022990610d88565b80601f016020809104026020016040519081016040528092919081815260200182805461025590610d88565b80156102a25780601f10610277576101008083540402835291602001916102a2565b820191906000526020600020905b81548152906001019060200180831161028557829003601f168201915b50505050509050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461033c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033390610e05565b60405180910390fd5b600061034b8686868686610184565b90508061038d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038490610e71565b60405180910390fd5b604051806020016040528086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815250600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000190816104349190611076565b509050508573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f16e46964c9cae88907d56d563590caa3c362261a8bbdf57e259f3a393ae56a8a60405160405180910390a3505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461054c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054390610e05565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036105bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b390611194565b60405180910390fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600080610652868686868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506106b8565b9050806106ac578573ffffffffffffffffffffffffffffffffffffffff167f72fe1ec9576e79584e0bd5228ae442ef5723b5ed7b2d7369ea5d4d10b55ab92e8686866040516106a39392919061120b565b60405180910390a25b80915050949350505050565b60008060006106c7858561087e565b91509150600060048111156106df576106de61123d565b5b8160048111156106f2576106f161123d565b5b14801561072a57508573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b1561073a57600192505050610877565b6000808773ffffffffffffffffffffffffffffffffffffffff16631626ba7e60e01b888860405160240161076f9291906112b0565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516107d9919061131c565b600060405180830381855afa9150503d8060008114610814576040519150601f19603f3d011682016040523d82523d6000602084013e610819565b606091505b509150915081801561082c575060208151145b80156108705750631626ba7e60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168180602001905181019061086e919061135f565b145b9450505050505b9392505050565b60008060418351036108bf5760008060006020860151925060408601519150606086015160001a90506108b3878285856108cf565b945094505050506108c8565b60006002915091505b9250929050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08360001c111561090a5760006003915091506109a8565b60006001878787876040516000815260200160405260405161092f94939291906113a8565b6020604051602081039080840390855afa158015610951573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361099f576000600192509250506109a8565b80600092509250505b94509492505050565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006109e6826109bb565b9050919050565b6109f6816109db565b8114610a0157600080fd5b50565b600081359050610a13816109ed565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610a3e57610a3d610a19565b5b8235905067ffffffffffffffff811115610a5b57610a5a610a1e565b5b602083019150836001820283011115610a7757610a76610a23565b5b9250929050565b60008083601f840112610a9457610a93610a19565b5b8235905067ffffffffffffffff811115610ab157610ab0610a1e565b5b602083019150836001820283011115610acd57610acc610a23565b5b9250929050565b600080600080600060608688031215610af057610aef6109b1565b5b6000610afe88828901610a04565b955050602086013567ffffffffffffffff811115610b1f57610b1e6109b6565b5b610b2b88828901610a28565b9450945050604086013567ffffffffffffffff811115610b4e57610b4d6109b6565b5b610b5a88828901610a7e565b92509250509295509295909350565b60008115159050919050565b610b7e81610b69565b82525050565b6000602082019050610b996000830184610b75565b92915050565b610ba8816109db565b82525050565b6000602082019050610bc36000830184610b9f565b92915050565b600060208284031215610bdf57610bde6109b1565b5b6000610bed84828501610a04565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610c30578082015181840152602081019050610c15565b60008484015250505050565b6000601f19601f8301169050919050565b6000610c5882610bf6565b610c628185610c01565b9350610c72818560208601610c12565b610c7b81610c3c565b840191505092915050565b60006020820190508181036000830152610ca08184610c4d565b905092915050565b60008160601b9050919050565b6000610cc082610ca8565b9050919050565b6000610cd282610cb5565b9050919050565b610cea610ce5826109db565b610cc7565b82525050565b600081905092915050565b82818337600083830152505050565b6000610d168385610cf0565b9350610d23838584610cfb565b82840190509392505050565b6000610d3b8286610cd9565b601482019150610d4c828486610d0a565b9150819050949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610da057607f821691505b602082108103610db357610db2610d59565b5b50919050565b7f4e6f74206f776e65720000000000000000000000000000000000000000000000600082015250565b6000610def600983610c01565b9150610dfa82610db9565b602082019050919050565b60006020820190508181036000830152610e1e81610de2565b9050919050565b7f566572696669636174696f6e206661696c000000000000000000000000000000600082015250565b6000610e5b601183610c01565b9150610e6682610e25565b602082019050919050565b60006020820190508181036000830152610e8a81610e4e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610f227fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610ee5565b610f2c8683610ee5565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000610f73610f6e610f6984610f44565b610f4e565b610f44565b9050919050565b6000819050919050565b610f8d83610f58565b610fa1610f9982610f7a565b848454610ef2565b825550505050565b600090565b610fb6610fa9565b610fc1818484610f84565b505050565b5b81811015610fe557610fda600082610fae565b600181019050610fc7565b5050565b601f82111561102a57610ffb81610ec0565b61100484610ed5565b81016020851015611013578190505b61102761101f85610ed5565b830182610fc6565b50505b505050565b600082821c905092915050565b600061104d6000198460080261102f565b1980831691505092915050565b6000611066838361103c565b9150826002028217905092915050565b61107f82610bf6565b67ffffffffffffffff81111561109857611097610e91565b5b6110a28254610d88565b6110ad828285610fe9565b600060209050601f8311600181146110e057600084156110ce578287015190505b6110d8858261105a565b865550611140565b601f1984166110ee86610ec0565b60005b82811015611116578489015182556001820191506020850194506020810190506110f1565b86831015611133578489015161112f601f89168261103c565b8355505b6001600288020188555050505b505050505050565b7f4e6f742076616c69642061646472657373000000000000000000000000000000600082015250565b600061117e601183610c01565b915061118982611148565b602082019050919050565b600060208201905081810360008301526111ad81611171565b9050919050565b6000819050919050565b6111c7816111b4565b82525050565b600082825260208201905092915050565b60006111ea83856111cd565b93506111f7838584610cfb565b61120083610c3c565b840190509392505050565b600060408201905061122060008301866111be565b81810360208301526112338184866111de565b9050949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600081519050919050565b60006112828261126c565b61128c81856111cd565b935061129c818560208601610c12565b6112a581610c3c565b840191505092915050565b60006040820190506112c560008301856111be565b81810360208301526112d78184611277565b90509392505050565b600081905092915050565b60006112f68261126c565b61130081856112e0565b9350611310818560208601610c12565b80840191505092915050565b600061132882846112eb565b915081905092915050565b61133c816111b4565b811461134757600080fd5b50565b60008151905061135981611333565b92915050565b600060208284031215611375576113746109b1565b5b60006113838482850161134a565b91505092915050565b600060ff82169050919050565b6113a28161138c565b82525050565b60006080820190506113bd60008301876111be565b6113ca6020830186611399565b6113d760408301856111be565b6113e460608301846111be565b9594505050505056fea26469706673582212208aaa6ce6a54009518a64d5dd634c21d9164444029f52905b17ced976d07049d464736f6c63430008120033",
}

// UserStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use UserStorageMetaData.ABI instead.
var UserStorageABI = UserStorageMetaData.ABI

// UserStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UserStorageMetaData.Bin instead.
var UserStorageBin = UserStorageMetaData.Bin

// DeployUserStorage deploys a new Ethereum contract, binding an instance of UserStorage to it.
func DeployUserStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserStorage, error) {
	parsed, err := UserStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UserStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserStorage{UserStorageCaller: UserStorageCaller{contract: contract}, UserStorageTransactor: UserStorageTransactor{contract: contract}, UserStorageFilterer: UserStorageFilterer{contract: contract}}, nil
}

// UserStorage is an auto generated Go binding around an Ethereum contract.
type UserStorage struct {
	UserStorageCaller     // Read-only binding to the contract
	UserStorageTransactor // Write-only binding to the contract
	UserStorageFilterer   // Log filterer for contract events
}

// UserStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserStorageSession struct {
	Contract     *UserStorage      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserStorageCallerSession struct {
	Contract *UserStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// UserStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserStorageTransactorSession struct {
	Contract     *UserStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UserStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserStorageRaw struct {
	Contract *UserStorage // Generic contract binding to access the raw methods on
}

// UserStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserStorageCallerRaw struct {
	Contract *UserStorageCaller // Generic read-only contract binding to access the raw methods on
}

// UserStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserStorageTransactorRaw struct {
	Contract *UserStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserStorage creates a new instance of UserStorage, bound to a specific deployed contract.
func NewUserStorage(address common.Address, backend bind.ContractBackend) (*UserStorage, error) {
	contract, err := bindUserStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserStorage{UserStorageCaller: UserStorageCaller{contract: contract}, UserStorageTransactor: UserStorageTransactor{contract: contract}, UserStorageFilterer: UserStorageFilterer{contract: contract}}, nil
}

// NewUserStorageCaller creates a new read-only instance of UserStorage, bound to a specific deployed contract.
func NewUserStorageCaller(address common.Address, caller bind.ContractCaller) (*UserStorageCaller, error) {
	contract, err := bindUserStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserStorageCaller{contract: contract}, nil
}

// NewUserStorageTransactor creates a new write-only instance of UserStorage, bound to a specific deployed contract.
func NewUserStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*UserStorageTransactor, error) {
	contract, err := bindUserStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserStorageTransactor{contract: contract}, nil
}

// NewUserStorageFilterer creates a new log filterer instance of UserStorage, bound to a specific deployed contract.
func NewUserStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*UserStorageFilterer, error) {
	contract, err := bindUserStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserStorageFilterer{contract: contract}, nil
}

// bindUserStorage binds a generic wrapper to an already deployed contract.
func bindUserStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UserStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserStorage *UserStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserStorage.Contract.UserStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserStorage *UserStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserStorage.Contract.UserStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserStorage *UserStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserStorage.Contract.UserStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserStorage *UserStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserStorage *UserStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserStorage *UserStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserStorage.Contract.contract.Transact(opts, method, params...)
}

// GetSender is a free data retrieval call binding the contract method 0x5e01eb5a.
//
// Solidity: function getSender() view returns(address)
func (_UserStorage *UserStorageCaller) GetSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserStorage.contract.Call(opts, &out, "getSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSender is a free data retrieval call binding the contract method 0x5e01eb5a.
//
// Solidity: function getSender() view returns(address)
func (_UserStorage *UserStorageSession) GetSender() (common.Address, error) {
	return _UserStorage.Contract.GetSender(&_UserStorage.CallOpts)
}

// GetSender is a free data retrieval call binding the contract method 0x5e01eb5a.
//
// Solidity: function getSender() view returns(address)
func (_UserStorage *UserStorageCallerSession) GetSender() (common.Address, error) {
	return _UserStorage.Contract.GetSender(&_UserStorage.CallOpts)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _userAddress) view returns(string)
func (_UserStorage *UserStorageCaller) GetUser(opts *bind.CallOpts, _userAddress common.Address) (string, error) {
	var out []interface{}
	err := _UserStorage.contract.Call(opts, &out, "getUser", _userAddress)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _userAddress) view returns(string)
func (_UserStorage *UserStorageSession) GetUser(_userAddress common.Address) (string, error) {
	return _UserStorage.Contract.GetUser(&_UserStorage.CallOpts, _userAddress)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _userAddress) view returns(string)
func (_UserStorage *UserStorageCallerSession) GetUser(_userAddress common.Address) (string, error) {
	return _UserStorage.Contract.GetUser(&_UserStorage.CallOpts, _userAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UserStorage *UserStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UserStorage *UserStorageSession) Owner() (common.Address, error) {
	return _UserStorage.Contract.Owner(&_UserStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UserStorage *UserStorageCallerSession) Owner() (common.Address, error) {
	return _UserStorage.Contract.Owner(&_UserStorage.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_UserStorage *UserStorageTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _UserStorage.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_UserStorage *UserStorageSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _UserStorage.Contract.ChangeOwner(&_UserStorage.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_UserStorage *UserStorageTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _UserStorage.Contract.ChangeOwner(&_UserStorage.TransactOpts, _newOwner)
}

// SetUser is a paid mutator transaction binding the contract method 0x7dc6b859.
//
// Solidity: function setUser(address _userAddress, string _dirHash, bytes _sig) returns()
func (_UserStorage *UserStorageTransactor) SetUser(opts *bind.TransactOpts, _userAddress common.Address, _dirHash string, _sig []byte) (*types.Transaction, error) {
	return _UserStorage.contract.Transact(opts, "setUser", _userAddress, _dirHash, _sig)
}

// SetUser is a paid mutator transaction binding the contract method 0x7dc6b859.
//
// Solidity: function setUser(address _userAddress, string _dirHash, bytes _sig) returns()
func (_UserStorage *UserStorageSession) SetUser(_userAddress common.Address, _dirHash string, _sig []byte) (*types.Transaction, error) {
	return _UserStorage.Contract.SetUser(&_UserStorage.TransactOpts, _userAddress, _dirHash, _sig)
}

// SetUser is a paid mutator transaction binding the contract method 0x7dc6b859.
//
// Solidity: function setUser(address _userAddress, string _dirHash, bytes _sig) returns()
func (_UserStorage *UserStorageTransactorSession) SetUser(_userAddress common.Address, _dirHash string, _sig []byte) (*types.Transaction, error) {
	return _UserStorage.Contract.SetUser(&_UserStorage.TransactOpts, _userAddress, _dirHash, _sig)
}

// Verify is a paid mutator transaction binding the contract method 0x2dd34f0f.
//
// Solidity: function verify(address _userAddress, string _dirHash, bytes _sig) returns(bool)
func (_UserStorage *UserStorageTransactor) Verify(opts *bind.TransactOpts, _userAddress common.Address, _dirHash string, _sig []byte) (*types.Transaction, error) {
	return _UserStorage.contract.Transact(opts, "verify", _userAddress, _dirHash, _sig)
}

// Verify is a paid mutator transaction binding the contract method 0x2dd34f0f.
//
// Solidity: function verify(address _userAddress, string _dirHash, bytes _sig) returns(bool)
func (_UserStorage *UserStorageSession) Verify(_userAddress common.Address, _dirHash string, _sig []byte) (*types.Transaction, error) {
	return _UserStorage.Contract.Verify(&_UserStorage.TransactOpts, _userAddress, _dirHash, _sig)
}

// Verify is a paid mutator transaction binding the contract method 0x2dd34f0f.
//
// Solidity: function verify(address _userAddress, string _dirHash, bytes _sig) returns(bool)
func (_UserStorage *UserStorageTransactorSession) Verify(_userAddress common.Address, _dirHash string, _sig []byte) (*types.Transaction, error) {
	return _UserStorage.Contract.Verify(&_UserStorage.TransactOpts, _userAddress, _dirHash, _sig)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UserStorage *UserStorageTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _UserStorage.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UserStorage *UserStorageSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UserStorage.Contract.Fallback(&_UserStorage.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UserStorage *UserStorageTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UserStorage.Contract.Fallback(&_UserStorage.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UserStorage *UserStorageTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserStorage.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UserStorage *UserStorageSession) Receive() (*types.Transaction, error) {
	return _UserStorage.Contract.Receive(&_UserStorage.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UserStorage *UserStorageTransactorSession) Receive() (*types.Transaction, error) {
	return _UserStorage.Contract.Receive(&_UserStorage.TransactOpts)
}

// UserStorageSetUserSuccIterator is returned from FilterSetUserSucc and is used to iterate over the raw logs and unpacked data for SetUserSucc events raised by the UserStorage contract.
type UserStorageSetUserSuccIterator struct {
	Event *UserStorageSetUserSucc // Event containing the contract specifics and raw log

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
func (it *UserStorageSetUserSuccIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserStorageSetUserSucc)
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
		it.Event = new(UserStorageSetUserSucc)
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
func (it *UserStorageSetUserSuccIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserStorageSetUserSuccIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserStorageSetUserSucc represents a SetUserSucc event raised by the UserStorage contract.
type UserStorageSetUserSucc struct {
	Sender      common.Address
	UserAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetUserSucc is a free log retrieval operation binding the contract event 0x16e46964c9cae88907d56d563590caa3c362261a8bbdf57e259f3a393ae56a8a.
//
// Solidity: event SetUserSucc(address indexed sender, address indexed userAddress)
func (_UserStorage *UserStorageFilterer) FilterSetUserSucc(opts *bind.FilterOpts, sender []common.Address, userAddress []common.Address) (*UserStorageSetUserSuccIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _UserStorage.contract.FilterLogs(opts, "SetUserSucc", senderRule, userAddressRule)
	if err != nil {
		return nil, err
	}
	return &UserStorageSetUserSuccIterator{contract: _UserStorage.contract, event: "SetUserSucc", logs: logs, sub: sub}, nil
}

// WatchSetUserSucc is a free log subscription operation binding the contract event 0x16e46964c9cae88907d56d563590caa3c362261a8bbdf57e259f3a393ae56a8a.
//
// Solidity: event SetUserSucc(address indexed sender, address indexed userAddress)
func (_UserStorage *UserStorageFilterer) WatchSetUserSucc(opts *bind.WatchOpts, sink chan<- *UserStorageSetUserSucc, sender []common.Address, userAddress []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _UserStorage.contract.WatchLogs(opts, "SetUserSucc", senderRule, userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserStorageSetUserSucc)
				if err := _UserStorage.contract.UnpackLog(event, "SetUserSucc", log); err != nil {
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

// ParseSetUserSucc is a log parse operation binding the contract event 0x16e46964c9cae88907d56d563590caa3c362261a8bbdf57e259f3a393ae56a8a.
//
// Solidity: event SetUserSucc(address indexed sender, address indexed userAddress)
func (_UserStorage *UserStorageFilterer) ParseSetUserSucc(log types.Log) (*UserStorageSetUserSucc, error) {
	event := new(UserStorageSetUserSucc)
	if err := _UserStorage.contract.UnpackLog(event, "SetUserSucc", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserStorageVertifySignFailIterator is returned from FilterVertifySignFail and is used to iterate over the raw logs and unpacked data for VertifySignFail events raised by the UserStorage contract.
type UserStorageVertifySignFailIterator struct {
	Event *UserStorageVertifySignFail // Event containing the contract specifics and raw log

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
func (it *UserStorageVertifySignFailIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserStorageVertifySignFail)
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
		it.Event = new(UserStorageVertifySignFail)
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
func (it *UserStorageVertifySignFailIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserStorageVertifySignFailIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserStorageVertifySignFail represents a VertifySignFail event raised by the UserStorage contract.
type UserStorageVertifySignFail struct {
	UserAddress common.Address
	Hash        [32]byte
	Sig         []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVertifySignFail is a free log retrieval operation binding the contract event 0x72fe1ec9576e79584e0bd5228ae442ef5723b5ed7b2d7369ea5d4d10b55ab92e.
//
// Solidity: event VertifySignFail(address indexed userAddress, bytes32 hash, bytes sig)
func (_UserStorage *UserStorageFilterer) FilterVertifySignFail(opts *bind.FilterOpts, userAddress []common.Address) (*UserStorageVertifySignFailIterator, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _UserStorage.contract.FilterLogs(opts, "VertifySignFail", userAddressRule)
	if err != nil {
		return nil, err
	}
	return &UserStorageVertifySignFailIterator{contract: _UserStorage.contract, event: "VertifySignFail", logs: logs, sub: sub}, nil
}

// WatchVertifySignFail is a free log subscription operation binding the contract event 0x72fe1ec9576e79584e0bd5228ae442ef5723b5ed7b2d7369ea5d4d10b55ab92e.
//
// Solidity: event VertifySignFail(address indexed userAddress, bytes32 hash, bytes sig)
func (_UserStorage *UserStorageFilterer) WatchVertifySignFail(opts *bind.WatchOpts, sink chan<- *UserStorageVertifySignFail, userAddress []common.Address) (event.Subscription, error) {

	var userAddressRule []interface{}
	for _, userAddressItem := range userAddress {
		userAddressRule = append(userAddressRule, userAddressItem)
	}

	logs, sub, err := _UserStorage.contract.WatchLogs(opts, "VertifySignFail", userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserStorageVertifySignFail)
				if err := _UserStorage.contract.UnpackLog(event, "VertifySignFail", log); err != nil {
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

// ParseVertifySignFail is a log parse operation binding the contract event 0x72fe1ec9576e79584e0bd5228ae442ef5723b5ed7b2d7369ea5d4d10b55ab92e.
//
// Solidity: event VertifySignFail(address indexed userAddress, bytes32 hash, bytes sig)
func (_UserStorage *UserStorageFilterer) ParseVertifySignFail(log types.Log) (*UserStorageVertifySignFail, error) {
	event := new(UserStorageVertifySignFail)
	if err := _UserStorage.contract.UnpackLog(event, "VertifySignFail", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
