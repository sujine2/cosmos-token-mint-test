package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sujine/test/x/test/types"
)

// SetSymbolList set a specific symbolList in the store from its index
func (k Keeper) SetSymbolList(ctx sdk.Context, symbolList types.SymbolList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SymbolListKeyPrefix))
	b := k.cdc.MustMarshal(&symbolList)
	store.Set(types.SymbolListKey(
		symbolList.Index,
	), b)
}

// GetSymbolList returns a symbolList from its index
func (k Keeper) GetSymbolList(
	ctx sdk.Context,
	index string,

) (val types.SymbolList, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SymbolListKeyPrefix))

	b := store.Get(types.SymbolListKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSymbolList removes a symbolList from the store
func (k Keeper) RemoveSymbolList(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SymbolListKeyPrefix))
	store.Delete(types.SymbolListKey(
		index,
	))
}

// GetAllSymbolList returns all symbolList
func (k Keeper) GetAllSymbolList(ctx sdk.Context) (list []types.SymbolList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SymbolListKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SymbolList
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
