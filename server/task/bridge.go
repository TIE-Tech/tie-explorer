package task

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strconv"
	"strings"
)

type BridgeTxRelation struct {
	RPC       string
	Hash      string
	ProcessId int64
}

type BridgeProcessRelation struct {
	RPC          string
	BlockNumber  int64
	ToAddress    string
	ProcessId    int64
	Topic        string
	ContractHash string
	QueryCount   int
}

var (
	BridgeTxMap    = make(map[int64]BridgeTxRelation)
	TxProcessFlag  = false
	BridgeLogMap   = make(map[int64]BridgeProcessRelation)
	logProcessFlag = false
)

var (
	Failed    int64 = -1
	Finished  int64 = 0
	Pending   int64 = 1
	Confirmed int64 = 2
)

func WatchTxHash() error {
	if TxProcessFlag {
		return nil
	}
	TxProcessFlag = true
	for _, item := range BridgeTxMap {
		GetTxInfo(item)
	}
	TxProcessFlag = false
	return nil
}

func WatchLogRecords() error {
	if logProcessFlag {
		return nil
	}
	logProcessFlag = true
	for _, item := range BridgeLogMap {
		CheckLogs(item)
	}
	logProcessFlag = false
	return nil
}

func GetBlockHeight(rpc string) int64 {
	client, _ := ethclient.Dial(rpc)
	defer client.Close()
	n, _ := client.BlockNumber(context.Background())
	return int64(n)
}

func GetTxInfo(relation BridgeTxRelation) {
	delete(BridgeTxMap, relation.ProcessId)

	client, _ := ethclient.Dial(relation.RPC)
	defer client.Close()
	_, pending, err := client.TransactionByHash(context.Background(), common.HexToHash(relation.Hash))
	if pending || err != nil {
		fmt.Println("tx is pending", err)
		BridgeTxMap[relation.ProcessId] = relation
	}
	r, err := client.TransactionReceipt(context.Background(), common.HexToHash(relation.Hash))
	if err != nil {
		BridgeTxMap[relation.ProcessId] = BridgeTxRelation{
			RPC:       relation.RPC,
			Hash:      relation.Hash,
			ProcessId: relation.ProcessId,
		}
	} else {
		o := orm.NewOrm()
		if r.Status == 1 {
			o.Raw("UPDATE \"bridge_process\" SET \"state\"=" + strconv.FormatInt(Confirmed, 10) + ", \"block_number\"=" + r.BlockNumber.String() + " WHERE \"id\"=" + strconv.FormatInt(relation.ProcessId, 10)).Exec()
		} else {
			o.Raw("UPDATE \"bridge_process\" SET \"state\"=" + strconv.FormatInt(Failed, 10) + ", \"block_number\"=" + r.BlockNumber.String() + " WHERE \"id\"=" + strconv.FormatInt(relation.ProcessId, 10)).Exec()
		}
	}
}

func CheckLogs(relation BridgeProcessRelation) {
	delete(BridgeLogMap, relation.ProcessId)

	rpcClient, err := ethclient.Dial(relation.RPC)
	defer rpcClient.Close()
	filterLogs, err := rpcClient.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(relation.BlockNumber),
		ToBlock:   big.NewInt(relation.BlockNumber + 5000),
		Addresses: []common.Address{
			common.HexToAddress(relation.ContractHash),
		},
		Topics: [][]common.Hash{
			{common.HexToHash(relation.Topic)},
		},
	})

	if err != nil {
		BridgeLogMap[relation.ProcessId] = BridgeProcessRelation{
			RPC:          relation.RPC,
			ContractHash: relation.ContractHash,
			Topic:        relation.Topic,
			BlockNumber:  relation.BlockNumber,
			ToAddress:    relation.ToAddress,
			ProcessId:    relation.ProcessId,
			QueryCount:   relation.QueryCount + 1,
		}
	}

	txHash := ""
	for _, item := range filterLogs {
		fmt.Println(strings.ToLower(item.Topics[2].Hex()))
		fmt.Println(strings.ToLower(relation.ToAddress[2:]))
		if strings.Contains(strings.ToLower(item.Topics[2].Hex()), strings.ToLower(relation.ToAddress[2:])) {
			amount, _ := hexutil.DecodeUint64("0x" + strings.TrimLeft(hexutil.Encode(item.Data)[2:], "0"))
			amountBN := new(big.Int).SetUint64(amount)
			decimals := new(big.Int).SetUint64(1000000000000000000)
			fmt.Println("transfer amount:", amountBN.Div(amountBN, decimals))
			txHash = item.TxHash.Hex()
			break
		}
	}
	if txHash != "" {
		o := orm.NewOrm()
		o.Raw("UPDATE \"bridge_process\" SET \"state\"=" + strconv.FormatInt(Confirmed, 10) + ", \"hash\"='" + txHash[2:] + "' WHERE \"id\"=" + strconv.FormatInt(relation.ProcessId, 10)).Exec()
	} else if relation.QueryCount <= 500 {
		fmt.Println("get data from ", len(filterLogs), " log records, Retry the", relation.QueryCount, "th time, retry a maximum of 500.")
		BridgeLogMap[relation.ProcessId] = BridgeProcessRelation{
			RPC:          relation.RPC,
			ContractHash: relation.ContractHash,
			Topic:        relation.Topic,
			BlockNumber:  relation.BlockNumber,
			ToAddress:    relation.ToAddress,
			ProcessId:    relation.ProcessId,
			QueryCount:   relation.QueryCount + 1,
		}
	}
}

func NewBridgeTask() {
	bridgeTxTk := toolbox.NewTask("address_tk", "*/5 * * * * *", WatchTxHash)
	toolbox.AddTask("bridge_tx_fetcher_task", bridgeTxTk)
	bridgeLogTk := toolbox.NewTask("address_tk", "*/5 * * * * *", WatchLogRecords)
	toolbox.AddTask("bridge_log_fetcher_task", bridgeLogTk)
}
