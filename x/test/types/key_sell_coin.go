package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SellCoinKeyPrefix is the prefix to retrieve all SellCoin
	SellCoinKeyPrefix = "SellCoin/value/"
)

// SellCoinKey returns the store key to retrieve a SellCoin from the index fields
func SellCoinKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
