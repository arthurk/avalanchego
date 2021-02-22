// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gsubnetlookup

import (
	"context"

	"github.com/corpetty/avalanchego/ids"
	"github.com/corpetty/avalanchego/snow"
	"github.com/corpetty/avalanchego/vms/rpcchainvm/gsubnetlookup/gsubnetlookupproto"
)

// Server is a messenger that is managed over RPC.
type Server struct {
	aliaser snow.SubnetLookup
}

// NewServer returns a vm instance connected to a remote vm instance
func NewServer(aliaser snow.SubnetLookup) *Server {
	return &Server{aliaser: aliaser}
}

// SubnetID ...
func (s *Server) SubnetID(
	_ context.Context,
	req *gsubnetlookupproto.SubnetIDRequest,
) (*gsubnetlookupproto.SubnetIDResponse, error) {
	chainID, err := ids.ToID(req.ChainID)
	if err != nil {
		return nil, err
	}
	id, err := s.aliaser.SubnetID(chainID)
	if err != nil {
		return nil, err
	}
	return &gsubnetlookupproto.SubnetIDResponse{
		Id: id[:],
	}, nil
}
