package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

// Transaction struct represents a transaction in the blockchain. It contains the transaction ID, inputs, and outputs.

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

// SetID sets the transaction ID by encoding the transaction data and hashing it.
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// CoinbaseTx creates a new coinbase transaction with the given parameters.
// CoinbaseTx The function takes the recipient address and an optional data string.
// CoinbaseTx If data is not provided, it will be set to "Coins to [recipient address]".
func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{100, to}

	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
	tx.SetID()

	return &tx
}

// NewTransaction creates a new transaction with the given parameters.
// NewTransaction The function takes the sender and recipient addresses, amount, and a reference to the blockchain.
// NewTransaction The function uses the FindSpendableOutputs method to find spendable outputs and create the inputs for the new transaction.
// NewTransaction If the sender does not have enough funds, the function will panic.
func NewTransaction(from, to string, amount int, chain *Blockchain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutputs := chain.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("Error: not enough funds")
	}

	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		Handle(err)

		for _, out := range outs {
			input := TxInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxOutput{amount, to})

	if acc > amount {
		outputs = append(outputs, TxOutput{acc - amount, from})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}

// IsCoinbase checks if a transaction is a coinbase transaction
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}
