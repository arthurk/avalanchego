// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertex

import (
	"github.com/corpetty/avalanchego/ids"
	"github.com/corpetty/avalanchego/snow/consensus/snowstorm"
	"github.com/corpetty/avalanchego/snow/engine/common"
)

// DAGVM defines the minimum functionality that an avalanche VM must
// implement
type DAGVM interface {
	common.VM

	// Return any transactions that have not been sent to consensus yet
	Pending() []snowstorm.Tx

	// Convert a stream of bytes to a transaction or return an error
	Parse(tx []byte) (snowstorm.Tx, error)

	// Retrieve a transaction that was submitted previously
	Get(ids.ID) (snowstorm.Tx, error)
}
