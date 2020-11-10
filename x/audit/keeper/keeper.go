package keeper

import (
	"fmt"
	"sort"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	atypes "github.com/ovrclk/akash/types"
	"github.com/ovrclk/akash/x/audit/types"
	ptypes "github.com/ovrclk/akash/x/provider/types"
)

// Keeper of the provider store
type Keeper struct {
	skey sdk.StoreKey
	cdc  codec.BinaryMarshaler
}

// NewKeeper creates and returns an instance for Market keeper
func NewKeeper(cdc codec.BinaryMarshaler, skey sdk.StoreKey) Keeper {
	return Keeper{cdc: cdc, skey: skey}
}

// Codec returns keeper codec
func (k Keeper) Codec() codec.BinaryMarshaler {
	return k.cdc
}

// GetProvider returns a provider with given validator and owner id's
func (k Keeper) GetProviderAttributes(ctx sdk.Context, id types.ProviderID) (types.Provider, bool) {
	store := ctx.KVStore(k.skey)

	key := providerKey(id)
	buf := store.Get(key)
	if buf == nil {
		return types.Provider{}, false
	}

	var val types.Provider
	k.cdc.MustUnmarshalBinaryBare(buf, &val)

	return val, true
}

// CreateOrUpdateProviderAttributes update signed provider attributes.
// creates new if key does not exist
// if key exists, existing values for matching pairs will be replaced
func (k Keeper) CreateOrUpdateProviderAttributes(ctx sdk.Context, id types.ProviderID, attr ptypes.Attributes) error {
	store := ctx.KVStore(k.skey)
	key := providerKey(id)

	prov := types.Provider{
		Owner:      id.Owner.String(),
		Validator:  id.Validator.String(),
		Attributes: attr,
	}

	buf := store.Get(key)
	if buf != nil {
		tmp := types.Provider{}
		k.cdc.MustUnmarshalBinaryBare(buf, &tmp)

		kv := make(map[string]string)

		for _, entry := range tmp.Attributes {
			kv[entry.Key] = entry.Value
		}

		for _, entry := range prov.Attributes {
			kv[entry.Key] = entry.Value
		}

		attr = ptypes.Attributes{}

		for ky, val := range kv {
			attr = append(attr, atypes.Attribute{
				Key:   ky,
				Value: val,
			})
		}

		sort.SliceStable(attr, func(i, j int) bool {
			return attr[i].Key < attr[j].Key
		})

		prov.Attributes = attr
	}

	store.Set(key, k.cdc.MustMarshalBinaryBare(&prov))
	return nil
}

func (k Keeper) DeleteProviderAttributes(ctx sdk.Context, id types.ProviderID, attr ptypes.Attributes) error {
	store := ctx.KVStore(k.skey)
	key := providerKey(id)

	buf := store.Get(key)
	if buf == nil {
		return types.ErrProviderNotFound
	}

	if attr == nil {
		store.Delete(key)
	} else {
		prov := types.Provider{
			Owner:      id.Owner.String(),
			Validator:  id.Validator.String(),
			Attributes: attr,
		}

		tmp := types.Provider{}
		k.cdc.MustUnmarshalBinaryBare(buf, &tmp)

		kv := make(map[string]string)

		for _, entry := range tmp.Attributes {
			kv[entry.Key] = entry.Value
		}

		for _, entry := range attr {
			if _, exists := kv[entry.Key]; !exists {
				return fmt.Errorf("trying to delete non-existsting attribute \"%s\" for validator/provider \"%s/%s\"",
					entry.Key,
					prov.Validator,
					prov.Owner)
			} else {
				delete(kv, entry.Key)
			}
		}

		attr = ptypes.Attributes{}

		for ky, val := range kv {
			attr = append(attr, atypes.Attribute{
				Key:   ky,
				Value: val,
			})
		}

		if len(attr) == 0 {
			store.Delete(key)
		} else {
			sort.SliceStable(attr, func(i, j int) bool {
				return attr[i].Key < attr[j].Key
			})

			prov.Attributes = attr

			store.Set(key, k.cdc.MustMarshalBinaryBare(&prov))
		}
	}
	return nil
}
