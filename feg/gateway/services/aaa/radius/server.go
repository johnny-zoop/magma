/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

// package radius implements AAA server's radius interface for accounting & authentication
package radius

import (
	"magma/feg/cloud/go/protos/mconfig"
	"magma/feg/gateway/services/aaa/protos"
)

// AuthServer - radius EAP server implementation
type AuthServer struct {
	authenticator protos.AuthenticatorServer
	authMethods   []byte
}

// AcctServer - radius accounting server implementation
type AcctServer struct {
	accounting protos.AccountingServer
}

// Server - radius server implementation
type Server struct {
	AuthServer
	AcctServer
	config *mconfig.RadiusConfig
}

// GetConfig returns server configs
func (s *Server) GetConfig() *mconfig.RadiusConfig {
	if s == nil {
		return defaultConfigs
	}
	return s.config
}

// New returns new radius server
func New(cfg *mconfig.RadiusConfig, authRPC protos.AuthenticatorServer, acctRpc protos.AccountingServer) *Server {
	return &Server{
		config:     validateConfigs(cfg),
		AuthServer: AuthServer{authenticator: authRPC},
		AcctServer: AcctServer{accounting: acctRpc},
	}
}