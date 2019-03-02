package omnijson

type OmniSendResult = string

type OmniSendCommand struct {
	FromAddress string
	ToAddress   string
	PropertyID  int32
	Amount      float64
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
