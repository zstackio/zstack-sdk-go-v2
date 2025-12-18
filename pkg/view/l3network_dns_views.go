// Copyright (c) ZStack.io, Inc.

package view

import "time"

// L3NetworkDnsInventoryView L3NetworkDns
type L3NetworkDnsInventoryView struct {
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"dns,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

