// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BareMetal2GatewayClusterRefInventoryView BareMetal2GatewayClusterRef
type BareMetal2GatewayClusterRefInventoryView struct {
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"gatewayUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

