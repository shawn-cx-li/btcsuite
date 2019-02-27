package btcjson

type GetNewAddressCommand struct{}

func (GetNewAddressCommand) Method() string {
	return "getnewaddress"
}

func (GetNewAddressCommand) ID() string {
	return "1"
}

func (GetNewAddressCommand) Params() []interface{} {
	return nil
}
