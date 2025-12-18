// Copyright (c) ZStack.io, Inc.

package param

// UpdateXskyBlockVolumeDetailParam UpdateXskyBlockVolume详细参数
type UpdateXskyBlockVolumeDetailParam struct {
	rest int64 `json:"burstTotalBw,omitempty"`
	rest int64 `json:"burstTotalIops,omitempty"`
	rest int64 `json:"maxTotalBw,omitempty"`
	rest int64 `json:"maxTotalIops,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateXskyBlockVolumeParam UpdateXskyBlockVolume请求参数
type UpdateXskyBlockVolumeParam struct {
	BaseParam
	Params UpdateXskyBlockVolumeDetailParam `json:"params"` // 详细参数
}

