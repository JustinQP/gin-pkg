package ethrpc

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (t *Web3RPC) LoadContract(address, ccabi, privateKey string) error {
	if !common.IsHexAddress(address) {
		return fmt.Errorf("invalid address type")
	}
	ccAddress := common.HexToAddress(address)

	// import ABI interface string
	parsedABI, err := abi.JSON(strings.NewReader(ccabi))
	if err != nil {
		return fmt.Errorf("error parsing abi data:%s", err)
	}

	t.Abi = &parsedABI
	t.CcAddress = ccAddress

	// generate contract instances
	t.Contract = bind.NewBoundContract(ccAddress, parsedABI, t.Client, t.Client, nil)

	err = t.LoadAccount(privateKey)
	if err != nil {
		return fmt.Errorf("invalid private key, read-only")
	}
	return nil
}

func (t *Web3RPC) ContractCall(opts *bind.CallOpts, results *[]interface{}, method string, params ...interface{}) error {
	return t.Contract.Call(opts, results, method, params)
}

// gasPrice If nil use the suggested value (Web3 web query)
// gasLimit If nil use the suggested value (Web3 web query)
func (t *Web3RPC) TxOpts(gasPrice, value *big.Int, gasLimit, nonce uint64) (*bind.TransactOpts, error) {
	var err error

	if nonce == 0 {
		nonce, err = t.Client.PendingNonceAt(context.Background(), *t.Account)
		if err != nil {
			return nil, fmt.Errorf("PendingNonceAt: %s", err)
		}
	}

	if gasPrice == nil {
		gasPrice, err = t.Client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, fmt.Errorf("SuggestGasPrice: %s", err)
		}
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(t.PrivKey, big.NewInt(CHAIN_ID))
	if err != nil {
		return nil, fmt.Errorf("NewKeyedTransactorWithChainID: %s", err)
	}

	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.Value = value // in wei

	transactOpts.GasLimit = gasLimit // in units
	transactOpts.GasPrice = gasPrice

	// transactOpts.NoSend = true

	return transactOpts, nil
}
