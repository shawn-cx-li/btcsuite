package btcsuite

import (
	"github.com/shawn-cx-li/btcsuite/omnijson"
)

/*
 * In Use
 */
func (c *Client) OmniGetTransaction(hash string) (omnijson.OmniGettransactionResult, error) {
	return futureOmniGetTransaction(c.do(omnijson.OmniGetTransactionCommand{
		Hash: hash,
	})).Receive()
}

func (c *Client) OmniGetBalance(addr string, propertyID int64) (omnijson.OmniGetBalanceResult, error) {
	return futureOmniGetBalance(c.do(omnijson.OmniGetBalanceCommand{
		Address:    addr,
		PropertyID: propertyID,
	})).Receive()
}

func (c *Client) OmniGetAllBalancesForID(propertyID int64) (omnijson.OmniGetAllBalancesForIDResult, error) {
	return futureOmniGetAllBalancesForID(c.do(omnijson.OmniGetAllBalancesForIDCommand{
		PropertyID: propertyID,
	})).Receive()
}

func (c *Client) OmniGetWalletBalances() (omnijson.OmniGetWalletBalancesResult, error) {
	return futrueOmniGetWalletBalances(c.do(omnijson.OmniGetWalletBalancesCommand{})).Receive()
}

func (c *Client) OmniGetWalletAddressBalances() (omnijson.OmniGetWalletAddressBalancesResult, error) {
	return futrueOmniGetWalletAddressBalances(c.do(omnijson.OmniGetWalletAddressBalancesCommand{})).Receive()
}

func (c *Client) OmniSend(fromAddr, toAddr, amount string, propertyId int64) (omnijson.OmniSendResult, error) {
	return futureOmniSend(c.do(omnijson.OmniSendCommand{
		FromAddress: fromAddr,
		ToAddress:   toAddr,
		PropertyID:  propertyId,
		Amount:      amount,
	})).Receive()
}

func (c *Client) OmniFundedSendall(fromAddr, toAddr, feeAddr string, ecosystem int64) (omnijson.OmniFundedSendallResult, error) {
	return futureOmniFundedSendall(c.do(omnijson.OmniFundedSendallCommand{
		FromAddress: fromAddr,
		ToAddress:   toAddr,
		FeeAddress:  feeAddr,
		Ecosystem:   ecosystem,
	})).Receive()
}

/*
 * Not In Use
 */
func (c *Client) OmniListBlockTransactions(block int64) (omnijson.OmniListBlockTransactionsResult, error) {
	return futureOmniListBlockTransactions(c.do(omnijson.OmniListBlockTransactionsCommand{
		Block: block,
	})).Receive()
}

func (c *Client) GetInfo() (omnijson.OmniGetInfoResult, error) {
	return futureGetInfo(c.do(omnijson.OmniGetInfoCommand{})).Receive()
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
