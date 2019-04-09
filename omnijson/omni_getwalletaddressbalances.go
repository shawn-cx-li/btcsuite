package omnijson

type OmniGetWalletAddressBalancesResult = []struct {
	Address  string               `json:"address"`
	Balances []OmniAddressBalance `json:"balances"`
}

type OmniGetWalletAddressBalancesCommand struct{}

func (OmniGetWalletAddressBalancesCommand) Method() string {
	return "omni_getwalletaddressbalances"
}

func (OmniGetWalletAddressBalancesCommand) ID() string {
	return "1"
}

func (OmniGetWalletAddressBalancesCommand) Params() []interface{} {
	return nil
}
