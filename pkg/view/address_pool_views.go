// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AddressPoolInventoryView AddressPool
type AddressPoolInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"startIp,omitempty"`
	rest string `json:"endIp,omitempty"`
	rest string `json:"netmask,omitempty"`
	rest string `json:"gateway,omitempty"`
	rest string `json:"networkCidr,omitempty"`
	rest int `json:"ipVersion,omitempty"`
	rest string `json:"addressMode,omitempty"`
	rest int `json:"prefixLen,omitempty"`
	rest string `json:"ipRangeType,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

