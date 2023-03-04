package evmDemo

import (
	"fmt"
	evmCommon "mtv/evm/common"
	evmManager "mtv/evm/manager"

	"github.com/beego/beego/v2/core/config"
	"github.com/ethereum/go-ethereum/accounts"
)

func SetUser() error {
	result := false
	rpcURL, _ := config.String("wallet::infuraURL")
	walletPassword, _ := config.String("wallet::walletPassword")
	walletPath, _ := config.String("wallet::walletPath")
	evmMgr, err := evmManager.NewEvmManager(rpcURL, walletPassword, walletPath)
	if err != nil {
		return err
	}

	storageContractAddress, _ := config.String("wallet::storageContractAddress")
	evmMgr.SetUserStorageContractAddress(storageContractAddress)

	const testUserAddress = "0x8420A24D450Da2dAB02C21ccEEd78C71a04E0005"
	const testIpfsDir = "0x8420A24D450Da2dAB02C21ccEEd78C71a04E0005"

	testSign := "0x6c724b2fb807bd337b853977f1d6e3f51488dc01955aac5b7b1719a9fa05f0f768263bcd3dff60c0d6d28731af6e9aadee5264ee7b30de63b5274aa7c67e219a1c"
	testSign, err = evmMgr.EvmUserStorage.Sign(testUserAddress, testIpfsDir)

	result, err = evmMgr.EvmUserStorage.SetUser(testUserAddress, testIpfsDir, testSign)
	if err != nil {
		return err
	}
	fmt.Printf("userStorage contract setUser result:(%v)\n", result)
	return err
}

func GetUser() error {
	rpcURL, _ := config.String("wallet::infuraURL")
	walletPassword, _ := config.String("wallet::walletPassword")
	walletPath, _ := config.String("wallet::walletPath")
	evmMgr, err := evmManager.NewEvmManager(rpcURL, walletPassword, walletPath)
	if err != nil {
		return err
	}

	storageContractAddress, _ := config.String("wallet::storageContractAddress")
	evmMgr.SetUserStorageContractAddress(storageContractAddress)

	const testUserAddress = "0x8420A24D450Da2dAB02C21ccEEd78C71a04E0005"
	ipfsDir, err := evmMgr.EvmUserStorage.GetUser(testUserAddress)
	if err != nil {
		return err
	}
	fmt.Printf("userStorage contract UserAddress:(%v):IpfsDir(%v)\n", testUserAddress, ipfsDir)

	return err
}

func CreateWallet() (*accounts.Account, error) {
	walletPassword, _ := config.String("wallet::walletPassword")
	walletPath, _ := config.String("wallet::walletPath")
	account, err := evmCommon.NewWallet(walletPath, walletPassword)
	if err != nil {
		return account, err
	}
	fmt.Println("userStorageConract contractAddress:", account.Address.Hex())
	return account, err
}

func GetOwner() (string, error) {
	rpcURL, _ := config.String("wallet::infuraURL")
	walletPassword, _ := config.String("wallet::walletPassword")
	walletPath, _ := config.String("wallet::walletPath")
	evmMgr, err := evmManager.NewEvmManager(rpcURL, walletPassword, walletPath)

	if err != nil {
		return "", err
	}

	ownerAddress, err := evmMgr.EvmUserStorage.GetOwner()
	if err != nil {
		return "", err
	}
	fmt.Println("userStorageConract owner's address:", ownerAddress)
	return ownerAddress, err
}

func Deploy() error {
	rpcURL, _ := config.String("wallet::infuraURL")
	walletPassword, _ := config.String("wallet::walletPassword")
	walletPath, _ := config.String("wallet::walletPath")
	evmMgr, err := evmManager.NewEvmManager(rpcURL, walletPassword, walletPath)

	contractAddress, tx, _, err := evmMgr.DeployUserStorageContract()
	if err != nil {
		return err
	}

	fmt.Println("userStorageConract contractAddress:", contractAddress)
	fmt.Println("userStorageConract deploy tx:", tx)
	return err
}

// need web socket
func WatchEvent() (bool, error) {

	// privateKey, _, _ := evmCommon.ReadWallet(password, walletPath)
	// evmMgr, err := evmManager.NewEvmManager(evmCommon.InfuraMumbaiURL, evmCommon.DefaultWalletPassword, evmCommon.DefaultWalletPath, evmCommon.DefaultUserStorageContractAddress)

	// contractAddressBytes, _ := hexutil.Decode(userStorageContractAddress)
	// contractAddress := common.BytesToAddress(contractAddressBytes)

	// contractInstance := newUserStorage(client, contractAddress)

	// userAddressesList := []common.Address{
	// 	contractAddress,
	// }
	// ch := make(chan *UserStorage.UserStorageIsValidSignatureNow)
	// _, err1 := contractInstance.UserStorageFilterer.WatchIsValidSignatureNow(nil, ch, userAddressesList)
	// if err1 != nil {
	// 	panic("[*] Error subscribing to events")
	// }
	// var newEvent *UserStorage.UserStorageIsValidSignatureNow = <-ch
	// fmt.Printf("[*] Event Data Received:(%v):", newEvent)
	return false, nil
}
