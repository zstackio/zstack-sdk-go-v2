// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VipPeerL3NetworkRefInventoryView VipPeerL3NetworkRef
type VipPeerL3NetworkRefInventoryView struct {
	VipUuid string `json:"vipUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

