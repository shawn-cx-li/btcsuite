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

	err = json.Unmarshal(data, &result)
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

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureGetReceivedByAddress chan *response

func (f futureGetReceivedByAddress) Receive() (btcjson.GetreceivedbyaddressResult, error) {
	var result btcjson.GetreceivedbyaddressResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureValidateAddress chan *response

func (f futureValidateAddress) Receive() (btcjson.ValidateAddressResult, error) {
	var result btcjson.ValidateAddressResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureSetTxFee chan *response

func (f futureSetTxFee) Receive() error {
	_, err := receive(f)
	if err != nil {
		return err
	}
	return nil
}

type futureSendToAddress chan *response

func (f futureSendToAddress) Receive() (btcjson.SendToAddressResult, error) {
	var result btcjson.SendToAddressResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, nil
}

type futureEstimateSmartFee chan *response

func (f futureEstimateSmartFee) Receive() (btcjson.EstimateSmartFeeResult, error) {
	var result btcjson.EstimateSmartFeeResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureEstimateFee chan *response

func (f futureEstimateFee) Receive() (btcjson.EstimateFeeResult, error) {
	var result btcjson.EstimateFeeResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}
