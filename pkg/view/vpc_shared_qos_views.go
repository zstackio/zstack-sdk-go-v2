// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcSharedQosInventoryView VpcSharedQos
type VpcSharedQosInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"vpcUuid,omitempty"`
	rest int64 `json:"bandwidth,omitempty"`
	rest []VpcSharedQosRefVipInventoryView `json:"vips,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

