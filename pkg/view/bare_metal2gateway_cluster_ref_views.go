// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2GatewayClusterRefInventoryView BareMetal2GatewayClusterRef
type BareMetal2GatewayClusterRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	GatewayUuid *string `json:"gatewayUuid,omitempty"`
}

