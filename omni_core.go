package btcsuite

import (
	"encoding/json"

	"github.com/shawn-cx-li/btcsuite/omnijson"
)

type futureOmniCreatePayloadSimpleSend chan *response

func (f futureOmniCreatePayloadSimpleSend) Receive() (omnijson.OmniCreatePayloadSimpleSendResult, error) {
	var result omnijson.OmniCreatePayloadSimpleSendResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniCreateRawTxChange chan *response

func (f futureOmniCreateRawTxChange) Receive() (omnijson.OmniCreateRawTxChangeResult, error) {
	var result omnijson.OmniCreateRawTxChangeResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniCreateRawTxOpReturn chan *response

func (f futureOmniCreateRawTxOpReturn) Receive() (omnijson.OmniCreateRawTxOpReturnResult, error) {
	var result omnijson.OmniCreateRawTxOpReturnResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniCreateRawTxReference chan *response

func (f futureOmniCreateRawTxReference) Receive() (omnijson.OmniCreateRawTxReferenceResult, error) {
	var result omnijson.OmniCreateRawTxReferenceResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureGetInfo chan *response

func (f futureGetInfo) Receive() (omnijson.OmniGetInfoResult, error) {
	var result omnijson.OmniGetInfoResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniGetTransaction chan *response

func (f futureOmniGetTransaction) Receive() (omnijson.OmniGettransactionResult, error) {
	var result omnijson.OmniGettransactionResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniListBlockTransactions chan *response

func (f futureOmniListBlockTransactions) Receive() (omnijson.OmniListBlockTransactionsResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := make(omnijson.OmniListBlockTransactionsResult, 0)

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniGetBalance chan *response

func (f futureOmniGetBalance) Receive() (omnijson.OmniGetBalanceResult, error) {
	var result omnijson.OmniGetBalanceResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniGetAllBalancesForID chan *response

func (f futureOmniGetAllBalancesForID) Receive() (omnijson.OmniGetAllBalancesForIDResult, error) {
	var result omnijson.OmniGetAllBalancesForIDResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futrueOmniGetWalletBalances chan *response

func (f futrueOmniGetWalletBalances) Receive() (omnijson.OmniGetWalletBalancesResult, error) {
	var result omnijson.OmniGetWalletBalancesResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futrueOmniGetWalletAddressBalances chan *response

func (f futrueOmniGetWalletAddressBalances) Receive() (omnijson.OmniGetWalletAddressBalancesResult, error) {
	var result omnijson.OmniGetWalletAddressBalancesResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniSend chan *response

func (f futureOmniSend) Receive() (omnijson.OmniFundedSendallResult, error) {
	var result omnijson.OmniFundedSendallResult

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
