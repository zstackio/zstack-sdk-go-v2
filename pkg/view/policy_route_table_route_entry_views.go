// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PolicyRouteTableRouteEntryInventoryView PolicyRouteTableRouteEntry
type PolicyRouteTableRouteEntryInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	TableUuid string `json:"tableUuid,omitempty"`
	DestinationCidr string `json:"destinationCidr,omitempty"`
	NextHopIp string `json:"nextHopIp,omitempty"`
	Distance int `json:"distance,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

