package utils

import (
	evm "mtv/evm/manager"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
)

func GetDFSPath(walletAddress string) (string, error) {
	// const walletAddress = "0x8420A24D450Da2dAB02C21ccEEd78C71a04E0005"
	rpcURL, _ := config.String("wallet::infuraURL")
	walletPassword, _ := config.String("wallet::walletPassword")
	walletPath, _ := config.String("wallet::walletPath")
	evmMgr, err := evm.NewEvmManager(rpcURL, walletPassword, walletPath)
	if err != nil {
		return "", err
	}

	storageContractAddress, _ := config.String("wallet::storageContractAddress")
	evmMgr.SetUserStorageContractAddress(storageContractAddress)

	dfsPath, err := evmMgr.EvmUserStorage.GetUser(walletAddress)
	if err != nil {
		return "", err
	}

	return dfsPath, nil
}

func SetDFSPath(walletAddress, dfsPath, sign string) (bool, error) {
	success := false
	rpcURL, _ := config.String("wallet::infuraURL")
	walletPassword, _ := config.String("wallet::walletPassword")
	walletPath, _ := config.String("wallet::walletPath")
	evmMgr, err := evm.NewEvmManager(rpcURL, walletPassword, walletPath)
	if err != nil {
		logs.Error(err)
		return false, err
	}

	storageContractAddress, _ := config.String("wallet::storageContractAddress")
	evmMgr.SetUserStorageContractAddress(storageContractAddress)

	success, err = evmMgr.EvmUserStorage.SetUser(walletAddress, dfsPath, sign)
	if err != nil {
		logs.Error(err)
		return false, err
	}

	return success, nil
}
