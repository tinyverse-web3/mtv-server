package storage

import (
	"context"
	"fmt"
	UserStorage "mtv/contracts"
	evmCommon "mtv/evm/common"
	"strconv"
	"time"

	"github.com/beego/beego/v2/core/config"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

type EvmUserStorage struct {
	ContractInstance *UserStorage.UserStorage
}

func NewEvmUserStorage(contractInstance *UserStorage.UserStorage) (*EvmUserStorage, *UserStorage.UserStorage, error) {
	var err error
	result := &EvmUserStorage{
		ContractInstance: contractInstance,
	}
	instance, err := result.getInstance()
	if err != nil {
		return nil, nil, err
	}
	return result, instance, nil
}

func (this *EvmUserStorage) getInstance() (*UserStorage.UserStorage, error) {
	if this.ContractInstance != nil {
		return this.ContractInstance, nil
	}

	contractAddressObject := common.HexToAddress(evmCommon.GlobalEvmParameter.UserStorageContractAddress)

	var err error
	this.ContractInstance, err = UserStorage.NewUserStorage(contractAddressObject, evmCommon.GlobalEvmParameter.EthClient)
	if err != nil {
		return nil, err
	}
	return this.ContractInstance, nil
}

func (this *EvmUserStorage) GetOwner() (string, error) {
	owner, err := this.ContractInstance.Owner(nil)
	if err != nil {
		return "", err
	}
	fmt.Println("userStorage contract owner's address:", owner.Hex())
	return owner.Hex(), nil
}

func (this *EvmUserStorage) Sign(userAddress string, hashDir string) (string, error) {
	hashData := solsha3.SoliditySHA3(
		[]string{"address", "string"},
		[]interface{}{
			userAddress,
			hashDir,
		},
	)
	fmt.Println("sign hashData:", hexutil.Encode(hashData))

	privateKeyStr := "0x66234d94782b104f840cdfc64f5ae5164e337bb960793a5e80ca15a70d313237"

	privateKeyByte, err := hexutil.Decode(privateKeyStr)
	if err != nil {
		return "", err
	}
	privateKeyData, err := crypto.ToECDSA(privateKeyByte)
	if err != nil {
		return "", err
	}

	signature, err := crypto.Sign(hashData, privateKeyData)
	if err != nil {
		return "", err
	}
	if signature[64] == 0 || signature[64] == 1 {
		signature[64] += 27
	}

	keccak256Hash := hexutil.Encode(signature)
	fmt.Println("EvmUserStorage signature: ", keccak256Hash)
	return keccak256Hash, nil
}

func (this *EvmUserStorage) GetSender() (string, error) {
	senderAddressObj, err := this.ContractInstance.GetSender(nil)
	if err != nil {
		return "", err
	}
	fmt.Printf("userStorage contract senderAddress:%v\n", senderAddressObj.Hex())

	return senderAddressObj.Hex(), err
}

func (this *EvmUserStorage) GetUser(userAddress string) (string, error) {
	userAddressObject := common.HexToAddress(userAddress)
	ipfsDir, err := this.ContractInstance.GetUser(nil, userAddressObject)
	if err != nil {
		return "", err
	}
	fmt.Printf("userStorage contract UserAddress:(%v) <=> IpfsDir(%v)\n", userAddress, ipfsDir)
	return ipfsDir, err
}

func (this *EvmUserStorage) SetUser(userAddress string, ipfsDir string, sign string) (bool, error) {
	result := false

	signData, err := hexutil.Decode(sign)
	signDataHash := common.BytesToHash(signData)
	fmt.Println("userStorage conract sign hash:", signDataHash.Hex())
	// signData, err := this.Sign(userAddress, ipfsDir)
	if err != nil {
		return result, err
	}

	auth, err := evmCommon.GlobalEvmParameter.NewTransaction()
	if err != nil {
		return result, err
	}

	userAddressObject := common.HexToAddress(userAddress)

	// tx, err := contractInstance.Verify(auth, userAddress, ipfsDir, sig)
	tx, err := this.ContractInstance.SetUser(auth, userAddressObject, ipfsDir, signData)
	if err != nil {
		return result, err
	}
	fmt.Println("userStorage contract tx sent:", tx.Hash().Hex())

	// wait for sync data in all block
	tryCount := 0

	var transactionReceipt *types.Receipt
	pullContractLogTryCount, _ := config.Int("wallet::pullContractLogTryCount")
	for transactionReceipt == nil {
		tryCount++
		if tryCount >= pullContractLogTryCount {
			break
		}
		time.Sleep(1 * time.Second)

		transactionReceipt, err = evmCommon.GlobalEvmParameter.EthClient.TransactionReceipt(context.Background(), tx.Hash())
		if err == nil {
			if transactionReceipt.Status == 1 {
				break
			} else {
				err = fmt.Errorf("transactionReceipt.Status(" + strconv.FormatUint(transactionReceipt.Status, 10) + ")!= 1")
				return result, err
			}
		} else {
			if err.Error() != "not found" {
				err = fmt.Errorf("not found transactionReceipt")
				return result, err
			}
		}
	}

	logBlockNumber := transactionReceipt.BlockNumber
	query := ethereum.FilterQuery{
		FromBlock: logBlockNumber,
		ToBlock:   logBlockNumber,
		Addresses: []common.Address{
			common.HexToAddress(evmCommon.GlobalEvmParameter.UserStorageContractAddress),
		},
	}

	logs, err := evmCommon.GlobalEvmParameter.EthClient.FilterLogs(context.Background(), query)
	if err != nil {
		err = fmt.Errorf("not found transactionReceipt")
		return result, err
	}

	for index := range logs {
		log := logs[index]
		setUserEvent, setUserErr := this.ContractInstance.ParseSetUserSucc(log)
		if setUserErr == nil {
			result = true
			fmt.Printf("userStorage contract setUserEvent:(%v)\n", setUserEvent)
		} else {
			signEvent, signErr := this.ContractInstance.UserStorageFilterer.ParseVertifySignFail(log)
			if signErr != nil {
				fmt.Println("userStorage Conract signFailErr:", signErr)
			} else {
				fmt.Printf("userStorage Conract signEvent:(%v)\n", signEvent)
			}
		}
	}
	return result, err
}
