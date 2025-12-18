// Copyright (c) ZStack.io, Inc.

package param

// UpdateBlockVolumeDetailParam UpdateBlockVolume详细参数
type UpdateBlockVolumeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateBlockVolumeParam UpdateBlockVolume请求参数
type UpdateBlockVolumeParam struct {
	BaseParam
	Params UpdateBlockVolumeDetailParam `json:"params"` // 详细参数
}

