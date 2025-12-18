// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VipPeerL3NetworkRefInventoryView VipPeerL3NetworkRef
type VipPeerL3NetworkRefInventoryView struct {
	rest string `json:"vipUuid,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

