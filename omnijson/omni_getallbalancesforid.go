package omnijson

type OmniGetAllBalancesForIDResult = []struct {
	Address  string  `json:"address"`
	Balance  float64 `json:"balance"`
	Reserved float64 `json:"reserved"`
	Frozen   float64 `json:"frozen"`
}

type OmniGetAllBalancesForIDCommand struct {
	PropertyID int64
}

func (OmniGetAllBalancesForIDCommand) Method() string {
	return "omni_getallbalancesforid"
}

func (OmniGetAllBalancesForIDCommand) ID() string {
	return "1"
}

func (cmd OmniGetAllBalancesForIDCommand) Params() []interface{} {
	return []interface{}{cmd.PropertyID}
}
