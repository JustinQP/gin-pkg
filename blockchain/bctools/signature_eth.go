package bctools

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

const EthSignStr = "Ethereum Signed Message:"

func EthSign(privKey string, msg string) (string, error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", err
	}

	msg = fmt.Sprintf("\x19%s\n%d%s", EthSignStr, len(msg), msg)
	dataHash := crypto.Keccak256Hash([]byte(msg))
	sig, err := crypto.Sign(dataHash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(sig), nil
}

func EthVerifySignature(pubKey string, msg string, sign string) bool {
	// dataHash := sha256.Sum256([]byte(msg))
	dataHash := crypto.Keccak256Hash([]byte(msg))

	sign = strings.TrimPrefix(sign, "0x")
	signature, err := hex.DecodeString(sign)
	if err != nil {
		return false
	}

	pubkey, err := hex.DecodeString(pubKey)
	if err != nil {
		return false
	}
	return crypto.VerifySignature(pubkey, dataHash[:], signature[:len(signature)-1])
}

func EthVerifySignAddress(address string, msg string, sign string) bool {
	sign = strings.TrimPrefix(sign, "0x")

	if len(sign) != 130 {
		return false
	}

	msg = fmt.Sprintf("\u0019%s\n%d%s", EthSignStr, len(msg), msg)
	dataHash := crypto.Keccak256Hash([]byte(msg))

	signature, err := hex.DecodeString(sign)
	if err != nil {
		return false
	}

	if signature[64] >= 27 {
		signature[64] -= 27
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(dataHash.Bytes(), signature)
	if err != nil {
		return false
	}

	addr := crypto.PubkeyToAddress(*sigPublicKeyECDSA)
	return strings.EqualFold(addr.Hex(), address)
}
