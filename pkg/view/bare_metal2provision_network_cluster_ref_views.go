// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2ProvisionNetworkClusterRefInventoryView BareMetal2ProvisionNetworkClusterRef
type BareMetal2ProvisionNetworkClusterRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	ClusterUuid string `json:"clusterUuid,omitempty"`
	NetworkUuid string `json:"networkUuid,omitempty"`
}

