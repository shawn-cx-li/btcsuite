package btcsuite

import (
	"gitlab.com/sdce/btcsuite/btcjson"
)

func (c *Client) GetBlockChainInfo() (btcjson.GetBlockChainInfoResult, error) {
	return futureGetBlockChainInfo(c.do(btcjson.GetBlockChainInfoCommand{})).Receive()
}

func (c *Client) ListUnspent(cmd btcjson.ListUnspentCommand) (btcjson.ListUnspentResult, error) {
	return futureListUnspent(c.do(cmd)).Receive()
}

func (c *Client) CreateRawTransaction(cmd btcjson.CreateRawTransactionCommand) (btcjson.CreateRawTransactionResult, error) {
	return futureCreateRawTransaction(c.do(cmd)).Receive()
}

func (c *Client) ImportAddress(address string, rescan bool) error {
	return futureImportAddress(c.do(btcjson.ImportAddressCommand{Adress: address, Rescan: rescan})).Receive()
}

func (c *Client) SendRawTransaction(cmd btcjson.SendRawTransactionCommand) (btcjson.SendRawTransactionResult, error) {
	return futureSendRawTransaction(c.do(cmd)).Receive()
}

func (c *Client) SignRawTransaction(cmd btcjson.SignRawTransactionCommand) (btcjson.SignRawTransactionResult, error) {
	return futureSignRawTransaction(c.do(cmd)).Receive()
}

func (c *Client) SignRawTransactionWithKey(cmd btcjson.SignRawTransactionWithKeyCommand) (btcjson.SignRawTransactionWithKeyResult, error) {
	return futureSignRawTransactionWithKey(c.do(cmd)).Receive()
}

func (c *Client) GetNewAddress() (string, error) {
	return futureGetNewAddress(c.do(btcjson.GetNewAddressCommand{})).Receive()
}

func (c *Client) GetTransaction(cmd btcjson.GetTransactionCommand) (btcjson.GetTransactionResult, error) {
	return futureGetTransaction(c.do(cmd)).Receive()
}

func (c *Client) GetBalance() (float64, error) {
	return futureGetBalance(c.do(btcjson.GetBalanceCommand{})).Receive()
}

func (c *Client) ValidateAddress(cmd btcjson.ValidateAddressCommand) (btcjson.ValidateAddressResult, error) {
	return futureValidateAddress(c.do(cmd)).Receive()
}

func (c *Client) SetTxFee(cmd btcjson.SetTxFeeCommand) error {
	return futureSetTxFee(c.do(cmd)).Receive()
}

func (c *Client) SendToAddress(cmd btcjson.SendToAddressCommand) (string, error) {
	return futureSendToAddress(c.do(cmd)).Receive()
}
