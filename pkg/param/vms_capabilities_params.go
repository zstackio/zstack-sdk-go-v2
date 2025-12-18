// Copyright (c) ZStack.io, Inc.

package param

// GetVmsCapabilitiesDetailParam GetVmsCapabilities详细参数
type GetVmsCapabilitiesDetailParam struct {
	rest []string `json:"vmUuids" validate:"required"` // 必填
}

// GetVmsCapabilitiesParam GetVmsCapabilities请求参数
type GetVmsCapabilitiesParam struct {
	BaseParam
	Params GetVmsCapabilitiesDetailParam `json:"params"` // 详细参数
}

