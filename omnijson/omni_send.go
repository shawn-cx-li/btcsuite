package omnijson

type OmniSendResult struct {
	Txid string
}

type OmniSendCommand struct {
	FromAddress string
	ToAddress   string
	PropertyID  int32
	Amount      string
}

func (OmniSendCommand) Method() string {
	return "omni_send"
}

func (OmniSendCommand) ID() string {
	return "1"
}

func (cmd OmniSendCommand) Params() []interface{} {
	return []interface{}{cmd.FromAddress, cmd.ToAddress, cmd.PropertyID, cmd.Amount}
}
