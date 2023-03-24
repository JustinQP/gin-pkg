package ethrpc

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type Web3RPC struct {
	Client    *ethclient.Client
	Contract  *bind.BoundContract
	Account   *common.Address
	PrivKey   *ecdsa.PrivateKey
	Abi       *abi.ABI
	CcAddress common.Address
}

const (
	TX_FAULT   = 0
	TX_SUCCESS = 1
)

var CHAIN_ID int64

const BLOCK_GAP = 1 // 区块间隔

func init() {
	// 获取环境变量 GIN_CHAIN_ID
	if chainId := os.Getenv("GIN_CHAIN_ID"); chainId != "" {
		CHAIN_ID, _ = strconv.ParseInt(chainId, 10, 64)
	} else {
		CHAIN_ID = 1
	}
}

func NewWeb3RPC(web3url string) (*Web3RPC, error) {
	// Initialize an RPC connection
	web3Client, err := ethclient.Dial(web3url)
	if err != nil {
		return nil, err
	}

	w3 := &Web3RPC{
		Client: web3Client,
	}

	return w3, nil
}

// Query the latest account balance, error return 0
func (t *Web3RPC) BalanceAt(address string) (*big.Int, error) {
	balance, err := t.Client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return big.NewInt(0), fmt.Errorf("query account balance err:%s", err.Error())
	}

	return balance, nil
}

func (t *Web3RPC) GetBlockTimeByNumber(height int64) (int64, error) {
	blockNumber := big.NewInt(height)
	block, err := t.Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return 0, err
	}

	return int64(block.Time()), nil
}

// Query pending or completed transactions
func (t *Web3RPC) GetPendingTx(txHash string) (*types.Receipt, error) {
	// 6 * 10 second
	for i := 0; i < 10; i++ {
		tx, isPending, err := t.Client.TransactionByHash(context.Background(), common.HexToHash(txHash))
		if err != nil {
			if errors.Is(err, ethereum.NotFound) {
				time.Sleep(BLOCK_GAP * time.Second)
				continue
			} else {
				return nil, errors.Wrap(err, txHash)
			}
		}

		if isPending {
			time.Sleep(time.Second * BLOCK_GAP)
			continue
		}

		receipt, err := t.Client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			return nil, err
		}

		if receipt.Status == TX_FAULT {
			return nil, fmt.Errorf("tx fault, txHash:%s", txHash)
		}

		return receipt, nil

	}

	return nil, fmt.Errorf("tx not found, txHash:%s", txHash)
}
