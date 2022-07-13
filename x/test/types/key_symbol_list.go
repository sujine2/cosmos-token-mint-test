package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SymbolListKeyPrefix is the prefix to retrieve all SymbolList
	SymbolListKeyPrefix = "SymbolList/value/"
)

// SymbolListKey returns the store key to retrieve a SymbolList from the index fields
func SymbolListKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
