// Copyright (c) ZStack.io, Inc.

package view

import "time"

// MirrorNetworkUsedIpInventoryView MirrorNetworkUsedIp
type MirrorNetworkUsedIpInventoryView struct {
	rest UsedIpInventoryView `json:"usedIpInventory,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"uuid,omitempty"`
}

