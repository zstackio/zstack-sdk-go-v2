// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BaremetalPxeServerClusterRefInventoryView BaremetalPxeServerClusterRef
type BaremetalPxeServerClusterRefInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"pxeServerUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

