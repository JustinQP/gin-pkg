package bctools

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func VerifyAddress(addr string) (string, bool) {
	a := common.HexToAddress(addr)
	aStr := a.String()
	return aStr, strings.EqualFold(addr, aStr)
}
