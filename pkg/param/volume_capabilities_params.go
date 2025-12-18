// Copyright (c) ZStack.io, Inc.

package param

// GetVolumeCapabilitiesDetailParam GetVolumeCapabilities详细参数
type GetVolumeCapabilitiesDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVolumeCapabilitiesParam GetVolumeCapabilities请求参数
type GetVolumeCapabilitiesParam struct {
	BaseParam
	Params GetVolumeCapabilitiesDetailParam `json:"params"` // 详细参数
}

