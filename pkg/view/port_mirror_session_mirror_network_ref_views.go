// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PortMirrorSessionMirrorNetworkRefInventoryView PortMirrorSessionMirrorNetworkRef
type PortMirrorSessionMirrorNetworkRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	SrcTunnelUuid string `json:"srcTunnelUuid,omitempty"`
	DstTunnelUuid string `json:"dstTunnelUuid,omitempty"`
	Type string `json:"type,omitempty"`
	SessionUuid string `json:"sessionUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	SrcTunnel MirrorNetworkUsedIpInventoryView `json:"srcTunnel,omitempty"`
	DstTunnel MirrorNetworkUsedIpInventoryView `json:"dstTunnel,omitempty"`
}

