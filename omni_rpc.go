package btcsuite

import (
	"gitlab.com/sdce/btcsuite/omnijson"
)

func (c *Client) OmniListBlockTransactions(block int64) (omnijson.OmniListBlockTransactionsResult, error) {
	return futureOmniListBlockTransactions(c.do(omnijson.OmniListBlockTransactionsCommand{
		Block: block,
	})).Receive()
}

func (c *Client) GetInfo() (omnijson.OmniGetInfoResult, error) {
	return futureGetInfo(c.do(omnijson.OmniGetInfoCommand{})).Receive()
}

func (c *Client) OmniGetTransaction(hash string) (omnijson.OmniGettransactionResult, error) {
	return futureOmniGetTransaction(c.do(omnijson.OmniGetTransactionCommand{
		Hash: hash,
	})).Receive()
}

func (c *Client) OmniCreatePayloadSimpleSend(cmd omnijson.OmniCreatePayloadSimpleSendCommand) (omnijson.OmniCreatePayloadSimpleSendResult, error) {
	return futureOmniCreatePayloadSimpleSend(c.do(cmd)).Receive()
}

func (c *Client) OmniCreateRawTxOpReturn(cmd omnijson.OmniCreateRawTxOpReturnCommand) (omnijson.OmniCreateRawTxOpReturnResult, error) {
	return futureOmniCreateRawTxOpReturn(c.do(cmd)).Receive()
}

func (c *Client) OmniCreateRawTxReference(cmd omnijson.OmniCreateRawTxReferenceCommand) (omnijson.OmniCreateRawTxReferenceResult, error) {
	return futureOmniCreateRawTxReference(c.do(cmd)).Receive()
}

func (c *Client) OmniCreateRawTxChange(cmd omnijson.OmniCreateRawTxChangeCommand) (omnijson.OmniCreateRawTxChangeResult, error) {
	return futureOmniCreateRawTxChange(c.do(cmd)).Receive()
}

func (c *Client) OmniGetBalance(cmd omnijson.OmniGetBalanceCommand) (omnijson.OmniGetBalanceResult, error) {
	return futureOmniGetBalance(c.do(cmd)).Receive()
}

func (c *Client) OmniGetAllBalancesForID(propertyID int32) (omnijson.OmniGetAllBalancesForIDResult, error) {
	return futureOmniGetAllBalancesForID(c.do(omnijson.OmniGetAllBalancesForIDCommand{PropertyID: propertyID})).Receive()
}

func (c *Client) OmniSend(cmd omnijson.OmniSendCommand) (omnijson.OmniSendResult, error) {
	return futureOmniSend(c.do(cmd)).Receive()
}

func (c *Client) OmniFundedSendall(cmd omnijson.OmniFundedSendallCommand) (omnijson.OmniFundedSendallResult, error) {
	return futureOmniFundedSendall(c.do(cmd)).Receive()
}
