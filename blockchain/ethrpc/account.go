package ethrpc

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func IsPrivKey(privateKey string) (string, bool) {
	if privateKey == "" {
		return "", false
	}

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", false
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", false
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return fromAddress.String(), true
}

func (t *Web3RPC) LoadAccount(privateKey string) error {
	if privateKey == "" {
		return nil
	}

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return err
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	t.Account = &fromAddress
	t.PrivKey = privKey
	return nil
}
