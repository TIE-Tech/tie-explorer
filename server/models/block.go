package models

import (
	"github.com/astaxie/beego/orm"
)

type Block struct {
	ParentHash       string  `json:"parentHash" orm:"size(128)"`
	Sha3Uncles       string  `json:"sha3Uncles" orm:"size(128)"`
	Miner            string  `json:"miner" orm:"size(128)"`
	StateRoot        string  `json:"stateRoot" orm:"size(128)"`
	TransactionsRoot string  `json:"transactionsRoot" orm:"size(128)"`
	ReceiptsRoot     string  `json:"receiptsRoot" orm:"size(128)"`
	LogsBloom        string  `json:"logsBloom"`
	Difficulty       uint64  `json:"difficulty"`
	TotalDifficulty  uint64  `json:"totalDifficulty"`
	SizeDecode       string  `json:"sizeDecode" orm:"size(128)"`
	Size             float64 `json:"size"`
	Number           uint64  `json:"number" orm:"unique"`
	GasLimit         uint64  `json:"gasLimit"`
	GasUsed          uint64  `json:"gasUsed"`
	Timestamp        uint64  `json:"timestamp"`
	DateTime         string  `json:"dateTime" orm:"size(128)"`
	ExtraData        string  `json:"extraData"`
	MixHash          string  `json:"mixHash" orm:"size(128)"`
	Nonce            uint64  `json:"nonce"`
	Hash             string  `json:"hash" orm:"unique;size(128)"`
	Txs              int64   `json:"txs"`
	Uncles           int64   `json:"uncles"`
	Reward           string  `json:"reward" orm:"size(128)"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(Block))
}
