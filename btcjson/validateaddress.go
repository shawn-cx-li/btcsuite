package btcjson

type ValidateAddressResult struct {
	IsValid      bool     `json:"isvalid"`
	Address      string   `json:"address,omitempty"`
	IsMine       bool     `json:"ismine,omitempty"`
	IsWatchOnly  bool     `json:"iswatchonly,omitempty"`
	IsScript     bool     `json:"isscript,omitempty"`
	PubKey       string   `json:"pubkey,omitempty"`
	IsCompressed bool     `json:"iscompressed,omitempty"`
	Account      string   `json:"account,omitempty"`
	Addresses    []string `json:"addresses,omitempty"`
	Hex          string   `json:"hex,omitempty"`
	Script       string   `json:"script,omitempty"`
	SigsRequired int32    `json:"sigsrequired,omitempty"`
}

type ValidateAddressCommand struct {
	Address string
}

func (ValidateAddressCommand) Method() string {
	return "validateaddress"
}

func (ValidateAddressCommand) ID() string {
	return "1"
}

func (cmd ValidateAddressCommand) Params() []interface{} {
	return []interface{}{cmd.Address}
}
