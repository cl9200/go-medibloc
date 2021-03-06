// Copyright (C) 2018  MediBloc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package rpc

import (
	"net"

	"github.com/medibloc/go-medibloc/core"
	"github.com/medibloc/go-medibloc/medlet/pb"
	"github.com/medibloc/go-medibloc/rpc/pb"
	"github.com/medibloc/go-medibloc/util/logging"
	"google.golang.org/grpc"
)

// Server is rpc server.
type Server struct {
	addrGrpc  string
	addrHTTP  string
	rpcServer *grpc.Server
}

// New returns NewServer.
func New(cfg *medletpb.Config) *Server {
	rpc := grpc.NewServer()
	return &Server{
		rpcServer: rpc,
		addrGrpc:  cfg.Rpc.RpcListen[0],
		addrHTTP:  cfg.Rpc.HttpListen[0],
	}
}

//Setup sets up server.
func (s *Server) Setup(bm *core.BlockManager, tm *core.TransactionManager) {
	api := newAPIService(bm, tm)
	rpcpb.RegisterApiServiceServer(s.rpcServer, api)
}

// Start starts rpc server.
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.addrGrpc)
	if err != nil {
		return err
	}
	go func() {
		if err := s.rpcServer.Serve(lis); err != nil {
			logging.Console().Error(err)
		}
	}()
	logging.Console().Info("GRPC Server is running...")

	s.RunGateway()
	return nil
}

// RunGateway runs rest gateway server.
func (s *Server) RunGateway() error {
	go func() {
		if err := httpServerRun(s.addrHTTP, s.addrGrpc); err != nil {
			logging.Console().Error(err)
		}
	}()
	logging.Console().Info("GRPC HTTP Gateway is running...")
	return nil
}

// Stop stops server.
func (s *Server) Stop() {
	s.rpcServer.Stop()
}
