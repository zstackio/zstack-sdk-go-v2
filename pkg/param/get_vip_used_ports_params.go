// Copyright (c) ZStack.io, Inc.

package param

// GetVipUsedPortsDetailParam GetVipUsedPorts detail param
type GetVipUsedPortsDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Protocol string `json:"protocol" validate:"required"`
}

// GetVipUsedPortsParam GetVipUsedPorts request param
type GetVipUsedPortsParam struct {
	BaseParam
	Params GetVipUsedPortsDetailParam `json:"params"`
}
