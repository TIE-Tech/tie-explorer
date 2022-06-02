package task

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
	"sync"
	"tie-explorer-go/models"
	"tie-explorer-go/rpc"
	"time"
)

var (
	fetchingBlock    bool
	cacheBlockNumber uint64
	currentBlock     uint64

	maxProcessNum = 120

	ErrorBlockChan = make(chan uint64)
)

var (
	ErrHasBlockFetching = errors.New("the block info is being processed")
)

func getLatestBlockNumberFromDB() uint64 {
	type latest struct {
		Max uint64
	}
	var max latest
	o := orm.NewOrm()
	o.Raw("SELECT \"number\" AS max FROM block ORDER BY \"number\" DESC LIMIT 1").QueryRow(&max)

	return max.Max
}

func parseBlockInfoByNumber(number uint64) {
	logs.Info("get [%d] block info", number)
	c := rpc.GetClient()

	block, err := c.GetBlockInfoByNumber(new(big.Int).SetUint64(number))
	if err != nil {
		fmt.Println("get block info err, ", err.Error())
		return
	}

	var reward = "0"
	if block.Transactions().Len() > 0 {
		for _, item := range block.Transactions() {
			if hex.EncodeToString(item.Data()) == "57305920" &&
				strings.ToLower(item.To().Hex()) == "0xc79543f253dbf1f7606499be536620c1b1358e1c" {
				reward = item.Value().String()
			}
			go func(hash common.Hash) {
				TransactionChan <- TransactionRelation{
					Timestamp:       block.Time(),
					TransactionHash: hash,
				}
			}(item.Hash())
		}
	}

	blockModel := &models.Block{
		ParentHash:       block.ParentHash().Hex(),
		Sha3Uncles:       block.UncleHash().Hex(),
		Miner:            block.Coinbase().Hex(),
		StateRoot:        block.Root().Hex(),
		TransactionsRoot: block.TxHash().Hex(),
		ReceiptsRoot:     block.ReceiptHash().Hex(),
		LogsBloom:        "0x" + hex.EncodeToString(block.Bloom().Bytes()),
		Difficulty:       block.Difficulty().Uint64(),
		TotalDifficulty:  block.Difficulty().Uint64(),
		SizeDecode:       block.Size().String(),
		Size:             float64(block.Size()),
		Number:           block.Number().Uint64(),
		GasLimit:         block.GasLimit(),
		GasUsed:          block.GasUsed(),
		Timestamp:        block.Time(),
		DateTime:         time.Unix(int64(block.Time()), 0).Format("2006-01-02 15:04:05"),
		ExtraData:        "0x" + hex.EncodeToString(block.Extra()),
		MixHash:          block.MixDigest().Hex(),
		Nonce:            block.Nonce(),
		Hash:             block.Hash().Hex(),
		Txs:              int64(block.Transactions().Len()),
		Uncles:           int64(len(block.Uncles())),
		Reward:           reward,
	}

	o := orm.NewOrm()
	_, err = o.InsertOrUpdate(blockModel, "hash")
	if err != nil {
		logs.Error("block info write to DB error, ", err.Error())
		ErrorBlockChan <- number
		return
	}

	go func() {
		AddressChan <- block.Coinbase()
	}()
}

func fetchNewestBlocks() error {
	if fetchingBlock {
		logs.Warn("%s", ErrHasBlockFetching.Error())
		return ErrHasBlockFetching
	}

	fetchingBlock = true

	c := rpc.GetClient()
	latestBlockNumber := c.GetLatestBlockNumber()

	if currentBlock <= 0 {
		currentBlock = latestBlockNumber
		cacheBlockNumber = latestBlockNumber
		dbBlockNumber := getLatestBlockNumberFromDB()
		go fetchLostBlocks(dbBlockNumber)
	}

	if currentBlock+1 <= latestBlockNumber {
		for n := currentBlock + 1; n <= latestBlockNumber; n++ {
			parseBlockInfoByNumber(n)
		}
	} else {
		for n := currentBlock - 3; n <= latestBlockNumber; n++ {
			parseBlockInfoByNumber(n)
		}
	}

	currentBlock = latestBlockNumber
	fetchingBlock = false
	return nil
}

func fetchLostBlocks(dbBlockNumber uint64) {
	logs.Info("processing the missing blocks")

	ticker := time.NewTicker(1 * time.Second / time.Duration(maxProcessNum))
	defer ticker.Stop()

	var wg sync.WaitGroup
	for i := cacheBlockNumber; i > dbBlockNumber; i-- {
		<-ticker.C
		wg.Add(1)

		go func(n uint64) {
			defer wg.Done()
			parseBlockInfoByNumber(n)
		}(i)
	}
	wg.Wait()
}

func NewBlockTask() {
	go func() {
		for {
			select {
			case n := <-ErrorBlockChan:
				time.Sleep(time.Second * 2)
				parseBlockInfoByNumber(n)
			}
		}
	}()

	blockTk := toolbox.NewTask("block_tk", "0/"+beego.AppConfig.String("interval")+" * * * * *", fetchNewestBlocks)
	toolbox.AddTask("block_fetcher_task", blockTk)
}

func DebugTask(number uint64) {
	parseBlockInfoByNumber(number)
}
