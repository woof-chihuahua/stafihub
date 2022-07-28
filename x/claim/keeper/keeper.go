package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stafihub/stafihub/x/claim/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		sudoKeeper types.SudoKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	sudoKeeper types.SudoKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		sudoKeeper: sudoKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetMerkleRoot(ctx sdk.Context, round uint64, root [32]byte) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.MerkleRootStoreKey(round), root[:])
}

func (k Keeper) GetMerkleRoot(ctx sdk.Context, round uint64) ([32]byte, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.MerkleRootStoreKey(round))
	if bts == nil {
		return [32]byte{}, false
	}

	var root [32]byte
	copy(root[:], bts)

	return root, true
}

func (k Keeper) SetClaimRound(ctx sdk.Context, round uint64) {
	store := ctx.KVStore(k.storeKey)

	bts := sdk.Uint64ToBigEndian(round)

	store.Set(types.ClaimRoundStoreKey, bts)
}

func (k Keeper) GetClaimRound(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ClaimRoundStoreKey)
	if bts == nil {
		return 0
	}

	return sdk.BigEndianToUint64(bts)
}

func (k Keeper) setClaimBitMap(ctx sdk.Context, claimRound, wordIndex, bitIndex uint64) {
	store := ctx.KVStore(k.storeKey)

	bts := sdk.Uint64ToBigEndian(bitIndex)

	store.Set(types.ClaimBitMapStoreKey(claimRound, wordIndex), bts)
}

func (k Keeper) getClaimBitMap(ctx sdk.Context, claimRound, wordIndex uint64) uint64 {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ClaimBitMapStoreKey(claimRound, wordIndex))
	if bts == nil {
		return 0
	}

	return sdk.BigEndianToUint64(bts)
}

func (k Keeper) IsClaimed(ctx sdk.Context, claimRound, index uint64) bool {
	claimedWordIndex := index / 64
	claimedBitIndex := index % 64

	mask := uint64(1 << claimedBitIndex)
	existBitIndex := k.getClaimBitMap(ctx, claimRound, claimedWordIndex)

	return (existBitIndex & mask) == mask
}

func (k Keeper) SetClaimed(ctx sdk.Context, claimRound, index uint64) {
	claimedWordIndex := index / 64
	claimedBitIndex := index % 64

	existBitIndex := k.getClaimBitMap(ctx, claimRound, claimedWordIndex)
	newBitIndex := existBitIndex | (1 << claimedBitIndex)

	k.setClaimBitMap(ctx, claimRound, claimedWordIndex, newBitIndex)
}
