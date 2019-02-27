package btcjson

type GetBalanceResult struct {
	Balance float64
}

type GetBalanceCommand struct {
}

func (GetBalanceCommand) Method() string {
	return "getbalance"
}

func (GetBalanceCommand) ID() string {
	return "1"
}

func (GetBalanceCommand) Params() []interface{} {
	return nil
}
