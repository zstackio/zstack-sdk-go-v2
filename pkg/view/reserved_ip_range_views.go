// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ReservedIpRangeInventoryView ReservedIpRange
type ReservedIpRangeInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	StartIp string `json:"startIp,omitempty"`
	EndIp string `json:"endIp,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// DeleteReservedIpRangeEventView DeleteReservedIpRangeEvent
type DeleteReservedIpRangeEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddReservedIpRangeEventView AddReservedIpRangeEvent
type AddReservedIpRangeEventView struct {
	Inventory ReservedIpRangeInventoryView `json:"inventory,omitempty"`
}

