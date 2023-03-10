package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/TessorNetwork/furya/config"

	"github.com/TessorNetwork/furya/x/token/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey sdk.StoreKey
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) IncTokens(ctx sdk.Context, address sdk.AccAddress, amount sdk.Fur) {
	if amount.IsZero() {
		return
	}

	k.SetBalance(ctx, address, k.GetBalance(ctx, address).Add(amount))
}

func (k Keeper) GetBalance(ctx sdk.Context, address sdk.AccAddress) sdk.Fur {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BalancePrefix)

	if !store.Has(address) {
		return config.InitialTokenBalance
	}

	var d sdk.Fur
	if err := d.Unmarshal(store.Get(address)); err != nil {
		panic(fmt.Errorf("failed to get balance for %s: %w", address, err))
	}
	return d
}

func (k Keeper) SetBalance(ctx sdk.Context, address sdk.AccAddress, amount sdk.Fur) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BalancePrefix)

	if amount.IsZero() {
		store.Delete(address)
		return
	}

	bz, err := amount.Marshal()
	if err != nil {
		panic(fmt.Errorf("failed to marshal new balance for %s: %w", address, err))
	}
	store.Set(address, bz)
}

func (k Keeper) ResetAccount(ctx sdk.Context, address sdk.AccAddress) {
	k.SetBalance(ctx, address, sdk.ZeroFur())
}

func (k Keeper) IterateBalance(ctx sdk.Context, handle func(address sdk.AccAddress, balance sdk.Fur) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BalancePrefix)

	it := store.Iterator(nil, nil)
	defer it.Close()

	for it.Valid() {
		address := sdk.AccAddress(it.Key())

		var delta sdk.Fur
		if err := delta.Unmarshal(it.Value()); err != nil {
			panic(fmt.Errorf("failed to get balance for %s: %w", address, err))
		}

		if handle(it.Key(), delta) {
			return
		}

		it.Next()
	}
}
