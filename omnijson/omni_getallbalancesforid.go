package omnijson

type OmniGetAllBalancesForIDResult = struct {
	Balances []OmniGetAllBalancesForIDBody
}

type OmniGetAllBalancesForIDBody = struct {
	Address  string
	Balance  float64
	Reserved float64
	Frozen   float64
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
