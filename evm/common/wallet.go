package common

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func ReadWallet(walletPassword string, walletPath string) (*ecdsa.PrivateKey, error) {
	walletData, err := ioutil.ReadFile(walletPath)
	if err != nil {
		return nil, err
	}
	privateKeyData, err := keystore.DecryptKey(walletData, walletPassword)
	if err != nil {
		return nil, err
	}
	// privateKey := crypto.FromECDSA(key.PrivateKey)
	privateKey := privateKeyData.PrivateKey
	pubicKey := crypto.FromECDSAPub(&privateKey.PublicKey)
	walletAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	fmt.Println("walletPrivateKey:", hexutil.Encode(crypto.FromECDSA(privateKey)))
	fmt.Println("walletPublicKey:", hexutil.Encode(pubicKey))
	fmt.Println("walletAddress:", walletAddress.Hex())

	return privateKey, nil
}

func NewWallet(walletPath string, walletPassword string) (*accounts.Account, error) {
	// defaultWalletPath = "./wallet"
	// defaultWalletPassword := "Sp9%6mH9XRHr"
	key := keystore.NewKeyStore(walletPath, keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := key.NewAccount(walletPassword)
	if err != nil {
		return nil, err
	}
	fmt.Println("accountAddress:", account.Address)
	return &account, err
}
