package btcsuite

import (
	"encoding/json"

	"gitlab.com/sdce/btcsuite/btcjson"
)

type futureCreateRawTransaction chan *response

func (f futureCreateRawTransaction) Receive() (btcjson.CreateRawTransactionResult, error) {
	var result btcjson.CreateRawTransactionResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureGetBlockChainInfo chan *response

func (f futureGetBlockChainInfo) Receive() (btcjson.GetBlockChainInfoResult, error) {
	var result btcjson.GetBlockChainInfoResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureListUnspent chan *response

func (f futureListUnspent) Receive() (btcjson.ListUnspentResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := make(btcjson.ListUnspentResult, 0)

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureImportAddress chan *response

func (f futureImportAddress) Receive() error {
	_, err := receive(f)
	return err
}

type futureSendRawTransaction chan *response

func (f futureSendRawTransaction) Receive() (btcjson.SendRawTransactionResult, error) {
	var result btcjson.SendRawTransactionResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureSignRawTransaction chan *response

func (f futureSignRawTransaction) Receive() (btcjson.SignRawTransactionResult, error) {
	var result btcjson.SignRawTransactionResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureSignRawTransactionWithKey chan *response

func (f futureSignRawTransactionWithKey) Receive() (btcjson.SignRawTransactionWithKeyResult, error) {
	var result btcjson.SignRawTransactionWithKeyResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureGetNewAddress chan *response

func (f futureGetNewAddress) Receive() (btcjson.GetNewAddressResult, error) {
	var result btcjson.GetNewAddressResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	var address string
	err = json.Unmarshal(data, &address)

	result = btcjson.GetNewAddressResult{
		Address: address,
	}
	return result, err
}

type futureGetTransaction chan *response

func (f futureGetTransaction) Receive() (btcjson.GetTransactionResult, error) {
	var result btcjson.GetTransactionResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureGetBalance chan *response

func (f futureGetBalance) Receive() (btcjson.GetBalanceResult, error) {
	var result btcjson.GetBalanceResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	var balance int64
	err = json.Unmarshal(data, &balance)

	result = btcjson.GetBalanceResult{
		Balance: balance,
	}
	return result, err
}
