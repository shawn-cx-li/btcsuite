package btcjson

type EstimateFeeCommand struct {
	Blocks uint32
}

type EstimateFeeResult = float64

func (EstimateFeeCommand) Method() string {
	return "estimatefee"
}

func (EstimateFeeCommand) ID() string {
	return "1"
}

func (cmd EstimateFeeCommand) Params() []interface{} {
	return []interface{}{cmd.Blocks}
}
