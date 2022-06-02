package common

import (
	"github.com/astaxie/beego/orm"
	"tie-explorer-go/models"
)

func AddTestTokens() {
	tokenBasic := models.BridgeAsset{
		Name:      "USDT",
		Precision: 18,
		Meta:      "/icon/usdt.svg",
		State:     true,
	}

	tokenTie := models.BridgeToken{
		Hash:           "B5fabF404a12687ACe6d6f524825AC2df4B49e9E",
		ChainId:        200,
		Name:           "USDT",
		TokenBasicName: "USDT",
		Property:       1,
		Precision:      18,
		ResourceId:     "0x000000000000000000000000000000c76ebe4a02bbc34786d860b355f5a5ce00",
		State:          true,
	}

	tokenBnb := models.BridgeToken{
		Hash:           "52Ad3084a9Ad4152E52FdEAb8644a05E9Eb1E006",
		ChainId:        97,
		Name:           "USDT",
		TokenBasicName: "USDT",
		Property:       1,
		Precision:      18,
		ResourceId:     "0x000000000000000000000000000000c76ebe4a02bbc34786d860b355f5a5ce00",
		State:          true,
	}

	tokenxTieBasic := models.BridgeAsset{
		Name:      "xTIE",
		Precision: 18,
		Meta:      "/icon/polygon.svg",
		State:     true,
	}

	tokenXTieTie := models.BridgeToken{
		Hash:           "209514Bd4B528e8248cDe4B2CD5C023F47746E38",
		ChainId:        200,
		Name:           "xTIE",
		TokenBasicName: "xTIE",
		Property:       1,
		Precision:      18,
		ResourceId:     "0x000000000000000000000000000000c76ebe4a02bbc34786d860b355f5a5ce02",
		State:          true,
	}

	tokenxTieBnb := models.BridgeToken{
		Hash:           "10D95320F84e96eDd3C69347Ebcf83a6B7fae0cb",
		ChainId:        97,
		Name:           "xTIE",
		TokenBasicName: "xTIE",
		Property:       1,
		Precision:      18,
		ResourceId:     "0x000000000000000000000000000000c76ebe4a02bbc34786d860b355f5a5ce02",
		State:          true,
	}

	chainTie := models.BridgeChain{
		ChainId: 200,
		Name:    "TIE",
		Rpc:     "http://119.28.23.120:8545/",
		Topic:   "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
	}

	chainBnb := models.BridgeChain{
		ChainId: 97,
		Name:    "BNB",
		Rpc:     "https://data-seed-prebsc-1-s1.binance.org:8545/",
		Topic:   "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
	}

	o := orm.NewOrm()
	o.Insert(&tokenBasic)
	o.Insert(&tokenTie)
	o.Insert(&tokenBnb)

	o.Insert(&tokenxTieBasic)
	o.Insert(&tokenXTieTie)
	o.Insert(&tokenxTieBnb)

	o.Insert(&chainTie)
	o.Insert(&chainBnb)
}
