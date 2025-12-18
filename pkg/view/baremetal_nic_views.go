// Copyright (c) ZStack.io, Inc.

package view

import "time"

// BaremetalNicInventoryView BaremetalNic
type BaremetalNicInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"baremetalInstanceUuid,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"baremetalBondingUuid,omitempty"`
	rest string `json:"mac,omitempty"`
	rest string `json:"ip,omitempty"`
	rest string `json:"netmask,omitempty"`
	rest string `json:"gateway,omitempty"`
	rest string `json:"metadata,omitempty"`
	rest bool `json:"pxe,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

