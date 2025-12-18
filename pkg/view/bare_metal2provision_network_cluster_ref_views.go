// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BareMetal2ProvisionNetworkClusterRefInventoryView BareMetal2ProvisionNetworkClusterRef
type BareMetal2ProvisionNetworkClusterRefInventoryView struct {
	ClusterUuid string `json:"clusterUuid,omitempty"`
	NetworkUuid string `json:"networkUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

