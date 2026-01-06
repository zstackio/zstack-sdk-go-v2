// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2GatewayClusterRefInventoryView BareMetal2GatewayClusterRef
type BareMetal2GatewayClusterRefInventoryView struct {
	ClusterUuid string `json:"clusterUuid,omitempty"`
	GatewayUuid string `json:"gatewayUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

