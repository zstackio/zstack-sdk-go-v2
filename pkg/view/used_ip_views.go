// Copyright (c) ZStack.io, Inc.

package view

import "time"

// UsedIpInventoryView UsedIp
type UsedIpInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"ipRangeUuid,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest int `json:"ipVersion,omitempty"`
	rest string `json:"ip,omitempty"`
	rest string `json:"netmask,omitempty"`
	rest string `json:"gateway,omitempty"`
	rest string `json:"usedFor,omitempty"`
	rest int64 `json:"ipInLong,omitempty"`
	rest interface{} `json:"ipInBinary,omitempty"`
	rest string `json:"vmNicUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

