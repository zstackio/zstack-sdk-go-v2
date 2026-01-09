// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcVpnIpSecConfigInventoryView VpcVpnIpSecConfig
type VpcVpnIpSecConfigInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	AccountName *string `json:"accountName,omitempty"`
	Name string `json:"name,omitempty"`
	EncodeAlgorithm *string `json:"encodeAlgorithm,omitempty"`
	AuthAlgorithm *string `json:"authAlgorithm,omitempty"`
	Pfs *string `json:"pfs,omitempty"`
	Lifetime *int `json:"lifetime,omitempty"`
	Description *string `json:"description,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

// CreateVpnIpsecConfigEventView CreateVpnIpsecConfigEvent
type CreateVpnIpsecConfigEventView struct {
	Inventory VpcVpnIpSecConfigInventoryView `json:"inventory,omitempty"`
}

// QueryVpcIpSecConfigFromLocalView QueryVpcIpSecConfigFromLocal
type QueryVpcIpSecConfigFromLocalView struct {
	Inventories []VpcVpnIpSecConfigInventoryView `json:"inventories,omitempty"`
}

