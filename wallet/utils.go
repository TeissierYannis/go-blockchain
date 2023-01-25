package wallet

import (
	"github.com/mr-tron/base58"
	"log"
)

// Base58Encode encodes the input byte slice using the base58 encoding scheme and returns the encoded byte slice
func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

// Base58Decode decodes the input byte slice using the base58 encoding scheme and returns the decoded byte slice
// Base58Decode It will panic if it fails to decode the input byte slice
func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}

	return decode
}
