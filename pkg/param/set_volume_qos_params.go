// Copyright (c) ZStack.io, Inc.

package param

// SetVolumeQosDetailParam SetVolumeQos详细参数
type SetVolumeQosDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"mode,omitempty"`
	rest int64 `json:"volumeBandwidth,omitempty"`
	rest int64 `json:"readBandwidth,omitempty"`
	rest int64 `json:"writeBandwidth,omitempty"`
	rest int64 `json:"totalBandwidth,omitempty"`
	rest int64 `json:"readIOPS,omitempty"`
	rest int64 `json:"writeIOPS,omitempty"`
	rest int64 `json:"totalIOPS,omitempty"`
}

// SetVolumeQosParam SetVolumeQos请求参数
type SetVolumeQosParam struct {
	BaseParam
	Params SetVolumeQosDetailParam `json:"params"` // 详细参数
}

