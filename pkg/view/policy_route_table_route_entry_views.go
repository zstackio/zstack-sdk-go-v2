// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PolicyRouteTableRouteEntryInventoryView PolicyRouteTableRouteEntry
type PolicyRouteTableRouteEntryInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"tableUuid,omitempty"`
	rest string `json:"destinationCidr,omitempty"`
	rest string `json:"nextHopIp,omitempty"`
	rest int `json:"distance,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

