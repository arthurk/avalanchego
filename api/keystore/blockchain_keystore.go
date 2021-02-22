// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package keystore

import (
	"github.com/corpetty/avalanchego/database"
	"github.com/corpetty/avalanchego/ids"
)

// BlockchainKeystore ...
type BlockchainKeystore struct {
	blockchainID ids.ID
	ks           *Keystore
}

// GetDatabase ...
func (bks *BlockchainKeystore) GetDatabase(username, password string) (database.Database, error) {
	return bks.ks.GetDatabase(bks.blockchainID, username, password)
}
