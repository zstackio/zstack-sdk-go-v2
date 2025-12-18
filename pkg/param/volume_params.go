// Copyright (c) ZStack.io, Inc.

package param

// UpdateVolumeDetailParam UpdateVolume详细参数
type UpdateVolumeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateVolumeParam UpdateVolume请求参数
type UpdateVolumeParam struct {
	BaseParam
	Params UpdateVolumeDetailParam `json:"params"` // 详细参数
}

