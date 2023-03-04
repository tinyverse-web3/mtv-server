package common

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/beego/beego/v2/core/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var GlobalEvmParameter *EvmParameter

type EvmParameter struct {
	PrivateKey                 *ecdsa.PrivateKey
	RpcURL                     string
	EthClient                  *ethclient.Client
	WalletPath                 string
	UserStorageContractAddress string
}

func SetGlobalEvmParamater(privateKey *ecdsa.PrivateKey, rpcURL string, ethClient *ethclient.Client, walletPath string) *EvmParameter {
	GlobalEvmParameter = &EvmParameter{
		PrivateKey: privateKey,
		RpcURL:     rpcURL,
		EthClient:  ethClient,
		WalletPath: walletPath,
	}
	return GlobalEvmParameter
}

func GetGlobalEvmParamater() *EvmParameter {
	return GlobalEvmParameter
}

func (this *EvmParameter) NewTransaction() (*bind.TransactOpts, error) {

	walletAddress := crypto.PubkeyToAddress(this.PrivateKey.PublicKey)

	nonce, err := this.EthClient.PendingNonceAt(context.Background(), walletAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := this.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := this.EthClient.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(this.PrivateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.GasPrice = gasPrice

	gasLimit, _ := config.Int("wallet::gasLimit")
	auth.GasLimit = uint64(gasLimit)
	auth.Nonce = big.NewInt(int64(nonce))

	return auth, nil
}
