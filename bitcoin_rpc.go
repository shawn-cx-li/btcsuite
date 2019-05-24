package btcsuite

import (
	"github.com/shawn-cx-li/btcsuite/btcjson"
)

/*
 * In Use
 */
func (c *Client) GetNewAddress(account string) (btcjson.GetNewAddressResult, error) {
	return futureGetNewAddress(c.do(btcjson.GetNewAddressCommand{
		Account: account,
	})).Receive()
}

func (c *Client) GetTransaction(txid string) (btcjson.GetTransactionResult, error) {
	return futureGetTransaction(c.do(btcjson.GetTransactionCommand{
		TxID: txid,
	})).Receive()
}

// GetBalance returns the balance of a specific account
// Use "*" for all account, default minConf is 1
func (c *Client) GetBalance(account string, minConf int) (btcjson.GetBalanceResult, error) {
	return futureGetBalance(c.do(btcjson.GetBalanceCommand{
		Account: account,
		MinConf: minConf,
	})).Receive()
}

// GetReceivedByAddress gets the balance of a specific address, including watch-only
func (c *Client) GetReceivedByAddress(address string, minConf int) (btcjson.GetreceivedbyaddressResult, error) {
	return futureGetBalance(c.do(btcjson.GetreceivedbyaddressCommand{
		Address: address,
		MinConf: minConf,
	})).Receive()
}

func (c *Client) ValidateAddress(addr string) (btcjson.ValidateAddressResult, error) {
	return futureValidateAddress(c.do(btcjson.ValidateAddressCommand{
		Address: addr,
	})).Receive()
}

func (c *Client) SetTxFee(amt float64) error {
	return futureSetTxFee(c.do(btcjson.SetTxFeeCommand{
		Amount: amt,
	})).Receive()
}

func (c *Client) SendToAddress(toAddr string, amt float64, comment, commentTo string) (btcjson.SendToAddressResult, error) {
	return futureSendToAddress(c.do(btcjson.SendToAddressCommand{
		Address:   toAddr,
		Amount:    amt,
		Comment:   comment,
		CommentTo: commentTo,
	})).Receive()
}

func (c *Client) EstimateSmartFee(blocks uint32) (btcjson.EstimateSmartFeeResult, error) {
	return futureEstimateSmartFee(c.do(btcjson.EstimateSmartFeeCommand{
		Blocks: blocks,
	})).Receive()
}

func (c *Client) EstimateFee(blocks uint32) (btcjson.EstimateFeeResult, error) {
	return futureEstimateFee(c.do(btcjson.EstimateFeeCommand{
		Blocks: blocks,
	})).Receive()
}

/*
 * Not In Use
 */
func (c *Client) GetBlockChainInfo() (btcjson.GetBlockChainInfoResult, error) {
	return futureGetBlockChainInfo(c.do(btcjson.GetBlockChainInfoCommand{})).Receive()
}

func (c *Client) ListUnspent(min int, addresses []string) (btcjson.ListUnspentResult, error) {
	return futureListUnspent(c.do(btcjson.ListUnspentCommand{
		Min:       min,
		Addresses: addresses,
	})).Receive()
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
