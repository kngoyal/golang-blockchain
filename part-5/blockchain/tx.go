package blockchain

type TxOutput struct {
	// Indivisible Outputs
	Value  int    // value token
	PubKey string // to unlock the tokens; derived from Script language
}

type TxInput struct {
	// references to previous outputs
	ID  []byte // reference to the txn that the output is inside
	Out int    // index where the output appears
	Sig string
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
