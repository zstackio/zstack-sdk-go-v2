// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcSharedQosInventoryView VpcSharedQos
type VpcSharedQosInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description *string `json:"description,omitempty"`
	L3NetworkUuid *string `json:"l3NetworkUuid,omitempty"`
	VpcUuid *string `json:"vpcUuid,omitempty"`
	Bandwidth *int64 `json:"bandwidth,omitempty"`
	Vips []VpcSharedQosRefVipInventoryView `json:"vips,omitempty"`
}

// CreateVpcSharedQosEventView CreateVpcSharedQosEvent
type CreateVpcSharedQosEventView struct {
	Inventory VpcSharedQosInventoryView `json:"inventory,omitempty"`
}

// QueryVpcSharedQosView QueryVpcSharedQos
type QueryVpcSharedQosView struct {
	Inventories []VpcSharedQosInventoryView `json:"inventories,omitempty"`
}

// UpdateVpcSharedQosEventView UpdateVpcSharedQosEvent
type UpdateVpcSharedQosEventView struct {
	Inventory VpcSharedQosInventoryView `json:"inventory,omitempty"`
}

// ChangeVpcSharedQosBandwidthEventView ChangeVpcSharedQosBandwidthEvent
type ChangeVpcSharedQosBandwidthEventView struct {
	Inventory VpcSharedQosInventoryView `json:"inventory,omitempty"`
}

// DeleteVpcSharedQosEventView DeleteVpcSharedQosEvent
type DeleteVpcSharedQosEventView struct {
	Success bool `json:"success,omitempty"`
}

