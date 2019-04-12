package btcjson

type GetBalanceCommand struct {
	Account string
	MinConf int `jsonrpcdefault:"1"`
}

type GetBalanceResult = float64

func (GetBalanceCommand) Method() string {
	return "getbalance"
}

func (GetBalanceCommand) ID() string {
	return "1"
}

func (cmd GetBalanceCommand) Params() []interface{} {
	return []interface{}{cmd.Account, cmd.MinConf}
}
