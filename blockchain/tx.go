package blockchain

// TxOutput represents a transaction output, which contains a value and a public key.
type TxOutput struct {
	Value  int
	PubKey string
}

// TxInput represents a transaction input, which contains the ID of the previous transaction, the index of the output it's referencing, and a signature.
type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

// CanUnlock is a method on TxInput that checks if the input can be unlocked using the provided data, which is typically the public key of the recipient.
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

// CanBeUnlocked is a method on TxOutput that checks if the output can be unlocked using the provided data, which is typically the public key of the recipient.
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
