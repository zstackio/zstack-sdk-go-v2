// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcVirtualRouteEntryInventoryView VpcVirtualRouteEntry
type VpcVirtualRouteEntryInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"vRouterType,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"destinationCidrBlock,omitempty"`
	rest string `json:"nextHopId,omitempty"`
	rest string `json:"virtualRouterUuid,omitempty"`
	rest string `json:"nextHopType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

