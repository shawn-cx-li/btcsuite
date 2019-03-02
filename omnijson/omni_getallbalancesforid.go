package omnijson

type OmniGetAllBalancesForIdResult = struct {
	Balances []OmniGetAllBalancesForIdBody
}

type OmniGetAllBalancesForIdBody = struct {
	Address  string
	Balance  string
	Reserved string
	Frozen   string
}

type OmniGetAllBalancesForIdCommand struct {
	PropertyID int32
}

func (OmniGetAllBalancesForIdCommand) Method() string {
	return "omni_getallbalancesforid"
}

func (OmniGetAllBalancesForIdCommand) ID() string {
	return "1"
}

func (cmd OmniGetAllBalancesForIdCommand) Params() []interface{} {
	return []interface{}{cmd.PropertyID}
}
