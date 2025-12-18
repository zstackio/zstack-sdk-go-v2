// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PortMirrorSessionMirrorNetworkRefInventoryView PortMirrorSessionMirrorNetworkRef
type PortMirrorSessionMirrorNetworkRefInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"srcTunnelUuid,omitempty"`
	rest string `json:"dstTunnelUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"sessionUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest MirrorNetworkUsedIpInventoryView `json:"srcTunnel,omitempty"`
	rest MirrorNetworkUsedIpInventoryView `json:"dstTunnel,omitempty"`
}

