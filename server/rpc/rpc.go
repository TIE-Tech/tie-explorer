package rpc

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type rpc struct {
	Client *ethclient.Client
}

var o = &rpc{
	Client: nil,
}

func GetClient() *rpc {
	if o.Client != nil {
		return o
	}

	client, err := ethclient.Dial(beego.AppConfig.String("jsonRpcUrl"))

	if err != nil {
		panic("RPC Client connect error")
	}

	defer client.Close()

	o.Client = client
	return o
}

func (r rpc) GetAccountBalance(address common.Address) *big.Int {
	balance, err := r.Client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return big.NewInt(0)
	}
	return balance
}

func (r rpc) GetTxCountByAddress(address common.Address) uint64 {
	balance, err := r.Client.NonceAt(context.Background(), address, nil)
	if err != nil {
		return 0
	}
	return balance
}

func (r rpc) GetLatestBlockNumber() uint64 {
	number, err := r.Client.BlockNumber(context.Background())

	if err != nil {
		return 0
	}

	return number
}

func (r rpc) GetBlockInfoByNumber(number *big.Int) (*types.Block, error) {
	block, err := r.Client.BlockByNumber(context.Background(), number)
	if err != nil {
		return nil, fmt.Errorf("get block: %s error, %s", number.String(), err.Error())
	}
	return block, nil
}

func (r rpc) GetBlockInfoByHash(hash common.Hash) (*types.Block, error) {
	block, err := r.Client.BlockByHash(context.Background(), hash)
	if err != nil {
		return nil, fmt.Errorf("get HASH: %s error, %s", hash.Hex(), err.Error())
	}
	return block, nil
}

func (r rpc) GetTxInfoByHash(hash common.Hash) (*types.Transaction, bool, error) {
	tx, pending, err := r.Client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return nil, false, fmt.Errorf("get tx error, %s", err.Error())
	}
	return tx, pending, nil
}

func (r rpc) GetTxReceiptInfo(hash common.Hash) (*types.Receipt, error) {
	receipt, err := r.Client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return nil, fmt.Errorf("get Receipt error, %s", err.Error())
	}
	return receipt, nil
}

func (r rpc) GetTxInfoByBlockNumberAndIndex(hash common.Hash, index uint) (*types.Transaction, error) {
	tx, err := r.Client.TransactionInBlock(context.Background(), hash, index)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	return tx, nil
}

func (r rpc) GetTxCountByBlockHash(hash common.Hash) (uint, error) {
	count, err := r.Client.TransactionCount(context.Background(), hash)
	if err != nil {
		return 0, fmt.Errorf("%s", err.Error())
	}
	return count, nil
}

func (r rpc) GetChainId() *big.Int {
	b, err := r.Client.ChainID(context.Background())
	if err != nil {
		return big.NewInt(0)
	}
	return b
}

func (r rpc) GetNetVersion() uint64 {
	// net_version
	b, err := r.Client.NetworkID(context.Background())
	if err != nil {
		return 0
	}
	return b.Uint64()
}

func (r rpc) GetLogs(query ethereum.FilterQuery) ([]types.Log, error) {
	l, err := r.Client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	return l, nil
}

func (r rpc) GetCode(account common.Address, number *big.Int) ([]byte, error) {
	c, err := r.Client.CodeAt(context.Background(), account, number)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	return c, nil
}

func (r rpc) GetPendingTxCount() uint {
	c, err := r.Client.PendingTransactionCount(context.Background())
	if err != nil {
		return 0
	}
	return c
}

func (r rpc) CallContractAtBlockNumber(msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	bytes, err := r.Client.CallContract(context.Background(), msg, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return bytes, nil
}

func (r rpc) CallContractAtBlockHash(msg ethereum.CallMsg, hash common.Hash) ([]byte, error) {
	//to := common.HexToAddress("0x8d338C427a746ADbbB06fDA2fe23A9f5e81587B2")
	//
	//msg := ethereum.CallMsg{
	//	From: common.HexToAddress("0x5573143eBA235545BE5548E1107dD7B92713EF18"),
	//	To:   &to,
	//}

	bytes, err := r.Client.CallContractAtHash(context.Background(), msg, hash)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return bytes, nil
}
