// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BareMetal2ProvisionNetworkClusterRefInventoryView BareMetal2ProvisionNetworkClusterRef
type BareMetal2ProvisionNetworkClusterRefInventoryView struct {
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"networkUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

