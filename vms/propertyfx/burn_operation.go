package propertyfx

import (
	"github.com/corpetty/avalanchego/vms/components/verify"
	"github.com/corpetty/avalanchego/vms/secp256k1fx"
)

// BurnOperation ...
type BurnOperation struct {
	secp256k1fx.Input `serialize:"true"`
}

// Outs ...
func (op *BurnOperation) Outs() []verify.State { return nil }
