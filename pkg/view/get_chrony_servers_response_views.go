// Copyright (c) ZStack.io, Inc.

package view

// GetChronyServersView GetChronyServers
type GetChronyServersView struct {
	Servers []ChronyServerInfoPairView `json:"servers,omitempty"`
	Success bool `json:"success,omitempty"`
}

