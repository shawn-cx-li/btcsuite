package omnilayer

import (
	"encoding/json"

	"github.com/shawn-cx-li/omnilayer-go/omnijson"
)

type futureCreateRawTransaction chan *response

func (f futureCreateRawTransaction) Receive() (omnijson.CreateRawTransactionResult, error) {
	var result omnijson.CreateRawTransactionResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureGetBlockChainInfo chan *response

func (f futureGetBlockChainInfo) Receive() (omnijson.GetBlockChainInfoResult, error) {
	var result omnijson.GetBlockChainInfoResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureListUnspent chan *response

func (f futureListUnspent) Receive() (omnijson.ListUnspentResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := make(omnijson.ListUnspentResult, 0)

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureImportAddress chan *response

func (f futureImportAddress) Receive() error {
	_, err := receive(f)
	return err
}

type futureSendRawTransaction chan *response

func (f futureSendRawTransaction) Receive() (omnijson.SendRawTransactionResult, error) {
	var result omnijson.SendRawTransactionResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureSignRawTransaction chan *response

func (f futureSignRawTransaction) Receive() (omnijson.SignRawTransactionResult, error) {
	var result omnijson.SignRawTransactionResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureSignRawTransactionWithKey chan *response

func (f futureSignRawTransactionWithKey) Receive() (omnijson.SignRawTransactionWithKeyResult, error) {
	var result omnijson.SignRawTransactionWithKeyResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniSend chan *response

func (f futureOmniSend) Receive() (omnijson.OmniSendResult, error) {
	var result omnijson.OmniSendResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniFundedSendall chan *response

func (f futureOmniFundedSendall) Receive() (omnijson.OmniFundedSendallResult, error) {
	var result omnijson.OmniFundedSendallResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}
