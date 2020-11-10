package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ProviderID struct {
	Owner     sdk.Address
	Validator sdk.Address
}
