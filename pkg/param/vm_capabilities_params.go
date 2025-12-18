// Copyright (c) ZStack.io, Inc.

package param

// GetVmCapabilitiesDetailParam GetVmCapabilities详细参数
type GetVmCapabilitiesDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVmCapabilitiesParam GetVmCapabilities请求参数
type GetVmCapabilitiesParam struct {
	BaseParam
	Params GetVmCapabilitiesDetailParam `json:"params"` // 详细参数
}

