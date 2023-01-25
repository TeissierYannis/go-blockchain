package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Difficulty TODO Dynamic difficulty (to implement)
// Difficulty is a constant value that represents the difficulty level of the proof of work algorithm.
// Difficulty The higher the difficulty, the more difficult it is to find a solution, and the more secure the blockchain is.
var Difficulty = 12

// ProofOfWork is a struct that contains a Block and a Target value. The target value is used in the proof of work algorithm to check if a solution is valid.
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProof creates a new ProofOfWork struct with the block and target value.
func NewProof(b *Block) *ProofOfWork {
	// Create a big.Int with the value of 1
	target := big.NewInt(1)
	// Left shift the bits of the big.Int by the number of bits necessary to reach the desired difficulty
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{b, target}

	// Run the proof of work algorithm
	_, _ = pow.Run()

	return pow
}

// InitData creates a byte slice with the previous block hash, transactions hash, nonce and difficulty
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.HashTransactions(),
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)

	return data
}

// Run performs the proof of work algorithm, repeatedly hashing the data with a nonce until the solution is found
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

// Validate performs the proof of work algorithm on the block's nonce to check if it is a valid solution
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

// ToHex converts an int64 number to a byte slice
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)

	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
