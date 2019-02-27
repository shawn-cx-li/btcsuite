package btcjson

type SetTxFeeCommand struct {
	Amount float64 // In BTC
}

func (SetTxFeeCommand) Method() string {
	return "settxfee"
}

func (SetTxFeeCommand) ID() string {
	return "1"
}

func (SetTxFeeCommand) Params() []interface{} {
	return nil
}
