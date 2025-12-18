// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VipQosInventoryView VipQos
type VipQosInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	Port int `json:"port,omitempty"`
	InboundBandwidth int64 `json:"inboundBandwidth,omitempty"`
	OutboundBandwidth int64 `json:"outboundBandwidth,omitempty"`
	Type string `json:"type,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

