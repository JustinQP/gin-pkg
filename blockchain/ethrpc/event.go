package ethrpc

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (t *Web3RPC) QueryLogsAt(query ethereum.FilterQuery) ([]types.Log, error) {
	logs, err := t.Client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (t *Web3RPC) NewFilterQuery(address string, fromBlock, toBlock *big.Int, topics [][]interface{}) (*ethereum.FilterQuery, error) {
	// convert topics type
	ts := make([][]common.Hash, len(topics))
	for k1, v1 := range topics {
		topic := make([]common.Hash, len(v1))
		for k2, v2 := range v1 {
			switch data := v2.(type) {
			case string:
				topic[k2] = common.HexToHash(data)
			case big.Int:
				topic[k2] = common.BigToHash(&data)
			case []byte:
				topic[k2] = common.BytesToHash(data)
			default:
				return nil, fmt.Errorf("undefined type")
			}
		}
		ts[k1] = topic
	}

	result := &ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Topics:    ts,
		Addresses: []common.Address{
			common.HexToAddress(address),
		},
	}

	return result, nil
}
