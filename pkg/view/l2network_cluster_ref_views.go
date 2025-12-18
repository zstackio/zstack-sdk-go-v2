// Copyright (c) ZStack.io, Inc.

package view

import "time"

// L2NetworkClusterRefInventoryView L2NetworkClusterRef
type L2NetworkClusterRefInventoryView struct {
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"l2NetworkUuid,omitempty"`
	rest string `json:"l2ProviderType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

