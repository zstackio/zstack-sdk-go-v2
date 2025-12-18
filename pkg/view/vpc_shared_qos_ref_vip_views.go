// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcSharedQosRefVipInventoryView VpcSharedQosRefVip
type VpcSharedQosRefVipInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest string `json:"sharedQosUuid,omitempty"`
	rest string `json:"vipUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

