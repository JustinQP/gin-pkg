package bctools

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func Eip712V4Sign(privKey string, domainSeparator, structHash string) (string, error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("\x19\x01%s%s", domainSeparator, structHash)
	dataHash := crypto.Keccak256Hash([]byte(msg))
	sig, err := crypto.Sign(dataHash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(sig), nil
}

func Eip712V4SignVerify(address string, domainSeparator, structHash string, sign string) bool {
	sign = strings.TrimPrefix(sign, "0x")
	if len(sign) != 130 {
		return false
	}

	// 计算 消息的  keccak245 hash
	msg := fmt.Sprintf("\x19\x01%s%s", domainSeparator, structHash)
	dataHash := crypto.Keccak256Hash([]byte(msg))

	// 签名格式 string -> 字节
	signature, err := hex.DecodeString(sign)
	if err != nil {
		return false
	}

	//fmt.Println("signature[64]=", signature[64])
	if signature[64] >= 27 {
		signature[64] -= 27
	}

	// 由 msgHash + 签名 ==> PublicKey
	sigPublicKeyECDSA, err := crypto.SigToPub(dataHash.Bytes(), signature)
	if err != nil {
		return false
	}

	addr := crypto.PubkeyToAddress(*sigPublicKeyECDSA)
	return strings.EqualFold(addr.Hex(), address)
}
