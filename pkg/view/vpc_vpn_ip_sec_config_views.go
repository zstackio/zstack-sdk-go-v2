// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VpcVpnIpSecConfigInventoryView VpcVpnIpSecConfig
type VpcVpnIpSecConfigInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccountName string `json:"accountName,omitempty"`
	EncodeAlgorithm string `json:"encodeAlgorithm,omitempty"`
	AuthAlgorithm string `json:"authAlgorithm,omitempty"`
	Pfs string `json:"pfs,omitempty"`
	Lifetime int `json:"lifetime,omitempty"`
	Description string `json:"description,omitempty"`
}

// CreateVpnIpsecConfigEventView CreateVpnIpsecConfigEvent
type CreateVpnIpsecConfigEventView struct {
	Inventory VpcVpnIpSecConfigInventoryView `json:"inventory,omitempty"`
}

// QueryVpcIpSecConfigFromLocalView QueryVpcIpSecConfigFromLocal
type QueryVpcIpSecConfigFromLocalView struct {
	Inventories []VpcVpnIpSecConfigInventoryView `json:"inventories,omitempty"`
}

