package btcjson

type EstimateSmartFeeCommand struct {
	Blocks uint32
}

type EstimateSmartFeeResult struct {
	FeeRate float64 `json:"feerate"`
	Blocks  uint32  `json:"blocks"`
}

func (EstimateSmartFeeCommand) Method() string {
	return "estimatesmartfee"
}

func (EstimateSmartFeeCommand) ID() string {
	return "1"
}

func (cmd EstimateSmartFeeCommand) Params() []interface{} {
	return []interface{}{cmd.Blocks}
}
