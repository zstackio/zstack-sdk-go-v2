// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteVipQosParamDetail DeleteVipQos detail param
type DeleteVipQosParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Port int `json:"port,omitempty"`
}

// DeleteVipQosParam DeleteVipQos request param
type DeleteVipQosParam struct {
	BaseParam
	DeleteVipQos DeleteVipQosParamDetail `json:"deleteVipQos"`
}
// SetVipQosParamDetail SetVipQos detail param
type SetVipQosParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Port int `json:"port,omitempty"`
	OutboundBandwidth int64 `json:"outboundBandwidth,omitempty"`
	InboundBandwidth int64 `json:"inboundBandwidth,omitempty"`
}

// SetVipQosParam SetVipQos request param
type SetVipQosParam struct {
	BaseParam
	SetVipQos SetVipQosParamDetail `json:"setVipQos"`
}
// GetVipQosParamDetail GetVipQos detail param
type GetVipQosParamDetail struct {
	Uuid string `json:"uuid,omitempty"`
}

// GetVipQosParam GetVipQos request param
type GetVipQosParam struct {
	BaseParam
	GetVipQos GetVipQosParamDetail `json:"getVipQos"`
}
