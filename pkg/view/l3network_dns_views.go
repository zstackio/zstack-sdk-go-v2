// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// L3NetworkDnsInventoryView L3NetworkDns
type L3NetworkDnsInventoryView struct {
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	Dns string `json:"dns,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

