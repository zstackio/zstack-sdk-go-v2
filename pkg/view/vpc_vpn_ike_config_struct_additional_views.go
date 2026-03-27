// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VpcVpnIkeConfigStructView VpcVpnIkeConfigStruct
type VpcVpnIkeConfigStructView struct {
	Psk string `json:"Psk,omitempty"`
	IkeVersion string `json:"IkeVersion,omitempty"`
	IkeMode string `json:"IkeMode,omitempty"`
	IkeEncAlg string `json:"IkeEncAlg,omitempty"`
	IkeAuthAlg string `json:"IkeAuthAlg,omitempty"`
	IkePfs string `json:"IkePfs,omitempty"`
	IkeLifetime int `json:"IkeLifetime,omitempty"`
	LocalId string `json:"LocalId,omitempty"`
	RemoteId string `json:"RemoteId,omitempty"`
}

