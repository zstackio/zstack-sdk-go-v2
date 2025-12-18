// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VpcVpnIkeConfigInventoryView VpcVpnIkeConfig
type VpcVpnIkeConfigInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"accountName,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"psk,omitempty"`
	rest string `json:"version,omitempty"`
	rest string `json:"mode,omitempty"`
	rest string `json:"encodeAlgorithm,omitempty"`
	rest string `json:"authAlgorithm,omitempty"`
	rest string `json:"pfs,omitempty"`
	rest int `json:"lifetime,omitempty"`
	rest string `json:"localIp,omitempty"`
	rest string `json:"remoteIp,omitempty"`
	rest string `json:"description,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

