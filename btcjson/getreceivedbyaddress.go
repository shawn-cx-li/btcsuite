package btcjson

type GetreceivedbyaddressCommand struct {
	Address string
	MinConf int `jsonrpcdefault:"1"`
}

type GetreceivedbyaddressResult = float64

func (GetreceivedbyaddressCommand) Method() string {
	return "getreceivedbyaddress"
}

func (GetreceivedbyaddressCommand) ID() string {
	return "1"
}

func (cmd GetreceivedbyaddressCommand) Params() []interface{} {
	return []interface{}{cmd.Address, cmd.MinConf}
}
