package btcjson

type SendToAddressCommand struct {
	Address   string
	Amount    float64
	Comment   *string
	CommentTo *string
}

func (SendToAddressCommand) Method() string {
	return "sendtoaddress"
}

func (SendToAddressCommand) ID() string {
	return "1"
}

func (cmd SendToAddressCommand) Params() []interface{} {
	return []interface{}{cmd.Address, cmd.Amount, cmd.Comment, cmd.CommentTo}
}
