// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcVpnIpSecConfigInventoryView VpcVpnIpSecConfig
type VpcVpnIpSecConfigInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"accountName,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"encodeAlgorithm,omitempty"`
	rest string `json:"authAlgorithm,omitempty"`
	rest string `json:"pfs,omitempty"`
	rest int `json:"lifetime,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

