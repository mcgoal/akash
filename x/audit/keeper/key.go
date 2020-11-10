package keeper

import (
	"bytes"

	"github.com/ovrclk/akash/x/audit/types"
)

var (
	prefixProviderID = []byte{0x01}
)

func providerKey(id types.ProviderID) []byte {
	buf := bytes.NewBuffer(prefixProviderID)
	if _, err := buf.Write(id.Validator.Bytes()); err != nil {
		panic(err)
	}

	if _, err := buf.Write(id.Owner.Bytes()); err != nil {
		panic(err)
	}

	return buf.Bytes()
}
