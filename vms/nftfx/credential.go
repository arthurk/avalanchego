package nftfx

import (
	"github.com/corpetty/avalanchego/vms/secp256k1fx"
)

// Credential ...
type Credential struct {
	secp256k1fx.Credential `serialize:"true"`
}
