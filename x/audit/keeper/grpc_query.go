package keeper

import (
	"context"

	"github.com/ovrclk/akash/x/audit/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (k Querier) Auditors(context.Context, *types.QueryAuditorsRequest) (*types.QueryAuditorsResponse, error) {
	panic("not implemented")
}
