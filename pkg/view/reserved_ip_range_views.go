// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ReservedIpRangeInventoryView ReservedIpRange
type ReservedIpRangeInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"startIp,omitempty"`
	rest string `json:"endIp,omitempty"`
	rest int `json:"ipVersion,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

