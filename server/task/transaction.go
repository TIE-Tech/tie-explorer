package task

import (
	"encoding/hex"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	base "tie-explorer-go/common"
	"time"

	"tie-explorer-go/models"
	"tie-explorer-go/rpc"
	"unicode"
)

type TransactionRelation struct {
	Timestamp       uint64
	TransactionHash common.Hash
}

var (
	fetchingTransaction bool
	addressMap          = make(map[string]common.Address)
	TransactionChan     = make(chan TransactionRelation)
	TransactionMap      = make(map[common.Hash]uint64)
)

var (
	ErrHasTxFetching = errors.New("the tx info is being processed")
)

type token struct {
	Decimals    string
	TotalSupply string
	Symbol      string
	Name        string
}

func parseTransactionInfoByHash(hash common.Hash, timestamp uint64) {
	delete(TransactionMap, hash)

	c := rpc.GetClient()
	o := orm.NewOrm()

	tx, pending, err := c.GetTxInfoByHash(hash)
	if err != nil {
		logs.Debug("tx hash: %s", hash.Hex())
		logs.Error("get tx info err: %s", err.Error())
		//TransactionMap[tx.Hash()] = timestamp
		return
	}

	msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()), nil)
	if err != nil {
		logs.Error("get tx message error, tx hash: %s", hash.Hex())
		return
	}

	v, r, s := tx.RawSignatureValues()

	tModel := &models.Transaction{
		Hash:      tx.Hash().Hex(),
		From:      msg.From().Hex(),
		Value:     tx.Value().String(),
		Nonce:     tx.Nonce(),
		GasPrice:  tx.GasPrice().Uint64(),
		Gas:       tx.Gas(),
		Timestamp: timestamp,
		DateTime:  time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05"),
		Input:     "0x" + hex.EncodeToString(tx.Data()),
		V:         "0x" + v.Text(16),
		R:         "0x" + r.Text(16),
		S:         "0x" + s.Text(16),
		Type:      tx.Type(),
		IsFake:    msg.IsFake(),
		Status:    0,
	}

	if msg.From().Hex() != "" {
		addressMap[msg.From().Hex()] = msg.From()
	}

	if tx.To() == nil {
		tModel.To = ""
	} else {
		tModel.To = tx.To().Hex()
		addressMap[tx.To().Hex()] = *tx.To()
	}

	if pending {
		logs.Debug("tx is pending")
		TransactionMap[tx.Hash()] = timestamp
	} else {
		receipt, err := c.GetTxReceiptInfo(hash)
		if err != nil {
			logs.Error("get tx receipt error, hash: %s", hash.Hex())
			TransactionMap[tx.Hash()] = timestamp
			return
		}

		tModel.Index = uint64(receipt.TransactionIndex)
		tModel.BlockHash = receipt.BlockHash.Hex()
		tModel.BlockNumber = receipt.BlockNumber.Uint64()
		tModel.GasUsed = receipt.GasUsed
		tModel.CumulativeGasUsed = receipt.CumulativeGasUsed
		tModel.Root = "0x" + hex.EncodeToString(receipt.PostState)
		tModel.Status = receipt.Status
		tModel.LogsBoom = "0x" + hex.EncodeToString(receipt.Bloom.Bytes())
		tModel.ContractAddress = receipt.ContractAddress.Hex()

		addressMap[receipt.ContractAddress.Hex()] = receipt.ContractAddress

		if len(receipt.Logs) > 0 {
			parseTransactionLogs(receipt.Logs, tx.To())
		}
	}

	_, err = o.InsertOrUpdate(tModel, "hash")
	if err != nil {
		logs.Error("tx info write to DB error, %s", err.Error())
		TransactionMap[tx.Hash()] = timestamp
		return
	}

	for _, addr := range addressMap {
		AddressChan <- addr
	}
}

func parseTransactionLogs(txLogs []*types.Log, to *common.Address) {
	logs.Info("parse the tx logs")
	o := orm.NewOrm()
	//var txLogArr []*models.Log
	for _, l := range txLogs {
		logModel := &models.Log{
			Address:          l.Address.Hex(),
			Data:             "0x" + hex.EncodeToString(l.Data),
			BlockNumber:      l.BlockNumber,
			TransactionHash:  l.TxHash.Hex(),
			TransactionIndex: 0,
			BlockHash:        l.BlockHash.Hex(),
			LogIndex:         uint64(l.Index),
			Removed:          l.Removed,
		}

		addressMap[l.Address.Hex()] = l.Address

		for i, t := range l.Topics {
			// token transfer
			if i == 0 {
				logModel.Topics0 = t.String()
				value := string(l.Data)
				if t.String() == base.TransferSignature {
					fromAddr := "0x" + l.Topics[1].Hex()[26:]
					toAddr := "0x" + l.Topics[2].Hex()[26:]
					addressMap[fromAddr] = common.HexToAddress(fromAddr)
					addressMap[toAddr] = common.HexToAddress(toAddr)
					transferModel := models.TokenTransfer{
						BlockNumber:          l.BlockNumber,
						BlockHash:            l.BlockHash.Hex(),
						TransactionHash:      l.TxHash.Hex(),
						LogIndex:             uint64(l.Index),
						TokenContractAddress: l.Address.Hex(),
						From:                 fromAddr,
						To:                   toAddr,
						Value:                value,
					}
					tokenType := "erc-20"
					tokenId := ""
					if len(l.Topics) == 4 {
						tokenType = "erc-721"
						tokenId = l.Topics[3].Big().String()
						transferModel.TokenId = tokenId
					}
					transferModel.TokenType = tokenType

					o.InsertOrUpdate(&transferModel, "transaction_hash,log_index")

					tokenModel := models.Token{
						ContractAddress: l.Address.Hex(),
					}
					err := o.Read(&tokenModel)
					if err != nil {
						tokenInfo := getTokenInfo(new(big.Int).SetUint64(l.BlockNumber), to)
						tokenModel.Name = tokenInfo.Name
						tokenModel.Symbol = tokenInfo.Symbol
						tokenModel.Decimals = tokenInfo.Decimals
						tokenModel.TotalSupply = tokenInfo.TotalSupply
						tokenModel.Type = tokenType
						o.InsertOrUpdate(&tokenModel, "contract_address")
					}

					tokenBalanceModel := models.AddressTokenBalance{
						Address:         toAddr,
						BlockNumber:     l.BlockNumber,
						ContractAddress: l.Address.Hex(),
						Value:           value,
						TokenId:         tokenId,
						TokenType:       tokenType,
					}
					o.InsertOrUpdate(&tokenBalanceModel, "address,contract_address")

					GetTokenBalance(toAddr, l.Address.Hex())

					addAddressName(l.Address.Hex(), tokenModel.Name)
				}
			} else if i == 1 {
				logModel.Topics1 = t.Hex()
			} else if i == 2 {
				logModel.Topics2 = t.Hex()
			} else if i == 3 {
				logModel.Topics3 = t.Hex()
			}
		}

		o.InsertOrUpdate(logModel, "transaction_hash,log_index")
		//txLogArr = append(txLogArr, logModel)
	}
	//o.InsertMulti(len(txLogArr), txLogArr)
}

func addAddressName(address, name string) {
	addressNameModel := models.AddressName{
		Address:  address,
		Name:     name,
		Primary:  false,
		Metadata: "",
	}
	o := orm.NewOrm()
	o.InsertOrUpdate(&addressNameModel, "address,name")
}

func getTokenInfo(blockNumber *big.Int, to *common.Address) *token {
	var arr = []string{
		base.DecimalsSignature,
		base.SymbolSignature,
		base.TotalSupplySignature,
		base.NameSignature,
	}

	t := token{
		Decimals:    "0",
		TotalSupply: "0",
		Symbol:      "",
		Name:        "",
	}

	c := rpc.GetClient()

	baseMsg := ethereum.CallMsg{
		To: to,
	}

	for i, f := range arr {
		baseMsg.Data = common.FromHex(f)
		result, _ := c.CallContractAtBlockNumber(baseMsg, blockNumber)
		if i == 0 {
			t.Decimals = hex.EncodeToString(result)
		} else if i == 1 {
			t.Symbol = trimZero(string(result))
		} else if i == 2 {
			t.TotalSupply = hex.EncodeToString(result)
		} else if i == 3 {
			t.Name = trimZero(string(result))
		}
	}

	return &t
}

func trimZero(s string) string {
	str := make([]rune, 0, len(s))
	for _, v := range []rune(s) {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) {
			continue
		}

		str = append(str, v)
	}
	return string(str)
}

func fetchTransactionFromMap() error {
	if fetchingTransaction {
		logs.Warn("%s", ErrHasTxFetching.Error())
		return ErrHasTxFetching
	}
	fetchingTransaction = true

	if len(TransactionMap) > 0 {
		for hash, timestamp := range TransactionMap {
			logs.Debug("hash:", hash.Hex())
			parseTransactionInfoByHash(hash, timestamp)
		}
	}

	fetchingTransaction = false
	return nil
}

func NewTransactionTask() {
	go func() {
		for {
			select {
			case info := <-TransactionChan:
				parseTransactionInfoByHash(info.TransactionHash, info.Timestamp)
			}
		}
	}()

	transactionTk := toolbox.NewTask("transaction_tk", "0/10 * * * * *", fetchTransactionFromMap)
	toolbox.AddTask("transaction_fetcher_task", transactionTk)
}
