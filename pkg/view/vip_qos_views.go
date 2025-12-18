// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VipQosInventoryView VipQos
type VipQosInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"vipUuid,omitempty"`
	rest int `json:"port,omitempty"`
	rest int64 `json:"inboundBandwidth,omitempty"`
	rest int64 `json:"outboundBandwidth,omitempty"`
	rest string `json:"type,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

