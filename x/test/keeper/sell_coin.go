package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sujine/test/x/test/types"
)

// SetSellCoin set a specific sellCoin in the store from its index
func (k Keeper) SetSellCoin(ctx sdk.Context, sellCoin types.SellCoin) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SellCoinKeyPrefix))
	b := k.cdc.MustMarshal(&sellCoin)
	store.Set(types.SellCoinKey(
		sellCoin.Index,
	), b)
}

// GetSellCoin returns a sellCoin from its index
func (k Keeper) GetSellCoin(
	ctx sdk.Context,
	index string,

) (val types.SellCoin, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SellCoinKeyPrefix))

	b := store.Get(types.SellCoinKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSellCoin removes a sellCoin from the store
func (k Keeper) RemoveSellCoin(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SellCoinKeyPrefix))
	store.Delete(types.SellCoinKey(
		index,
	))
}

// GetAllSellCoin returns all sellCoin
func (k Keeper) GetAllSellCoin(ctx sdk.Context) (list []types.SellCoin) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SellCoinKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SellCoin
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
