package omnijson

type OmniGetWalletBalancesResult = []struct {
	PropertyID int64  `json:"propertyid"`
	Name       string `json:"name"`
	Balance    string `json:"balance"`
	Reserved   string `json:"reserved"`
	Frozen     string `json:"frozen"`
}

type OmniGetWalletBalancesCommand struct{}

func (OmniGetWalletBalancesCommand) Method() string {
	return "omni_getwalletbalances"
}

func (OmniGetWalletBalancesCommand) ID() string {
	return "1"
}

func (OmniGetWalletBalancesCommand) Params() []interface{} {
	return nil
}
