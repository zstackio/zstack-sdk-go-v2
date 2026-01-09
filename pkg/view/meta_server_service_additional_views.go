// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MetaServerServiceView MetaServerService
type MetaServerServiceView struct {
	Url *string `json:"url,omitempty"`
	Up bool `json:"up,omitempty"`
	Role *string `json:"role,omitempty"`
	UsedMemoryInBytes int64 `json:"usedMemoryInBytes,omitempty"`
	SystemMemoryInBytes int64 `json:"systemMemoryInBytes,omitempty"`
	MaxMemoryInBytes int64 `json:"maxMemoryInBytes,omitempty"`
	ConnectedClients int64 `json:"connectedClients,omitempty"`
	MaxClients int64 `json:"maxClients,omitempty"`
	SyncInProgress bool `json:"syncInProgress,omitempty"`
	ReplLagBytes int64 `json:"replLagBytes,omitempty"`
	Version *string `json:"version,omitempty"`
}

