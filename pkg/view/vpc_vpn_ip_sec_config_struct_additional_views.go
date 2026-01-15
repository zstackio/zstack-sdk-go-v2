// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VpcVpnIpSecConfigStructView VpcVpnIpSecConfigStruct
type VpcVpnIpSecConfigStructView struct {
	IpsecEncAlg string `json:"IpsecEncAlg,omitempty"`
	IpsecAuthAlg string `json:"IpsecAuthAlg,omitempty"`
	IpsecPfs string `json:"IpsecPfs,omitempty"`
	IpsecLifetime int `json:"IpsecLifetime,omitempty"`
}

