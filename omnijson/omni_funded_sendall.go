package omnijson

type OmniFundedSendallResult struct {
	Txid string
}

type OmniFundedSendallCommand struct {
	FromAddress string
	ToAddress   string
	Ecosystem   int32
	FeeAddress  string
}

func (OmniFundedSendallCommand) Method() string {
	return "omni_funded_sendall"
}

func (OmniFundedSendallCommand) ID() string {
	return "1"
}

func (cmd OmniFundedSendallCommand) Params() []interface{} {
	return []interface{}{cmd.FromAddress, cmd.ToAddress, cmd.Ecosystem, cmd.FeeAddress}
}
