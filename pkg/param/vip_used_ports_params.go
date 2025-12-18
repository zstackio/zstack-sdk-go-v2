// Copyright (c) ZStack.io, Inc.

package param

// GetVipUsedPortsDetailParam GetVipUsedPorts详细参数
type GetVipUsedPortsDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"protocol" validate:"required"` // 必填
}

// GetVipUsedPortsParam GetVipUsedPorts请求参数
type GetVipUsedPortsParam struct {
	BaseParam
	Params GetVipUsedPortsDetailParam `json:"params"` // 详细参数
}

