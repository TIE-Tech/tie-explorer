package task

import (
	"encoding/hex"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strings"
	"tie-explorer-go/models"
	"tie-explorer-go/rpc"
	"time"
)

var (
	refreshAddress bool

	AddressChan = make(chan common.Address)
)

var (
	ErrHasAddressRefresh = errors.New("the account balance is being processed")
)

func refreshAllAddressBalance() error {
	if refreshAddress {
		logs.Warn("%s", ErrHasAddressRefresh.Error())
		return ErrHasAddressRefresh
	}

	logs.Info("refresh account balance")
	refreshAddress = true

	c := rpc.GetClient()
	o := orm.NewOrm()

	var addresses []models.Address
	o.QueryTable("address").All(&addresses)

	for _, address := range addresses {
		balance := c.GetAccountBalance(common.HexToAddress(address.Address))

		addressModel := &models.Address{
			Balance: balance.String(),
		}

		o.Update(addressModel, "address")

		dailyModel := models.AddressBalanceDaily{
			Address: address.Address,
			Date:    time.Now().Format("2006-01-02"),
			Value:   balance.String(),
		}
		o.InsertOrUpdate(&dailyModel, "address,date")

		GetTokenBalance(address.Address, "")
	}

	refreshAddress = false
	return nil
}

func getContractCode(address common.Address) []byte {
	c := rpc.GetClient()
	code, err := c.GetCode(address, nil)
	if err != nil {
		logs.Error("get contract code error, %s", err.Error())
		return nil
	}
	// TODO
	// https://github.com/ethereum/evmdasm
	return code
}

func getCallMsg(address, contractAddress string) ethereum.CallMsg {
	to := common.HexToAddress(contractAddress)
	b, _ := hexutil.Decode("0x70a08231000000000000000000000000" + address[2:])

	return ethereum.CallMsg{
		To:   &to,
		Data: b,
	}
}

func GetTokenBalance(address, contractAddress string) {
	c := rpc.GetClient()
	o := orm.NewOrm()

	if contractAddress == "" {
		var tokenList []models.Token
		o.QueryTable("token").All(&tokenList)
		for _, item := range tokenList {
			callMsg := getCallMsg(address, item.ContractAddress)
			result, err := c.CallContractAtBlockNumber(callMsg, nil)
			if err != nil {
				continue
			}
			value, _ := hexutil.DecodeBig("0x" + strings.TrimLeft(hex.EncodeToString(result), "0"))
			var tokenModel models.AddressTokenBalance
			err = o.QueryTable("address_token_balance").Filter("address__iexact", address).
				Filter("contract_address__iexact", contractAddress).One(&tokenModel)
			if err != nil {
				return
			}
			tokenModel.Value = value.String()
			o.InsertOrUpdate(tokenModel, "address,contract_address")
		}
	} else {
		callMsg := getCallMsg(address, contractAddress)
		result, err := c.CallContractAtBlockNumber(callMsg, nil)
		if err != nil {
			return
		}
		value, _ := hexutil.DecodeBig("0x" + strings.TrimLeft(hex.EncodeToString(result), "0"))
		var tokenModel models.AddressTokenBalance
		err = o.QueryTable("address_token_balance").Filter("address__iexact", address).
			Filter("contract_address__iexact", contractAddress).One(&tokenModel)
		if err != nil {
			return
		}
		tokenModel.Value = value.String()
		o.InsertOrUpdate(&tokenModel, "address,contract_address")
	}
}

func parseAddressInfo(address common.Address) {
	o := orm.NewOrm()
	c := rpc.GetClient()

	addressModel := &models.Address{}

	err := o.QueryTable("address").Filter("address", address.Hex()).One(addressModel)
	if err != nil {
		addressModel.Address = address.Hex()
		codeBytes := getContractCode(address)
		if len(codeBytes) > 0 {
			addressModel.IsContract = true
			addressModel.ContractCode = "0x" + hex.EncodeToString(codeBytes)
		} else {
			addressModel.IsContract = false
		}

		balance := c.GetAccountBalance(common.HexToAddress(address.Hex()))
		addressModel.Balance = balance.String()
		o.Insert(addressModel)

		dailyModel := models.AddressBalanceDaily{
			Address: address.Hex(),
			Date:    time.Now().Format("2006-01-02"),
			Value:   balance.String(),
		}
		o.InsertOrUpdate(&dailyModel, "address,date")
	}
}

func NewAddressTask() {
	go func() {
		for {
			select {
			case address := <-AddressChan:
				parseAddressInfo(address)
			}
		}
	}()

	addressTk := toolbox.NewTask("address_tk", "* * 1 * * *", refreshAllAddressBalance)
	toolbox.AddTask("address_fetcher_task", addressTk)
}
