package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// Block represents a block in a blockchain.
// It contains a Hash, a slice of Transactions, the previous block's Hash, and a Nonce.
type Block struct {
	Hash         []byte         // Hash is the SHA-256 hash of the block
	Transactions []*Transaction // Transactions is a slice of all the transactions in the block
	PrevHash     []byte         // PrevHash is the hash of the previous block in the blockchain
	Nonce        int            // Nonce is the number used to mine the block
}

// HashTransactions returns the SHA-256 hash of all the transactions in the block
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

// CreateBlock creates a new block with the given transactions and previous block's hash.
// It also mines the block by calling the Run method of a new instance of the Proof of Work
func CreateBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{[]byte{}, txs, prevHash, 0}

	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Genesis creates the first block in the blockchain with the given coinbase transaction
func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}

// Serialize converts a block struct to a slice of bytes
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

// Deserialize converts a slice of bytes to a block struct
func Deserialize(data []byte) *Block {
	var Block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&Block)

	Handle(err)

	return &Block
}

// Handle is a helper function to handle errors
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
