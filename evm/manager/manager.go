package evmManager

import (
	"context"
	"fmt"

	UserStorage "mtv/contracts"
	evmCommon "mtv/evm/common"
	evmUserStorage "mtv/evm/user_storage"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmMgr struct {
	EvmParamater   *evmCommon.EvmParameter
	EvmUserStorage *evmUserStorage.EvmUserStorage
}

func NewEvmManager(rpcURL string, walletPassword string, walletPath string) (*EvmMgr, error) {
	var err error
	result := &EvmMgr{
		EvmUserStorage: nil,
	}

	privateKey, err := evmCommon.ReadWallet(walletPassword, walletPath)
	if err != nil {
		return nil, err
	}

	ethClient, err := ethclient.DialContext(context.Background(), rpcURL)
	if err != nil {
		return nil, err
	}

	result.EvmParamater = evmCommon.SetGlobalEvmParamater(privateKey, rpcURL, ethClient, walletPath)
	return result, err
}

func (this *EvmMgr) SetUserStorageContractAddress(userStorageContractAddress string) (*evmUserStorage.EvmUserStorage, error) {
	var err error
	this.EvmParamater.UserStorageContractAddress = userStorageContractAddress
	this.EvmUserStorage, _, err = evmUserStorage.NewEvmUserStorage(nil)
	return this.EvmUserStorage, err
}

func (this *EvmMgr) DeployUserStorageContract() (string, string, *UserStorage.UserStorage, error) {
	auth, err := evmCommon.GlobalEvmParameter.NewTransaction()

	if err != nil {
		return "", "", nil, err
	}

	contractAddress, tx, contractInstance, err := UserStorage.DeployUserStorage(auth, evmCommon.GlobalEvmParameter.EthClient)
	this.EvmParamater.UserStorageContractAddress = contractAddress.Hex()
	this.EvmUserStorage.ContractInstance = contractInstance
	if err != nil {
		return "", "", nil, err
	}

	fmt.Println("userStorage conract contractAddress:", contractAddress.Hex())
	fmt.Println("userStorage conract deploy tx:", tx.Hash().Hex())

	return contractAddress.Hex(), tx.Hash().Hex(), contractInstance, nil
}
