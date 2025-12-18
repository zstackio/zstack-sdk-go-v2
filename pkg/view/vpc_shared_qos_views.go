// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcSharedQosInventoryView VpcSharedQos
type VpcSharedQosInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	VpcUuid string `json:"vpcUuid,omitempty"`
	Bandwidth int64 `json:"bandwidth,omitempty"`
	Vips []VpcSharedQosRefVipInventoryView `json:"vips,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

