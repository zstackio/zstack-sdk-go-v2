// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcVpnIkeConfigInventoryView VpcVpnIkeConfig
type VpcVpnIkeConfigInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccountName *string `json:"accountName,omitempty"`
	Psk *string `json:"psk,omitempty"`
	Version *string `json:"version,omitempty"`
	Mode *string `json:"mode,omitempty"`
	EncodeAlgorithm *string `json:"encodeAlgorithm,omitempty"`
	AuthAlgorithm *string `json:"authAlgorithm,omitempty"`
	Pfs *string `json:"pfs,omitempty"`
	Lifetime *int `json:"lifetime,omitempty"`
	LocalIp *string `json:"localIp,omitempty"`
	RemoteIp *string `json:"remoteIp,omitempty"`
	Description *string `json:"description,omitempty"`
}

// CreateVpnIkeConfigEventView CreateVpnIkeConfigEvent
type CreateVpnIkeConfigEventView struct {
	Inventory VpcVpnIkeConfigInventoryView `json:"inventory,omitempty"`
}

// QueryVpcIkeConfigFromLocalView QueryVpcIkeConfigFromLocal
type QueryVpcIkeConfigFromLocalView struct {
	Inventories []VpcVpnIkeConfigInventoryView `json:"inventories,omitempty"`
}

